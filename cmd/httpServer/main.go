package main

import (
	"flag"
	"fmt"
	"net/http"
)

var version = flag.String("VERSION", "1.0", "config filepath")

func main() {
	http.HandleFunc("/health", healthcheck)
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func index(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Print(header)
	fmt.Fprintln(w, "Hello World")
}
