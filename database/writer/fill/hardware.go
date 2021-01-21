package fill

import (
	"report-maker-server/server/model"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/google/uuid"
)

func Hardware(report datalib.Report) (hardware model.Hardware) {
	hardware.CPU_list = cpu(report)
	hardware.RAM_list = ram(report)
	hardware.HDD_list = hdd(report)
	hardware.Matherboard = board(report)

	return
}

func cpu(report datalib.Report) (list model.CPU_list) {
	list.ID = uuid.New().String()

	for _, v := range report.Hardware.CPUs {
		var cpu model.CPU

		cpu.ID = uuid.New().String()
		cpu.Manufacturer = v.Manufacturer
		cpu.Model = v.Name
		cpu.Frequency = int(v.MaxClockSpeed)
		cpu.Number_cores = int(v.NumberOfCores)
		cpu.Number_threads = int(v.ThreadCount)

		list.CPUs = append(list.CPUs, cpu)
	}

	return
}

func ram(report datalib.Report) (list model.RAM_list) {
	list.ID = uuid.New().String()

	for _, v := range report.Hardware.RAMs {
		var ram model.RAM

		ram.ID = uuid.New().String()
		ram.Manufacturer = v.Manufacturer
		ram.Frequency = int(v.Speed)
		ram.Serial_number = v.PartNumber
		ram.Size = int(v.Capacity)

		list.RAMs = append(list.RAMs, ram)
	}

	return
}

func hdd(in datalib.Report) (list model.HDD_list) {
	list.ID = uuid.New().String()

	for _, v := range in.Hardware.HDDs {
		var hdd model.HDD

		hdd.ID = uuid.New().String()
		hdd.Size = int(v.Size)
		hdd.Model = v.Model

		//TODO: match type in name of hdd if name contains SSD then type = 2
		hdd.Type = 1

		list.HDDs = append(list.HDDs, hdd)
	}

	return
}

func board(report datalib.Report) (board model.Matherboard) {
	board.ID = uuid.New().String()
	board.Name = report.Hardware.Board.Manufacturer
	board.Product = report.Hardware.Board.Product
	board.Model = report.Hardware.Board.Version

	return
}
