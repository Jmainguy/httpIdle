package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/kelseyhightower/envconfig"
	"os"
	"time"
)

var (
	IdleWait int
)
type Specification struct {
	IdleWait int `default:"10"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Connection received")

	fmt.Fprintf(w, "hello\n")
	time.Sleep(time.Duration(IdleWait) * time.Second)
}

func main() {
	var s Specification
	err := envconfig.Process("httpidle", &s)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	IdleWait = s.IdleWait

	fmt.Printf("Idle Wait: %d\n", IdleWait)

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
