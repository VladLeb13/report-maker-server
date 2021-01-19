package writer

import (
	"log"

	"report-maker-server/database/writer/fill"
	"report-maker-server/server/model"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/google/uuid"
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

func datalibToModel(in datalib.Report) (out model.TO_WR) {

	cpu_list_id := uuid.New().String()
	ram_list_id := uuid.New().String()
	hdd_list_id := uuid.New().String()

	out.Hardware.CPU_listID = cpu_list_id
	out.CPUs, out.CPU_list = fill.CPU(in, cpu_list_id)

	out.Hardware.RAM_listID = ram_list_id
	out.RAMs, out.RAM_list = fill.RAM(in, ram_list_id)

	out.Hardware.HDD_listID = hdd_list_id
	out.HDDs, out.HDD_list = fill.HDD(in, hdd_list_id)

	out.Matherboard = fill.Board(in.Hardware)
	out.Hardware.MatherboardID = out.Matherboard.ID

	out.Workstation.Name = in.Software.OS.Name
	out.Workstation.Comment = "buh"
	out.Workstation.Allow_analysis = 0
	out.Programs, out.Program_list = fill.Program(in)

	out.Fault_tolerance = fill.Fault_tolerance(in.Events)
	out.Perfomance = fill.Perfomance(in.Perfomance)

	return
}
