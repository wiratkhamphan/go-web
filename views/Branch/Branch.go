package branch

import (
	"fmt"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Branch(c *gin.Context) {
	selectedBranch := c.Param("branch")                               // Get the selected branch with the file extension
	tmplPath := fmt.Sprintf("views/Branch/Branch/%s", selectedBranch) // No need to add ".html" here

	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(c.Writer, nil)
}
