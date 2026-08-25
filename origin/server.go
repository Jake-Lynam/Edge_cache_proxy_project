package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
