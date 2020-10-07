package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	var host string
	if len(arguments) == 1 {
		host = "localhost:5123"
	} else {
		host = arguments[1]
	}

	// CLIENT_START OMIT
	c, err := net.Dial("tcp", host) // HL
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	words := []string{"hello", "world", "this", "is", "a", "test"}

	for _, word := range words {
		fmt.Println("sending: " + word)
		fmt.Fprintf(c, word+"\n")                         // HL
		message, _ := bufio.NewReader(c).ReadString('\n') // HL
		fmt.Print("response: " + message)
		time.Sleep(2 * time.Second)
	}
	// CLIENT_END OMIT
}
