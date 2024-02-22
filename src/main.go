package main

import (
	"lvsena/go-webserver/server"
)

func main() {
	s := &server.Http{
		Port: "8080",
	}

	s.Start()
}
