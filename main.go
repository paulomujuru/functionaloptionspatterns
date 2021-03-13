package main

import (
	"github.com/paulomujuru/functionaloptionspatterns/server"
	"time"
)

func main() {
	srv := server.New("localhost", 8080, time.Minute)
	srv.Start()
	time.Sleep(time.Second)
	srv.Stop()
}
