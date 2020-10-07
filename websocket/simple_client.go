package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	// "net/url"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	var host string
	if len(arguments) == 1 {
		host = "localhost:5125"
	} else {
		host = arguments[1]
	}

	// CLIENT_START OMIT
	c, _, err := websocket.DefaultDialer.Dial("ws://"+host+"/ws", nil) // HL
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	words := []string{"hello", "world", "this", "is", "a", "test"}

	for _, word := range words {
		fmt.Println("sending: " + word)
		c.WriteMessage(websocket.TextMessage, []byte(word+"\n")) // HL
		_, message, _ := c.ReadMessage()                         // HL
		fmt.Print("->: " + string(message))
		time.Sleep(2 * time.Second)
	}
	// CLIENT_END OMIT
}
