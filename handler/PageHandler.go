package handler

import (
	"html/template"
	"net/http"

	"github.com/ioannisGiak89/goWebApp/model"
)

type PageHandler struct {
	templates *template.Template
}

func NewPageHandler(templates *template.Template) *PageHandler {
	return &PageHandler{
		templates: templates,
	}
}

func (ph *PageHandler) getTemplates() *template.Template {
	return ph.templates
}

func (ph *PageHandler) RenderTemplate(w http.ResponseWriter, tmpl string, p *model.Page) {
	err := ph.getTemplates().ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
