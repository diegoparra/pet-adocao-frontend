// Package router provides route.
package router

import (
	"github.com/diegoparra/pet-adocao-frontend/src/router/routes"
	"github.com/gorilla/mux"
)

// Generate returns a router with all the routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
