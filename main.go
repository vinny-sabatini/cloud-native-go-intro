package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(GetServePort(), nil)
}

// GetServePort returns environment variable SERVE_PORT if set otherwise, 8080 is returned
func GetServePort() string {
	//TODO: do this as a command line option instead
	port := os.Getenv("SERVE_PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I'm only a hello world now, but I'll be a real program one day")
}