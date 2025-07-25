package handlers

import (
	"github.com/JamesDuf/go-api-tutorial/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux) { //capital H tells the compiler this function can be imported in other packages, //lowercase means private
	//Global Middleware //Applied all the time to any endpoint we make
	r.Use(chimiddle.StripSlashes) //removes trailing slashes from the URL

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization) //apply auth middleware to all routes under /account
		router.Get("/coins", getCoinBalance) //GET request to /account/coins will call getCoinBalance function
	})

}
