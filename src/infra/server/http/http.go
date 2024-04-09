package server

import (
	"encoding/json"
	"net/http"
)

type Http struct {
	Port string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	response := map[string]bool{"check": true}
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Contet-Type", "application/json")

	w.Write(jsonResponse)
}

func (h Http) Start() error {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":"+h.Port, nil)
	return err
}
