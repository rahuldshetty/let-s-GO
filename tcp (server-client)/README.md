# TCP Client & Server

Simple message transfer program built with TCP net component from GO library.

# Setup

Run the server: `go run server.go 4004`

Run the client: `go run client.go 127.0.0.1:4004`

Client can type the message and send it to server. 
Server responds back with timestamp.
Client can terminate the connection by sending "STOP" message.