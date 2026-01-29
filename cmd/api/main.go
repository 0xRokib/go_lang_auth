package main

import (
	"log"
	"net/http"
	"time"

	"github.com/0xRokib/golang_auth/internal/httpserver"
)

func main() {
	router := httpserver.NewRouter()

	srv := &http.Server{
		Addr:              ":5000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Server Running on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
