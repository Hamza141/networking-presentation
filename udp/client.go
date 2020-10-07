package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	var host string
	if len(arguments) == 1 {
		host = "localhost:5126"
	} else {
		host = arguments[1]
	}

	c, err := net.Dial("udp", host)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	buf := make([]byte, 1024)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text)

		n, _ := c.Read(buf)
		fmt.Print("response: " + string(buf[0:n]))
	}
}
