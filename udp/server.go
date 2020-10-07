package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	arguments := os.Args
	var PORT int
	var err error
	if len(arguments) == 1 {
		PORT = 5126
	} else {
		PORT, err = strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// SERVER_START OMIT
	fmt.Println("Waiting for connection...")
	l, err := net.ListenUDP("udp", &net.UDPAddr{Port: PORT}) // HL
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close() // OMIT

	buf := make([]byte, 1024)
	for {
		// receive data
		n, retAddr, _ := l.ReadFromUDP(buf) // HL

		fmt.Print("-> ", string(buf[0:n]))

		// send data
		l.WriteToUDP([]byte(time.Now().Format(time.RFC3339) + "\n"), retAddr) // HL
	}
	// SERVER_END OMIT
}
