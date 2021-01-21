package writer

import (
	"report-maker-server/database/writer/fill"
	"report-maker-server/server/model"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

func AddRecord(data interface{}) {
	switch data.(type) {
	case *datalib.Report:
		fromRreiver(data.(*datalib.Report))
	case model.TO_WR:
		fromApp(data.(model.TO_WR))
	}
}

func fromRreiver(report *datalib.Report) {
	data := datalibToModel(*report)

	fromApp(data)
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

func fromApp(data model.TO_WR) {
	//TODO: write to db report
}
