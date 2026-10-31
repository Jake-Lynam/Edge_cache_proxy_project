package main

import (
	"io"
	"log"
	"net/http"
)

const backend = "http://localhost:8080"

func main() {
	cache := NewLRU(5)

	http.HandleFunc("/", cache.handler)
	log.Println("proxy listening on :9090, forwarding to", backend)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func (c *LRU_Cache) handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.RequestURI() // path plus query string, e.g. "/search?q=dogs"

	if u := c.Get(key); u != "" {
		w.Write([]byte(u))
		return
	}

	resp, err := http.Get(backend + key) //requests the same path from the server and gets resp/error
	if err != nil {
		http.Error(w, "bad gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close() // closes stream after getting response. Defer is a function that ensures the action it represents
	// completes even if the outer function ends earlier than expected.

	body, err := io.ReadAll(resp.Body) //reads response to allow for manipulation in program
	if err != nil {
		http.Error(w, "bad gateway", http.StatusBadGateway)
		return
	}

	c.Set(key, string(body))
	w.Write(body) //sends the body to the client in form of bytes attached to the body of the packet
}

/*
listenAndServe needs to be equal to a var, need a way to represent



*/
