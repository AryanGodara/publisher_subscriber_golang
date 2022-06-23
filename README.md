main.go currently is just a dummy implementation of pub-sub.

```
go run main.go
```

For working examples :-
- 1. godoc_websocket_eg
```
cd godoc_websocket_eg  
(First Terminal)  
go run server.go
(Second Terminal)  
go run client.go <message_you_want>
(Third Terminal)  
go run client.go <message_you_want>
```
It only prints the message being sent back to respective clients, not pub-sub model

- 2. gorilla_echo_eg
Just a copy of the example code of gorilla/websocket repository. Ignore this.