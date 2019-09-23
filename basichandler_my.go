package main

import (
	"io"
	"net/http"
)

// hello world, the web server

func SendHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello guys!\n")
	w.Write([]byte("Hello guys with write!\n"))
}

func main() {
	http.HandleFunc("/", SendHello)
	http.ListenAndServe(":8000", nil)
}
