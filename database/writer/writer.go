package writer

import (
	"github.com/google/uuid"
	"log"
	"report-maker-server/database/writer/fill"
	"report-maker-server/server/model"
	"strconv"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

func AddRecord(data interface{}) {
	switch data.(type) {
	case *datalib.Report:
		fromRreiver(data.(*datalib.Report))
	case model.TO_WR:
		fromApp(data)
	}
}

func fromRreiver(data *datalib.Report) {
	//TODO: write to db report

	log.Printf("%+v", data)
}

func fromApp(data model.TO_WR) {

}

func datalibToModel(in datalib.Report) (out model.TO_WR) {

	var cpu_list []model.CPU
	for _, v := range in.Hardware.CPUs {
		var cpu model.CPU

		cpu.Manufacturer = v.Manufacturer
		cpu.Model = v.Name
		cpu.Frequency = int(v.MaxClockSpeed)
		cpu.Number_cores = int(v.NumberOfCores)
		cpu.Number_threads = int(v.ThreadCount)

		cpu_list = append(cpu_list, cpu)
	}

	var ram_list []model.RAM
	for _, v := range in.Hardware.RAMs {
		var ram model.RAM

		ram.Manufacturer = v.Manufacturer
		ram.Frequency = int(v.Speed)
		ram.Serial_number = v.PartNumber
		ram.Size = int(v.Capacity)

		ram_list = append(ram_list, ram)
	}

	var hdd_list []model.HDD
	for _, v := range in.Hardware.HDDs {
		var hdd model.HDD

		hdd.Size = int(v.Size)
		hdd.Model = v.Model

		//TODO: match type in name of hdd if name contains SSD then type = 2
		hdd.Type = 1

		hdd_list = append(hdd_list, hdd)
	}

	var programs_list []model.Program
	for _, v := range in.Software.Programs {
		var program model.Program

		program.Manufacturer = v.Vendor
		program.Name = v.Name
		program.Install_on = v.InstallDate
		program.Version = v.Version

		programs_list = append(programs_list, program)
	}

	var perfomance model.Perfomance

	if len(in.Perfomance.HDD) != 0 {
		perfomance.HDD = strconv.Itoa(int(in.Perfomance.HDD[0].PercentDiskTime))
	}

	if len(in.Perfomance.RAM) != 0 {
		perfomance.RAM = strconv.Itoa(int(in.Perfomance.RAM[0].AvailableMBytes))
	}

	if len(in.Perfomance.CPU) != 0 {
		perfomance.CPU = strconv.Itoa(int(in.Perfomance.CPU[0].PercentProcessorUtility))
	}

	perfomance.Cluster = 0
	perfomance.ID = uuid.New().String()

	var flt model.Fault_tolerance

	flt.Cluster = 0
	flt.Number_of_error = len(in.Events.List)
	flt.Backup = 1
	flt.Commissioning_date = "2020-06-22T00:00:00Z"
	flt.ID = uuid.New().String()

	//fill

	out.Workstation.Name = in.Software.OS.Name
	out.Workstation.Comment = "buh"
	out.Workstation.Allow_analysis = 0

	out.Fault_tolerance = flt
	out.Perfomance = perfomance

	out.Program_list.ID = uuid.New().String()

	for _, v := range programs_list {
		uuid := uuid.New().String()

		out.Program_list.ProgramID = append(out.Program_list.ProgramID, uuid)
		v.ID = uuid
	}

	board_id := uuid.New().String()
	out.Hardware.MatherboardID = board_id
	out.Matherboard.ID = board_id
	out.Matherboard.Name = in.Hardware.Board.Manufacturer
	out.Matherboard.Product = in.Hardware.Board.Product
	out.Matherboard.Model = in.Hardware.Board.Version

	cpu_list_id := uuid.New().String()
	out.Hardware.CPU_listID = cpu_list_id
	out.CPU_list.ID = cpu_list_id

	for _, v := range cpu_list {
		id := uuid.New().String()
		out.CPU_list.CPUID = append(out.CPU_list.CPUID, id)
		v.ID = id
	}
	out.CPUs = cpu_list

	ram_list_id := uuid.New().String()
	out.Hardware.RAM_listID = ram_list_id
	out.RAM_list.ID = ram_list_id

	for _, v := range ram_list {
		id := uuid.New().String()

		out.RAM_list.RAMID = append(out.RAM_list.RAMID, id)
		v.ID = id
	}
	out.RAMs = ram_list

	hdd_list_id := uuid.New().String()
	out.Hardware.HDD_listID = hdd_list_id
	out.HDD_list.ID = hdd_list_id

	for _, v := range hdd_list {
		id := uuid.New().String()

		out.HDD_list.HDDID = append(out.HDD_list.HDDID, id)
		v.ID = id

	}
	out.HDDs = hdd_list

	return
}
