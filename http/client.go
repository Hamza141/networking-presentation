package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	var host string
	if len(arguments) == 1 {
		host = "http://localhost:5124"
	} else {
		host = arguments[1]
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		resp, err := http.Post(host+"/send", "text/plain", strings.NewReader(text+"\n"))

		if err != nil {
			fmt.Println(err)
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Print("->: " + string(body))
	}
}
