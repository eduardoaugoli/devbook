package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Carregar insere o templete hmlt na variavel template
func CarregarTempletes() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))

}

// Renderiza uma pagina html
func ExecutarTemplete(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)

}
