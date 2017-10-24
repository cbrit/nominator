package main

import (
	"io"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Success!\n")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
