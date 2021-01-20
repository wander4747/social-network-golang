package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates = Load templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// RunTemplate =  execute template html
func RunTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
