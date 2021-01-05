package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func Reports(w http.ResponseWriter, r *http.Request) {
	//TODO: context: config templates file
	path := "/home/worker/Studing/report-maker-server/src/server/templates/"

	tmpl, err := template.ParseFiles(path + "reports.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {
		tmpl.Execute(w, nil)
	}
}
