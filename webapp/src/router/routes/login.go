package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesLogin = []Route{
	{
		URI:                   "/",
		Method:                http.MethodGet,
		Function:              controllers.LoadViewLogin,
		RequireAuthentication: false,
	},
	{
		URI:                   "/login",
		Method:                http.MethodGet,
		Function:              controllers.LoadViewLogin,
		RequireAuthentication: false,
	},
	{
		URI:                   "/login",
		Method:                http.MethodPost,
		Function:              controllers.DoLogin,
		RequireAuthentication: false,
	},
}
