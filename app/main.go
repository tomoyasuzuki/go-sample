package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/tomoyasuzuki/go-sample/pkg/handlers"
	"log"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.RootHandler)
	mux.Get("/about", handlers.AboutHandler)

	return mux
}

func main() {
	sv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	log.Fatal(sv.ListenAndServe())
}
