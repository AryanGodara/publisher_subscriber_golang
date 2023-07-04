```
For RC: I have written many program after this one, bigger PRs, with more code and features, working with gRPC, etc.
But those are pieces of collaboration, and mostly use external libraries; and work with the rest of the codebase, but this is 
something completely independent, using only gorrila mux as that is even recommended on the official golang docs, as it's better than the
websocket implementation in the socket.

I had another repo of http-client in golang, but that was too big for review according to the question, so I did not submit that.
```

This server supports both http and websocket protocols on two different ports.

## To run the HTTP-plus-WebSocket Server

```bash
$ go run cmd/server/main.go
```

# To run the client, in different terminal(s)
```bash
$ go run cmmd/client/main.go   (in as many terminals as you want)
```

All the clients connected to this server will receive the echoed message from the postman client.

## For sending HTTP post requests, use Postman:
```
POST: localhost:4000/message
```
In body - "Type out anything you want in raw format of text or JSON"

## For sending Websocket requests, use Postman websocket, connect to
```
localhost:8080/socket
```
Then just send messages through the established websocket connection on the postman client, and it'll be echoed back to you.