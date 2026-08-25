package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	IdleWait int
)

type Specification struct {
	IdleWait  int    `default:"10"`
	KeepAlive bool   `default:"true"`
	Addr      string `default:":8080"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Connection received")

	if _, err := fmt.Fprintln(w, "hello"); err != nil {
		log.Printf("failed to write response: %v", err)
	}
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
	fmt.Printf("Keep Alives: %t\n", s.KeepAlive)
	fmt.Printf("Addr: %s\n", s.Addr)

	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: s.Addr,
	}
	server.SetKeepAlivesEnabled(s.KeepAlive)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
