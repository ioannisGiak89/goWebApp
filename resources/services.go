package resources

import (
	"net/http"
)

// Services contains all of the shared services in the application.
type Services struct {
	// Handlers
	ViewHandler http.HandlerFunc
	EditHandler http.HandlerFunc
	SaveHandler http.HandlerFunc
}
