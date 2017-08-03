package handler

import (
	"net/http"

	"github.com/ioannisGiak89/goWebApp/model"
)

type EditHandler struct {
	pageHandler *PageHandler
}

func NewEditHandler(ph *PageHandler) *EditHandler {
	return &EditHandler{
		pageHandler: ph,
	}
}

func (eh *EditHandler) EditHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	title := url[len("/edit/"):]
	p, err := model.LoadPage(url)
	if err != nil {
		p = &model.Page{Title: title}
	}
	eh.pageHandler.RenderTemplate(w, "edit", p)
}
