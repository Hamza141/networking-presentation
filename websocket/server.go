package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	arguments := os.Args
	var PORT string
	if len(arguments) == 1 {
		PORT = ":5125"
	} else {
		PORT = ":" + arguments[1]
	}

	// SERVER_START OMIT
	upgrader := websocket.Upgrader{}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { // HL
		c, err := upgrader.Upgrade(w, r, nil) // HL
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			// incoming message
			mt, message, _ := c.ReadMessage() // HL
			fmt.Printf("->: %s", message)

			// send response
			err = c.WriteMessage(mt, []byte(time.Now().Format(time.RFC3339)+"\n")) // HL
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
	fmt.Println("Waiting for connection...")
	log.Fatal(http.ListenAndServe(PORT, nil))
	// SERVER_END OMIT
}
