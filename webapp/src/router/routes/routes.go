package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route = Repressents all routes
type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

// Configure = puts all routes inside the router
func Configure(router *mux.Router) *mux.Router {
	routes := routesLogin
	routes = append(routes, routesUser...)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
