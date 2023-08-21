package home

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func Home(h *gin.Context) {
	tmpl := template.Must(template.ParseFiles("views/Home/html/home.html"))
	tmpl.Execute(h.Writer, nil)
}
