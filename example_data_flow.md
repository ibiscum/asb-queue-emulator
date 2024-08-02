## 1. Using HTTP (POST command)

1. **Client Application**: The client application wants to send a message to the queue. It makes an HTTP POST request directly to the `/queue/message` endpoint of the emulator.

2. **Azure Service Bus Emulator's API**: 
   - Receives the HTTP POST request from the client.
   - Parses the message from the request body.
   - Processes any necessary logic (e.g., validation, formatting).

3. **Communication with RabbitMQ**:
   - The emulator's API sends the message to the RabbitMQ instance, likely using a RabbitMQ client library in Go.
   - RabbitMQ receives the message and adds it to the appropriate queue.
   - RabbitMQ sends an acknowledgment back to the emulator's API.

4. **Response to Client**:
   - Once the acknowledgment is received from RabbitMQ, the emulator's API sends an HTTP response back to the client, indicating that the message has been queued successfully.

## 2. Using AMQP (from client's SDK)

1. **Client using Azure SDK**: The client wants to send a message to the queue. It uses the Azure SDK to do this, which will attempt to send the message over AMQP.

2. **AMQP to HTTP Gateway in the Emulator**:
   - The gateway listens for AMQP messages.
   - Upon receiving the AMQP request from the client's SDK, the gateway translates this into an HTTP POST request to the emulator's API (the same `/queue/message` endpoint as in the HTTP flow).

3. **Azure Service Bus Emulator's API**:
   - Receives the translated HTTP POST request from the gateway.
   - Parses the message as it would in the direct HTTP flow.
   - Sends the message to RabbitMQ and waits for an acknowledgment.

4. **Communication with RabbitMQ**:
   - Identical to the HTTP flow, the API sends the message to RabbitMQ.
   - RabbitMQ confirms by sending an acknowledgment back to the emulator's API.

5. **Response to Client**:
   - The emulator's API sends a response back to the AMQP to HTTP Gateway.
   - The gateway then translates this HTTP response into an appropriate AMQP response.
   - The client's Azure SDK receives this AMQP response, interpreting it as a confirmation from the Azure Service Bus.

In summary, the key difference in the AMQP flow is the introduction of the AMQP to HTTP Gateway. This gateway essentially translates AMQP operations to equivalent HTTP requests and vice versa, allowing the Azure SDK to interface with the emulator.