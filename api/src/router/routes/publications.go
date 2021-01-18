package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Route{
	{
		URI:                   "/publications",
		Method:                http.MethodPost,
		Function:              controllers.CreatePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications",
		Method:                http.MethodGet,
		Function:              controllers.AllPublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{id}",
		Method:                http.MethodGet,
		Function:              controllers.FindPublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/publications",
		Method:                http.MethodGet,
		Function:              controllers.FindPublicationByUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{id}/like",
		Method:                http.MethodPost,
		Function:              controllers.LikePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{id}/unlike",
		Method:                http.MethodPost,
		Function:              controllers.UnlikePublication,
		RequireAuthentication: true,
	},
}
