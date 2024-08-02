from proton import Message, Disposition, Link, Receiver, Sender
from proton._events import Event
from proton.handlers import MessagingHandler
from proton.reactor import Container
from proton import SSL, SSLDomain
import os
import amqp_translator as translator
import re
from proton import int32, timestamp
import argparse
# Import module for logging in a docker container
import logging
import sys
logging.basicConfig(stream=sys.stdout, level=logging.INFO)

RCV_SETTLE_MODES = {
    Link.RCV_FIRST: 1,
    Link.RCV_SECOND: 2,
}

SND_SETTLE_MODES = {
    Link.SND_MIXED: 1,
    Link.SND_UNSETTLED: 2,
    Link.SND_SETTLED: 3,
}

def parse_args():
    # We just need a --use-tls flag
    parser = argparse.ArgumentParser()
    parser.add_argument("--use-tls", action="store_true", default=False, required=False, help="Use TLS")
    return parser.parse_args()
    
class SimpleServer(MessagingHandler):
    
    def __init__(self, address, use_tls):
        super(SimpleServer, self).__init__()
        self.address = address
        self.queue_name = ""
        self.use_tls = use_tls
        self.instant_reply = False
        self.rcv_settle_mode = 0
        self.snd_settle_mode = 0
        self.stored_links = {}
        self.cert_name = "cert.pem"
        self.key_name = "key.pem"
        self.cert_location = "./certs"

    def on_start(self, event):
        if self.use_tls:
            logging.info("Starting server with TLS...")
            domain = SSLDomain(SSLDomain.MODE_SERVER)
            cert_location = os.getenv("CERT_LOCATION", self.cert_location)
            cert_path = os.path.join(cert_location, self.cert_name)
            key_path = os.path.join(cert_location, self.key_name)
            domain.set_credentials(cert_path, key_path, "")
            domain.set_peer_authentication(SSLDomain.ANONYMOUS_PEER)
            domain.set_trusted_ca_db(cert_path)
            domain.allow_unsecured_client()
            try:
                self.acceptor = event.container.listen(self.address, ssl_domain=domain)
            except Exception as e:
                logging.error("Error starting server:", e)
                raise e
        else:
            logging.info("Starting server...")
            self.acceptor = event.container.listen(self.address)

    def on_connection_init(self, event):
        logging.info("Connection initiated by a client.")

    def on_connection_bound(self, event):
        logging.info("Connection bound to transport.")

    def on_connection_error(self, event):
        logging.info("Connection error:", event.connection.remote_condition)

    def on_session_opened(self, event: Event):
        logging.info("Session opened.")

    def on_session_error(self, event):
        logging.info("Session error:", event.session.remote_condition)

    def on_session_closed(self, event):
        logging.info("Session closed.")


    def on_link_opened(self, event: Event):
        logging.info(f"Remote source: {event.link.remote_source}, Remote target: {event.link.remote_target}")
        logging.info(f"Remote source address: {event.link.remote_source.address}, Remote target address: {event.link.remote_target.address}")

        if not self.stored_links.get(event.link.remote_target.address, None):
            self.stored_links.update({event.link.remote_target.address: {}})
    
        if event.link.remote_target.address and event.link.is_sender:
            # This link wants us to send a message to the client
            self.stored_links[event.link.remote_target.address].update({"sender": event.link})
        elif event.link.remote_target.address and event.link.is_receiver:
            self.stored_links[event.link.remote_target.address].update({"receiver": event.link})

        logging.info(f"Current stored links: {self.stored_links}")
        if self.instant_reply and self.queue_name in event.link.remote_source.address:
            self._request_queue_message(event)

    def on_link_error(self, event):
        logging.info("Link error:", event.link.remote_condition)

    def on_link_closed(self, event):
        logging.info("Link closed.")
        # Remove the link from the stored links
        if event.link.name in self.stored_links:
            if event.link.is_sender:
                del self.stored_links[event.link.remote_target.address]["sender"]
            else:
                del self.stored_links[event.link.remote_target.address]["receiver"]
            
            if len(self.stored_links[event.link.remote_target.address]) == 0:
                del self.stored_links[event.link.remote_target.address]
        

    def on_link_init(self, event):
        link = event.link
        logging.info(f"Link initiated. Link name: {link.name}, Link remote source: {link.remote_source}, Link remote target: {link.remote_target}")
        if link.is_receiver:
            self.rcv_settle_mode = RCV_SETTLE_MODES[link.rcv_settle_mode]
            self.snd_settle_mode = SND_SETTLE_MODES[link.snd_settle_mode]
        if link.is_sender:
            self.rcv_settle_mode = RCV_SETTLE_MODES[1]
            self.snd_settle_mode = SND_SETTLE_MODES[2]

    def on_link_opening(self, event):
        logging.info("Link opening.")

    def on_link_remote_open(self, event: Event):
        logging.info("LINK REMOTE open triggered.")
        link = event.link
        if link.is_receiver: # Affects client sending to queue
            link.rcv_settle_mode = 0
            link.snd_settle_mode = 2
        else: # Affects the client requesting from queue
            link.snd_settle_mode = 1
            link.rcv_settle_mode = 1

    def on_message(self, event: Event):
        message: Message = event.message
        logging.info("Received message:", message)
        self.accept(event.delivery)
        logging.info("Message accepted.")

        if (message.reply_to and "cbs" in message.reply_to) or (message.properties and message.properties.get("operation", "") == "put-token"):
            # It's in the form of "amqps://<hostname>/<queue_name>"
            # It might also have "amqp://" instead of "amqps://
            # Finally it might be in the form of "amqps://<hostname>/<queue_name>/$cbs" or "amqps://<hostname>/<queue_name>/$management"
            # We need to remove the "/$cbs" or "/$management" from the end
            match = re.search(r'(amqp(s)?|sb)://.+?/(?P<queue_name>.+?)(/\$cbs|/\$management)?$', message.properties["name"])
            if match:
                self.queue_name = match.group("queue_name")
                logging.info("Interaction with queue:", self.queue_name)
            else:
                logging.info("Unable to find queue name in message properties.")
                return
            
            if not message.reply_to:
                message.reply_to = "$cbs"
            if "sb" in message.properties["name"]:
                # They aren't expecting to go through complete AMQP communication
                # We'll just start sending messages to the next receiver they open.
                self.instant_reply = True
            else:
                self.instant_reply = False
            self._send_auth_acceptance(event)
            logging.info("Sent acceptance message.")

        elif message.reply_to is not None:
            self._request_queue_message(event)
        
        elif message.body is not None:
            self._store_queue_message(event)
        else:
            logging.info("Unsure how to handle message")
        
    def _store_queue_message(self, event: Event):
        logging.info("Storing message in queue: ", self.queue_name)
        translator.HandleAmqpMessage(self.queue_name, event.message, translator.WRITE)
        
    def _request_queue_message(self, event: Event):
        logging.info("Requesting message from queue: ", self.queue_name)

        request = translator.HandleAmqpMessage(self.queue_name, None, translator.READ)

        request.content_type = "application/json"
        request.annotations = {
            "x-opt-enqueue-sequence-number": 2,
            "x-opt-enqueued-time": timestamp(1631582072.491),
            "x-opt-locked-until": timestamp(1631582072.491),
            "x-opt-sequence-number": 2144123,
        }

        target_address = ""
        correlation_id = 0
        if self.instant_reply:
            target_address = event.link.remote_target.address
        else:
            target_address = event.message.reply_to
            correlation_id = event.message.id

        request.correlation_id = correlation_id
        request.to = target_address
        request.properties = {
            'statusCode': int32(200),
            'statusDescription': 'OK',
            'com.microsoft:tracking-id': None,
            'errorCondition': None,
        }
        logging.info("Sending request:", request)
        # Convert the full request to bytes
        encoded_message = request.encode()
        # Convert the bytes to a memoryview
        memory_view_body = memoryview(encoded_message)
        request.body = {
            "messages": 
            [
                {"body": request.body,
                 "message": memory_view_body
                 },

            ]
        }

        link = self.stored_links[target_address]["sender"]
        self._send_client_message(link, request)
    
    def _send_auth_acceptance(self, event: Event):
        logging.info("Sending authentication acceptance message.")
        
        # Create a message with the required properties
        response = Message()
        response.body = ""
        response.correlation_id = event.message.id  # Set correlation_id to original message's message_id

        # Set the 'to' field to the 'reply_to' field of the original message
        response.to = event.message.reply_to

        response.properties = {
            'status-code': int32(202),
            'status-description': "Accepted"
        }
        
        logging.info(f"Stored links: {self.stored_links}")
        logging.info(f"Reply to: {event.message.reply_to}")
        link = self.stored_links[event.message.reply_to]["sender"]
        self._send_client_message(link, response)
        logging.info("Sent authentication acceptance message.")

    def _send_client_message(self, link, message):
        logging.info("Sending message to client:", message)
        link.send(message)
        logging.info("Sent message to client.")

if __name__ == "__main__":
    args = parse_args()
    try:
        host = os.getenv("HOST", "localhost")
        if args.use_tls:
            port = os.getenv("PORT", "5671")
        else:
            port = os.getenv("PORT", "5672")
        server = SimpleServer(host + ":" + port, args.use_tls)
        logging.info("Server running on " + host + ":" + port)
        Container(server).run()

    except KeyboardInterrupt:
        logging.info("\nServer interrupted by user. Exiting...")