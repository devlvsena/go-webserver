package main

import (
	"fmt"
	server "lvsena/go-webserver/infra/server/http"
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
