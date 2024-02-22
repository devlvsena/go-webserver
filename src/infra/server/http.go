package server

import (
	"fmt"
	"net/http"
)

type Http struct {
	Port string
}

func (h Http) Start() {
	err := http.ListenAndServe(":"+h.Port, nil)
	if err != nil {
		fmt.Print(err)
	}
}
