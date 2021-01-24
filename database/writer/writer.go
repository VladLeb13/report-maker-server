package writer

import (
	"database/sql"

	"report-maker-server/database/writer/fill"
	"report-maker-server/server/model"
	"report-maker-server/tools"

	"github.com/VladLeb13/report-maker-lib/datalib"
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
	data.Workstation.Name = report.Software.OS.Name
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

	db.Exec(`select *`)
}
