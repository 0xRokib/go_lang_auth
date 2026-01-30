package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/0xRokib/golang_auth/internal/app"
	"github.com/0xRokib/golang_auth/internal/httpserver"
)

func main() {
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Println(err)
		}
	}()

	router := httpserver.NewRouter()

	srv := &http.Server{
		Addr:              ":5000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Server Running on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
