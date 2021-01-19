package fill

import (
	"report-maker-server/server/model"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/google/uuid"
)

func CPU(in datalib.Report, cpu_list_id string) (cpu_list []model.CPU, cpu model.CPU_list) {
	for _, v := range in.Hardware.CPUs {
		var cpu model.CPU

		cpu.Manufacturer = v.Manufacturer
		cpu.Model = v.Name
		cpu.Frequency = int(v.MaxClockSpeed)
		cpu.Number_cores = int(v.NumberOfCores)
		cpu.Number_threads = int(v.ThreadCount)

		cpu_list = append(cpu_list, cpu)
	}

	cpu.ID = cpu_list_id

	for _, v := range cpu_list {
		id := uuid.New().String()
		cpu.CPUID = append(cpu.CPUID, id)
		v.ID = id
	}

	return
}

func RAM(in datalib.Report, ram_list_id string) (ram_list []model.RAM, ram model.RAM_list) {
	for _, v := range in.Hardware.RAMs {
		var ram model.RAM

		ram.Manufacturer = v.Manufacturer
		ram.Frequency = int(v.Speed)
		ram.Serial_number = v.PartNumber
		ram.Size = int(v.Capacity)

		ram_list = append(ram_list, ram)
	}

	ram.ID = ram_list_id

	for _, v := range ram_list {
		id := uuid.New().String()
		ram.RAMID = append(ram.RAMID, id)
		v.ID = id
	}

	return
}

func HDD(in datalib.Report, hdd_list_id string) (hdd_list []model.HDD, hdd model.HDD_list) {
	for _, v := range in.Hardware.HDDs {
		var hdd model.HDD

		hdd.Size = int(v.Size)
		hdd.Model = v.Model

		//TODO: match type in name of hdd if name contains SSD then type = 2
		hdd.Type = 1

		hdd_list = append(hdd_list, hdd)
	}

	hdd.ID = hdd_list_id
	for _, v := range hdd_list {
		id := uuid.New().String()

		hdd.HDDID = append(hdd.HDDID, id)
		v.ID = id

	}

	return
}

func Board(hardware datalib.Hardware) (board model.Matherboard) {
	board.ID = uuid.New().String()
	board.Name = hardware.Board.Manufacturer
	board.Product = hardware.Board.Product
	board.Model = hardware.Board.Version

	return
}
