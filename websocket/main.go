package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn){

	for {
		fmt.Println("trying to connect...")
		var reply string

		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't Receive")
			break
		}

		fmt.Println("Received back from Client:" + reply);

		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)

		if err := websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't Send", err)
			break
		}

	}

}

func main() {
	fmt.Println("Starting websocket server...")
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":8080", nil) ; err != nil {
		log.Fatal("ListenAndSever:", err);
	}

}
