# gpool_example

show how to use [gpool](https://github.com/cloudfstrife/gpool).

```
.
|-- cmd
|   |-- client                      //client main
|   |   `-- client.go              
|   `-- server                      //server main 
|       `-- server.go
|-- dial
|   |-- connection.go               //connection item 
|   `-- dial.go                     //pool 
|-- general.toml                    //config file 
|-- go.mod
`-- go.sum
```

## build 

```
go build ./cmd/client/
go build ./cmd/server/
./server
# open another terminal
./client
```
