package main

import (
	"github.com/DaminAlaskarovTC/bookings/pkg/config"
	"github.com/DaminAlaskarovTC/bookings/pkg/handlers"
	"net/http"

	// go mod tidy om ervoor te zorgen dat deze 2 chi packages worden gevoegd aan demod file en dat garbage wordt verwijdert uit de .mod file
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
