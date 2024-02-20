package main

import (
	"fmt"
	"net/http"
)

func main() {

}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!<h1>")
}
