package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	var PORT string
	if len(arguments) == 1 {
		PORT = ":5123"
	} else {
		PORT = ":" + arguments[1]
	}

	// SERVER_START OMIT
	l, err := net.Listen("tcp", PORT) // HL
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close() // OMIT

	fmt.Println("Waiting for connection...")
	c, _ := l.Accept() // HL
	fmt.Println("Connected")

	for {
		// receive data
		netData, err := bufio.NewReader(c).ReadString('\n') // HL
		if err != nil {
			fmt.Println("Connection closed")
			return
		}

		fmt.Print("-> ", string(netData))

		// send data
		c.Write([]byte(time.Now().Format(time.RFC3339) + "\n")) // HL
	}
	// SERVER_END OMIT
}
