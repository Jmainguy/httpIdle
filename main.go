package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Connection received")

	fmt.Fprintf(w, "hello\n")
	time.Sleep(10 * time.Second)
}

func main() {

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
