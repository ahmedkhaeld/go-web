package main

import (
	"github.com/ahmedkhaeld/go-web/pkg/config"
	"github.com/ahmedkhaeld/go-web/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// Gracefully absorb panics and prints the stack trace
	mux.Use(middleware.Recoverer)

	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux

}
