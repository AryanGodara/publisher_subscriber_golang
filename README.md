## To run the program (Now, root directory of repo, contains the proper files):-
```
$ go run server.go

In different terminal(s)
$ go run client.go   (in as many terminals as you want)


For sending HTTP post requests, use Postman:
POST: localhost:4000/postmessage

In body - "Type out anything you want"
```


## client.go and server.go in the root directory of repo :-
### Instead of a hub to store the clients. This makes use of global variables: `msg` and `flag`.
1. Whenever a POST request is sent to the HTTP server, it updates the value of global variable `msg`.
2. Then, it marks the `flag` bool variable as true
3. The websocket Handler if always listening, and in the case when `flag == true`: it sends the `msg` to the client, and turns `flag = false`  

```
$ cd pub-sub-hub
$ go run server.go

In different terminal(s)
$ go run client.go   (in as many terminals as you want)
```
