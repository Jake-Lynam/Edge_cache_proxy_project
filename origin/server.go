package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
