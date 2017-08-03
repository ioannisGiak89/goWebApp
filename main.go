package main

import (
	"fmt"
	"net/http"

	"github.com/ioannisGiak89/goWebApp/handler"
	"github.com/ioannisGiak89/goWebApp/resources"
)

func main() {
	templates := resources.LoadTemplates()
	pageHandler := handler.NewPageHandler(templates)

	services := resources.Services{
		ViewHandler: handler.NewViewHandler(pageHandler).ViewHandler,
		EditHandler: handler.NewEditHandler(pageHandler).EditHandler,
		SaveHandler: handler.NewSaveHandler(pageHandler).SaveHandler,
	}

	for _, route := range resources.BuildRoutes(services) {
		http.HandleFunc(route.Path, route.Handler)
	}

	fmt.Println("Listening on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
