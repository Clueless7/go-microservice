package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/foo" {
			w.Write([]byte("foo!"))

			return
		}
	}

	w.Write([]byte("Hello, World"))
}
