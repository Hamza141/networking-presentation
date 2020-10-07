package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	var host string
	if len(arguments) == 1 {
		host = "localhost:5126"
	} else {
		host = arguments[1]
	}

	// CLIENT_START OMIT
	c, err := net.Dial("udp", host) // HL
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close() // OMIT

	words := []string{"hello", "world", "this", "is", "a", "test"}

	buf := make([]byte, 1024)
	for _, word := range words {
		fmt.Println("sending: " + word)
		fmt.Fprintf(c, word+"\n") // HL
		n, _ := c.Read(buf)       // HL
		fmt.Print("response: " + string(buf[0:n]))
		time.Sleep(2 * time.Second)
	}
	// CLIENT_END OMIT
}
