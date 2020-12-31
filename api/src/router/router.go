package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate = return a router with configures routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
