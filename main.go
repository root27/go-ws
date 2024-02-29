package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {

	// Socket server

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,

		CheckOrigin: func(r *http.Request) bool {

			return true

		},
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {

			log.Printf("Error: %v", err)

		}

		log.Println("Client connected")

		defer conn.Close()

		// Read message from client

		for {

			_, msg, err := conn.ReadMessage()

			if err != nil {

				log.Printf("Error reading message: %v", err)

			}

			log.Printf("Message from client: %v", string(msg))

			// Write message back to client

			if err = conn.WriteMessage(websocket.TextMessage, []byte("Thank you for your message")); err != nil {

				log.Printf("Error writing message: %v", err)

			}

		}

	})

	log.Println("Server listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
