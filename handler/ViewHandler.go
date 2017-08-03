package handler

import (
	"net/http"

	"github.com/ioannisGiak89/goWebApp/model"
)

type ViewHandler struct {
	pageHandler *PageHandler
}

func NewViewHandler(ph *PageHandler) *ViewHandler {
	return &ViewHandler{
		pageHandler: ph,
	}
}

func (vh *ViewHandler) ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := model.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	vh.pageHandler.RenderTemplate(w, "view", p)
}
