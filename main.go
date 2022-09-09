package main

import (
	"app-engine-golang-test/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	h := handler.NewHttpHandler("ca-nakaharakodai-test-358502")
	http.HandleFunc("/debug", h.Debug)
	http.HandleFunc("/info", h.Info)
	http.HandleFunc("/error", h.Error)
	http.HandleFunc("/critical", h.Critical)
	http.HandleFunc("/warning", h.Warning)
	http.HandleFunc("/", h.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("PORT: %v", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
