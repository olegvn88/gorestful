package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/rint", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%d\n", rand.Int())
	})

	mux.HandleFunc("/rfloat", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%f\n", rand.Float64())
	})

	http.ListenAndServe(":8000", mux)
}
