package server

import (
	"net/http"
)

type Http struct {
	Port string
}

func (h Http) Start() error {
	err := http.ListenAndServe(":"+h.Port, nil)
	return err
}
