package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// carrega template insere os tamplates html na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// Executar renderiza uma pagina html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
