package controller

import (
	"fmt"
	"html/template"

	"report-maker-server/config"
	"report-maker-server/tools"

	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context) {
	app_ctx := ctx.MustGet("app-context").(*tools.AppContex)
	cnf := app_ctx.Context.Value("config").(config.Config)

	path := cnf.Template_path

	tmpl, err := template.ParseFiles(path + "search.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {
		tmpl.Execute(ctx.Writer, nil)
	}

}

//Searching - найденый комп
type Form_data struct {
	PC_Name            string
	CPU                string
	CPU_Freq           string
	CPU_Cores          string
	CPU_Thred          string
	RAM_Size           string
	RAM_Freq           string
	HDD_Type           string
	Perfomance_Cluster string
	FaultT_Cluster     string
}

func Searching(ctx *gin.Context) {
	app_ctx := ctx.MustGet("app-context").(*tools.AppContex)
	cnf := app_ctx.Context.Value("config").(config.Config)

	//r := ctx.Request

	// pc_name := r.FormValue("pc_name")
	// ctx.Redirect(http.StatusFound, "/home")

	fd := Form_data{
		PC_Name:            "PC495",
		CPU:                "core2duo",
		CPU_Freq:           "2777",
		CPU_Cores:          "4",
		CPU_Thred:          "4",
		RAM_Size:           "8",
		RAM_Freq:           "1333",
		HDD_Type:           "ssd",
		Perfomance_Cluster: "5",
		FaultT_Cluster:     "4",
	}

	path := cnf.Template_path
	tmpl, err := template.ParseFiles(path + "searching.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {
		tmpl.Execute(ctx.Writer, fd)
	}

}
