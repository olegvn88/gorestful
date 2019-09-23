package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

// CustomServeMux is a struct which can be a multiplexer
type CustomServeMux struct {
}

// This is the function handler to be overridden
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Your random number is: %f\n", rand.Float64(), )
	f := rand.Float64()

	res := strconv.FormatFloat(f, 'g', 1, 64)
	w.Write([]byte("Your random number is: " + res + "\n"))
}

func main() {
	// Any struct that has serveHTTP function can be a multiplexer
	mux := &CustomServeMux{}
	http.ListenAndServe(":8000", mux)
}
