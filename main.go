package main

import (
	"github.com/paulomujuru/functionaloptionspatterns/server"
	"time"
)

func main() {
	srv := server.New(
		server.WithHost("localhost"),
		server.WithMaxConn(100),
		server.WithPort(8080),
		server.WithTime(time.Minute),
	)
	srv.Start()
	time.Sleep(time.Minute)
	srv.Stop()
}
