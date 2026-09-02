package server

import (
	"net/http"

	"api-email-auch/internal/verify"
)

func NewMux(handler *verify.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /send", handler.SendHandler)
	mux.HandleFunc("GET /verify/{hash}", handler.VerifyHandler)

	return mux
}
