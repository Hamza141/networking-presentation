package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	var host string
	if len(arguments) == 1 {
		host = "http://localhost:5124"
	} else {
		host = arguments[1]
	}

	// CLIENT_START OMIT
	words := []string{"hello", "world", "this", "is", "a", "test"}

	for _, word := range words {
		resp, err := http.Post(host+"/send", "text/plain", strings.NewReader(word+"\n")) // HL

		if err != nil {
			fmt.Println(err)
			return
		}

		body, _ := ioutil.ReadAll(resp.Body) // HL
		fmt.Print("->: " + string(body))
		time.Sleep(2 * time.Second)
	}
	// CLIENT_END OMIT
}
