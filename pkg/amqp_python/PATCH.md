# Patching the ASB SDK
Currently, the ASB SDK has frames marked as optional which will crash the SDK if not included (not really optional). Additionally, the SDK does not easily support non-TLS connections.

For dealing with non-TLS (for now and Python specific), open up the SDK's _connection file. Within it, look for a condition written as 

```python
elif "sasl_credential" in kwargs:
```

Should be around line 128, change it to:

```python
elif "sasl_credential" not in kwargs:
```

In the _transport file, there is an AbstractTransport class. Replace the empty read/write functions with:

```python
    def _read(self, n, initial=False, buffer=None, _errnos=(errno.EAGAIN, errno.EINTR)):
        length = 0
        view = buffer or memoryview(bytearray(n))
        nbytes = self._read_buffer.readinto(view)
        toread = n - nbytes
        length += nbytes

        while toread:
            try:
                nbytes = self.sock.recv_into(view[length:])
            except socket.error as exc:
                if exc.errno in _errnos:
                    if initial:
                        raise socket.timeout()
                    continue
                raise
            if not nbytes:
                raise IOError("Server unexpectedly closed connection")

            length += nbytes
            toread -= nbytes

        return view

    def _write(self, s):
        """Completely write a string to the peer."""
        while s:
            try:
                n = self.sock.send(s)
            except ValueError:
                # Socket might be closed.
                n = 0
            if not n:
                raise IOError("Socket closed")
            s = s[n:]
```

As for the non-optional frame fixes. Back in the _connection file, look for frame[4] in the \_incoming\_open function. Add better conditionals:
```python
            if len(frame) > 4 and frame[4] is not None:
                self._remote_idle_timeout = frame[4] / 1000  # Convert to seconds
                self._remote_idle_timeout_send_frame = (
                    self._idle_timeout_empty_frame_send_ratio * self._remote_idle_timeout
                )
```

Go slightly below it and change the remote_properties object from frame[#] to empty:
```python
self._remote_properties = {}  # type: Dict[str, str]
```

In the "link" file, in \_incoming\_attach, set these two values to empty:
```python
        self.offered_capabilities = []  # offered_capabilities
        self.remote_properties = {}
```

Finally, in the receiver file, in \_incoming\_transfer update the received payload section as follows:
```python
            if self._received_payload or (len(frame) > 5 and frame[5]):  # more
                self._received_payload.extend(frame[-1])
            if frame[-1]:
                if self._received_payload:
                    message = decode_payload(memoryview(self._received_payload))
                    self._received_payload = bytearray()
                else:
                    message = decode_payload(frame[-1])
```