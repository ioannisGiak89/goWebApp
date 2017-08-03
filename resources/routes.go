package resources

import (
	"net/http"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

// BuildRoutes takes pre-configured handlers, and assigns them to routes.
func BuildRoutes(services Services) []Route {
	return []Route{
		// View
		{"/view/", services.ViewHandler},
		{"/edit/", services.EditHandler},
		{"/save/", services.SaveHandler},
	}
}
