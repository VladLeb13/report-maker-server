package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"

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

type Form_data struct {
	PC_Name            string
	CPU                string
	CPU_Freq           int
	CPU_Cores          int
	CPU_Thred          int
	RAM_Size           int
	RAM_Freq           int
	HDD_Type           int
	Perfomance_Cluster int
	FaultT_Cluster     int
}

const select_FORM_DATA = `SELECT  Workstation.Name AS "PC_Name"
							, CPU.Model AS "CPU"
							, CPU.Frequency AS "CPU_Freq"
							, CPU.Number_cores AS "CPU_Cores"
							, CPU.Number_threads AS "CPU_Thred"
							, RAM.Size AS "RAM_Size"
							, RAM.Frequency AS "RAM_Freq"
							, HDD.Type AS "HDD_Type"
							, Perfomance.Cluster AS "Perfomance_Cluster"
							, Fault_tolerance.Cluster AS "FaultT_Cluster"
						  FROM Workstation
						  	INNER JOIN Perfomance ON Perfomance.ID = Workstation.PerfomanceID
						  	INNER JOIN Fault_tolerance ON Fault_tolerance.ID = Workstation.Fault_toleranceID
							INNER JOIN Hardware ON Hardware.ID = Workstation.HardwareID
							INNER JOIN CPU_list ON CPU_list.ID =  Hardware.CPU_listID
							INNER JOIN CPU ON CPU.ID = CPU_list.CPUID
							INNER JOIN RAM_list ON RAM_list.ID =  Hardware.RAM_listID
							INNER JOIN RAM ON RAM.ID = RAM_list.RAMID
							INNER JOIN HDD_list ON HDD_list.ID =  Hardware.HDD_listID
							INNER JOIN HDD ON HDD.ID = HDD_list.HDDID
						  WHERE Name = $1
						  LIMIT 1`

//Searching - найденый комп
func Searching(ctx *gin.Context) {
	app_ctx := ctx.MustGet("app-context").(*tools.AppContex)
	cnf := app_ctx.Context.Value("config").(config.Config)
	db := app_ctx.Context.Value("database").(*sql.DB)

	pc_name := ctx.Request.FormValue("pc_name")

	fd := Form_data{}
	if pc_name != "" {
		err := db.QueryRow(select_FORM_DATA, pc_name).Scan(
			&fd.PC_Name,
			&fd.CPU,
			&fd.CPU_Freq,
			&fd.CPU_Cores,
			&fd.CPU_Thred,
			&fd.RAM_Size,
			&fd.RAM_Freq,
			&fd.HDD_Type,
			&fd.Perfomance_Cluster,
			&fd.FaultT_Cluster,
		)
		if err != nil {
			log.Println("Error of query select_FORM_DATA: ", err)
		}
	}

	path := cnf.Template_path
	tmpl, err := template.ParseFiles(path + "searching.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {

		tmpl.Execute(ctx.Writer, fd)
	}

}
