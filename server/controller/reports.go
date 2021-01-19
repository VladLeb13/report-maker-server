package controller

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

func Reports(ctx *gin.Context) {
	//TODO: context: config templates file
	path := "/home/lebedev/Документы/srv/src/report-maker-server/server/templates/"

	tmpl, err := template.ParseFiles(path + "reports.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {
		tmpl.Execute(ctx.Writer, nil)
	}
}
