package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
)

func main() {
	arguments := os.Args
	var host string
	if len(arguments) == 1 {
		host = "localhost:5125"
	} else {
		host = arguments[1]
	}

	u := url.URL{Scheme: "ws", Host: host, Path: "/ws"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		c.WriteMessage(websocket.TextMessage, []byte(text))

		_, message, _ := c.ReadMessage()
		fmt.Print("->: " + string(message))
	}
}
