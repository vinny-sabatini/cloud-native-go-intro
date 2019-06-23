package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/vinny-sabatini/cloud-native-go-intro/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)

	http.HandleFunc("/api/books", api.BooksHandleFunc)

	http.ListenAndServe(GetServePort(), nil)
}

// GetServePort returns environment variable SERVE_PORT if set otherwise, 8080 is returned
func GetServePort() string {
	port := os.Getenv("SERVE_PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

// index is the root entrypoint to the http server
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I'm only a hello world now, but I'll be a real program one day")
}

// echo prints the value of the URI parameter "message"
func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
