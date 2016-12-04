package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", origin)
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}

func origin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World by Go!")
}