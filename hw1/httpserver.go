package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	version := os.Getenv("VERSION")
	if version != "" {
		fmt.Fprintf(w, `Header["VERSION"]=%q`, version)
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("IP address: ", ip, "Status code: ", http.StatusForbidden)
		log.Fatalf("error getting IP address %v", err)
	}
	log.Println("IP address: ", ip, "Status code: ", http.StatusOK)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
