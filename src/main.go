package main

import (
	"fmt"
	"lvsena/go-webserver/infra/server"
)

func main() {
	port := "8080"

	s := &server.Http{
		Port: port,
	}

	fmt.Println("Starting Http Server on port", port)

	err := s.Start()
	if err != nil {
		fmt.Println(err)
	}
}
