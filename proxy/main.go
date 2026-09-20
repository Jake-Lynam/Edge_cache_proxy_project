package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	cache := NewLRU(5)

	http.HandleFunc("/", cache.handler)
	http.ListenAndServe(":9090", nil)
}

func (c *LRU_Cache) handler(w http.ResponseWriter, r *http.Request) {
	u := c.Get(r.URL.Path)
	if u != "" {
		w.Write([]byte(u))
		return
	}

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

	c.Set(r.URL.Path, string(body))
	w.Write(body) //sends the body to the client in form of bytes attached to the body of the packet
}

/*
listenAndServe needs to be equal to a var, need a way to represent 



*/