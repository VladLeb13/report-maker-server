package writer

import (
	"database/sql"
	"log"

	"report-maker-server/database/writer/fill"
	"report-maker-server/database/writer/update"
	"report-maker-server/database/writer/write"
	"report-maker-server/server/model"
	"report-maker-server/tools"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/google/uuid"
)

func AddRecord(ctx *tools.AppContex, data interface{}) {
	switch data.(type) {
	case *datalib.Report:
		fromRreiver(ctx, data.(*datalib.Report))
	case model.TO_WR:
		fromApp(ctx, data.(model.TO_WR))
	}
}

func fromRreiver(ctx *tools.AppContex, report *datalib.Report) {
	data := datalibToModel(*report)

	fromApp(ctx, data)
}

func datalibToModel(report datalib.Report) (data model.TO_WR) {
	data.Workstation.ID = uuid.New().String()
	data.Workstation.Name = report.Software.OS.UserName
	data.Workstation.Comment = "buh"
	data.Workstation.Allow_analysis = 0

	data.Workstation.Hardware = fill.Hardware(report)
	data.Workstation.Program_list = fill.Program(report)
	data.Workstation.Fault_tolerance = fill.Fault_tolerance(report)
	data.Workstation.Perfomance = fill.Perfomance(report)

	return
}

func fromApp(ctx *tools.AppContex, data model.TO_WR) {
	db := ctx.Context.Value("database").(*sql.DB)

	if data.Workstation.ID == "" || data.Workstation.Name == "" ||
		data.Workstation.Hardware.ID == "" ||
		data.Workstation.Program_list.ID == "" ||
		data.Workstation.Perfomance.ID == "" ||
		data.Workstation.Fault_tolerance.ID == "" {
		return
	}

	if check_record(db, data.Workstation.Name) {
		err := update.LoadLink(db, &data)
		if err != nil {
			log.Println("Fail update Workstation ", data.Workstation.Name, " error: ", err)
			return
		}
		update.Workstation(db, data)
		update.Perfomance(db, data)
		update.Fault_tolerance(db, data)
		update.Hardware(db, data)
		update.Software(db, data)
		return
	}

	write.Workstation(db, data)
	write.Perfomance(db, data)
	write.Fault_tolerance(db, data)
	write.Hardware(db, data)
	write.Software(db, data)
}

const check_WORKSATION = `SELECT ID
						   FROM Workstation
						   WHERE Name = $1`

func check_record(db *sql.DB, workstation_name string) (avalible bool) {
	var workstation_id string
	err := db.QueryRow(check_WORKSATION, workstation_name).Scan(&workstation_id)
	if err != nil {
		log.Println("Error query of check_WORKSATION: ", err)
	}

	if workstation_id == "" {
		return
	}

	log.Println("Workstation ", workstation_name, " avalible")

	avalible = true
	return
}
