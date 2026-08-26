package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080") //makes request to server and gets resp/error
	if err != nil {
		fmt.Println("error making request:", err)
		return
	}
	defer resp.Body.Close() // closes stream after getting response

	body, err := io.ReadAll(resp.Body) //reads response to allow for manipulation in program
	if err != nil {
		fmt.Println("error reading body:", err)
		return
	}

	fmt.Println("got response:", string(body))
}
