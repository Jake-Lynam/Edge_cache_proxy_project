package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080") //makes request to server and gets resp/error
	if err != nil {
		fmt.Println("error making request:", err)
		return
	}
	defer resp.Body.Close() // closes stream after getting response. Defer is a function that ensures the action it represents
							// completes even if the outer function ends earlier than expected. 

	body, err := io.ReadAll(resp.Body) //reads response to allow for manipulation in program
	if err != nil {
		fmt.Println("error reading body:", err)
		return
	}

	w.Write(body)
}

/*
listenAndServe needs to be equal to a var, need a way to represent 



*/