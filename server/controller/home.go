package controller

import (
	"fmt"
	"html/template"

	"report-maker-server/config"
	"report-maker-server/tools"

	"github.com/gin-gonic/gin"
)

type content struct {
	Message string
}

func Home(ctx *gin.Context) {
	app_ctx := ctx.MustGet("app-context").(*tools.AppContex)
	cnf := app_ctx.Context.Value("config").(config.Config)

	path := cnf.Template_path

	tmpl, err := template.ParseFiles(path + "home.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {
		tmpl.Execute(ctx.Writer, nil)
	}

	return

}
