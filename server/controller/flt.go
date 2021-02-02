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

type data_For_Analysis struct {
	work_id string
	hard_id string
	perf_id string
	flt_id  string
}

type Form_data_flt struct {
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

const get_Data = ` 	SELECT  Workstation.ID AS "work_id"   
						, HardwareID AS "hard_id"
 						, PerfomanceID AS "perf_id"
 						, Fault_toleranceID AS "flt_id" 				
					FROM Workstation
								WHERE Allow_analysis = 1`

const select_FORM_DATA_FLT = `SELECT  Workstation.Name AS "PC_Name"
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
						  WHERE Workstation.ID = $1
						  LIMIT 1`

//Searching - найденый комп
func Flt(ctx *gin.Context) {
	app_ctx := ctx.MustGet("app-context").(*tools.AppContex)
	cnf := app_ctx.Context.Value("config").(config.Config)
	db := app_ctx.Context.Value("database").(*sql.DB)

	rows, err := db.Query(get_Data)
	if err != nil {
		log.Println("Error in query get_Data ", err)
	}

	var resp []data_For_Analysis
	for rows.Next() {
		var v data_For_Analysis
		rows.Scan(&v.work_id, &v.hard_id, &v.perf_id, &v.flt_id)

		resp = append(resp, v)
	}
	rows.Close()

	var data_form []Form_data_flt
	for _, v := range resp {
		fd := Form_data_flt{}
		err := db.QueryRow(select_FORM_DATA_FLT, v.work_id).Scan(
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
			log.Println("Error of query select_FORM_DATA_FLT: ", err)
		}
		data_form = append(data_form, fd)
	}

	type f struct {
		D   []Form_data_flt
		All float64
	}

	frm := f{}
	frm.D = data_form

	length := len(data_form)

	var buff float64
	for _, v := range data_form {
		buff = float64(int(buff) + v.FaultT_Cluster)
	}

	frm.All = buff / float64(length)

	path := cnf.Template_path
	tmpl, err := template.ParseFiles(path + "flt-report.html")
	if err != nil {
		fmt.Println("error parce template")
	} else {
		tmpl.Execute(ctx.Writer, frm)
	}

}
