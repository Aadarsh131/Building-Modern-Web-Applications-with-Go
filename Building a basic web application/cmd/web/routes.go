package main

import (
	"net/http"

	"github.com/Aadarsh131/Building-Modern-Web-Applications-with-Go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {

	//using the package "github.com/bmizerany/pat"
	// multiplexer := pat.New()

	// multiplexer.Get("/", http.HandlerFunc(handlers.Home))
	// multiplexer.Get("/About", http.HandlerFunc(handlers.About))

	//using the "chi" package
	multiplexer := chi.NewRouter()

	multiplexer.Use(middleware.Recoverer) //chi middleware
	multiplexer.Use(WriteToConsole)       //my custor middleware
	multiplexer.Use(CSRFprotection)

	multiplexer.Get("/", handlers.Home)
	multiplexer.Get("/About", handlers.About)
	return multiplexer
}
