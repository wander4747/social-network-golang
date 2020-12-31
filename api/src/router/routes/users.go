package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.Create,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.All,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		Function:              controllers.Find,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodPut,
		Function:              controllers.Update,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodDelete,
		Function:              controllers.Delete,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/follow",
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.UnFollowUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/followers",
		Method:                http.MethodGet,
		Function:              controllers.Followers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/following",
		Method:                http.MethodGet,
		Function:              controllers.Following,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/password",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePassword,
		RequireAuthentication: true,
	},
}
