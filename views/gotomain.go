// views/views.go
package views

import (
	branch "po/views/Branch"
	home "po/views/Home"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("views/login/login.html"))
	tmpl.Execute(c.Writer, nil)
}

func WD(router *gin.Engine) {
	router.GET("/", Welcome)
	router.GET("/Home", home.Home)
	router.GET("/:branch", branch.Branch) // Use :branch as a parameter
	router.Static("/views", "./views")
}
