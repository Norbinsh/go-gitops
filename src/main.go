package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", goGitops)
	http.ListenAndServe(":80", nil)
}

func goGitops(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from: %s", getHostname())
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "not available"
	}
	return hostname
}
