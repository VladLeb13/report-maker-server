package writer

import (
	"log"

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

func fromRreiver(data *datalib.Report) {
	//TODO: write to db report

	log.Printf("%+v", data)
}

func fromApp(data model.TO_WR) {

}

func datalibToModel(report datalib.Report) (out model.TO_WR) {
	out.Workstation.Name = report.Software.OS.Name
	out.Workstation.Comment = "buh"
	out.Workstation.Allow_analysis = 0

	out.Workstation.Hardware = fill.Hardware(report)
	out.Workstation.Program_list = fill.Program(report)
	out.Workstation.Fault_tolerance = fill.Fault_tolerance(report)
	out.Workstation.Perfomance = fill.Perfomance(report)

	return
}
