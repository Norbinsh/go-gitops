package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", goGitops)
	http.ListenAndServe(":80", nil)
}

func goGitops(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UserAgent Is: %s", r.UserAgent())
}
