package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/health", healthcheck)
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200"))
}

func index(w http.ResponseWriter, r *http.Request) {
	rHeader := r.Header
	for k, v := range rHeader {
		for _, value := range v {
			w.Header().Set(k, value)
		}
	}
	fmt.Println("VERSION", os.Getenv("VERSION"))
	if os.Getenv("VERSION") != "" {
		w.Header().Set("VERSION", os.Getenv("VERSION"))
	}
	fmt.Fprintln(w, "Hello World")
}
