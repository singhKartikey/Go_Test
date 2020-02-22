package main

import (
	"Go_Test/api"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	api.UpperCase("test")

	http.ListenAndServe(":80", nil)
}
