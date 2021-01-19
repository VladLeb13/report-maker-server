package controller

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	//TODO: context: config templates file
	path := "/home/worker/Studing/report-maker-server/src/server/templates/"

	tmpl, err := template.ParseFiles(path + "home.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {
		tmpl.Execute(ctx.Writer, nil)
	}

	return

}
