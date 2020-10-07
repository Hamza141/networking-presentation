package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	var PORT string
	if len(arguments) == 1 {
		PORT = ":5124"
	} else {
		PORT = ":" + arguments[1]
	}

	// SERVER_START OMIT
	fmt.Println("Waiting for connection...")

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		reader := bufio.NewReader(r.Body)
		text, _ := reader.ReadString('\n')
		// incoming request
		fmt.Print("-> ", text)

		// send response
		fmt.Fprintf(w, "%s", time.Now().Format(time.RFC3339)+"\n")
	})

	log.Fatal(http.ListenAndServe(PORT, nil))
	// SERVER_END OMIT
}
