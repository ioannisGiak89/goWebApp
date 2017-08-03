package resources

import (
	"html/template"
)

func LoadTemplates() *template.Template {
	return template.Must(template.ParseGlob("view/*/*"))
	//return template.Must(template.ParseFiles(strings.Join(getTemplates(), ",")))
}

func getTemplates() []string {
	return []string{
		"view/edit/edit.html",
		"view/save/save.html",
		"view/view/view.html",
	}
}
