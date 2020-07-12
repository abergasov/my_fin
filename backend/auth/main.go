package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	buildTime string = "_dev"
	buildHash string = "_dev"
)

func main() {
	log.Println("Auth app. App build time:", buildTime)
	log.Println("Auth app. App build hash:", buildHash)

	http.HandleFunc("/api/auth/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":8090", nil)
}
