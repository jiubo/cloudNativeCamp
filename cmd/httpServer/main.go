package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/health", healthcheck)
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
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
	clientip := getCurrentIP(r)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientip)
}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
