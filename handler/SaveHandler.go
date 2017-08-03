package handler

import (
	"net/http"

	"github.com/ioannisGiak89/goWebApp/model"
)

type SaveHandler struct {
	pageHandler *PageHandler
}

func NewSaveHandler(ph *PageHandler) *SaveHandler {
	return &SaveHandler{
		pageHandler: ph,
	}
}

func (sh *SaveHandler) SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	p := &model.Page{Title: title, Body: []byte(r.FormValue("body"))}
	p.Save()
	sh.pageHandler.RenderTemplate(w, "save", p)
}
