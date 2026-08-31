package main

import (
	"log"
	"net/http"

	"api-email-auch/internal/config"
	"api-email-auch/internal/server"
	"api-email-auch/internal/verify"
)

func main() {
	cfg := config.LoadConfig()

	handler := verify.NewHandler(&cfg)
	mux := server.NewMux(handler)

	log.Println("server started on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
