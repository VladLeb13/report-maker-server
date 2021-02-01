package write

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

const hardware_INSERT_HARDWARE = `INSERT INTO Hardware(ID)
								VALUES($1)`

//hardware-board
const (
	hardware_UPDATE_HARDWARE_SET_BOARD = `UPDATE Hardware
								SET MatherboardID = $1
								WHERE ID = $2`

	matherboard_INSERT_BOARD = `INSERT INTO Matherboard(ID, Name, Model, Product)
								VALUES($1, $2, $3, $4)`
)

//hardware-cpu
const (
	hardware_UPDATE_HARDWARE_SET_CPU_LIST = `UPDATE Hardware
								SET CPU_listID = $1
								WHERE ID = $2`

	cpu_list_INSERT_CPU_LIST = `INSERT INTO CPU_list(ID)
								VALUES($1)`

	cpu_list_UPDATE_CPU_list_SET_CPUID = `UPDATE CPU_list
								SET CPUID = $1
								WHERE ID = $2`

	cpu_INSERT_CPU = `INSERT INTO CPU(ID, Model, Manufacturer, Frequency, Number_cores, Number_threads)
								VALUES($1, $2, $3, $4, $5, $6)`
)

//hardware-ram
const (
	hardware_UPDATE_HARDWARE_SET_RAM_LIST = `UPDATE Hardware
								SET RAM_listID = $1
								WHERE ID = $2`

	ram_list_INSERT_RAM_LIST = `INSERT INTO RAM_listID(ID)
								VALUES($1)`

	ram_list_UPDATE_RAM_list_SET_RAMID = `UPDATE RAM_list
								SET RAMID = $1
								WHERE ID = $2`

	ram_INSERT_RAM = `INSERT INTO RAM(ID, Manufacturer,Size, Frequency, Serial_number)
								VALUES($1, $2, $3, $4, $5)`
)

//hardware-hdd
const (
	hardware_UPDATE_HARDWARE_SET_HDD_LIST = `UPDATE Hardware
								SET HDD_list = $1
								WHERE ID = $2`

	hdd_list_INSERT_HDD_LIST = `INSERT INTO HDD_list(ID)
								VALUES($1)`

	hdd_list_UPDATE_HDD_list_SET_HDDID = `UPDATE HDD_list
								SET HDDID = $1
								WHERE ID = $2`

	hdd_INSERT_HDD = `INSERT INTO HDD((ID, Model,Size, Type)
								VALUES($1, $2, $3, $4)`
)

func Hardware(db *sql.DB, data model.TO_WR) {

	_, err := db.Exec(hardware_INSERT_HARDWARE, data.Workstation.Hardware.ID)
	if err != nil {
		log.Println("Error write: hardware_INSERT_HARDWARE ", err)
	}

	if data.Workstation.Hardware.Matherboard.ID != "" {
		_, err := db.Exec(hardware_UPDATE_HARDWARE_SET_BOARD, data.Workstation.Hardware.Matherboard.ID, data.Workstation.Hardware.ID)
		if err != nil {
			log.Println("Error write: hardware_UPDATE_HARDWARE_SET_BOARD ", err)
		}

		_, err = db.Exec(matherboard_INSERT_BOARD,
			data.Workstation.Hardware.Matherboard.ID,
			data.Workstation.Hardware.Matherboard.Name,
			data.Workstation.Hardware.Matherboard.Model,
			data.Workstation.Hardware.Matherboard.Product,
		)
		if err != nil {
			log.Println("Error write: matherboard_INSERT_BOARD ", err)
		}

	}

	if data.Workstation.Hardware.CPU_list.ID != "" {
		_, err = db.Exec(hardware_UPDATE_HARDWARE_SET_CPU_LIST, data.Workstation.Hardware.CPU_list.ID, data.Workstation.Hardware.ID)
		if err != nil {
			log.Println("Error write: hardware_UPDATE_HARDWARE_SET_CPU_LIST ", err)
		}

		_, err = db.Exec(cpu_list_INSERT_CPU_LIST, data.Workstation.Hardware.CPU_list.ID)
		if err != nil {
			log.Println("Error write: cpu_list_INSERT_CPU_LIST ", err)
		}

		for _, v := range data.Workstation.Hardware.CPU_list.CPUs {

			_, err = db.Exec(cpu_INSERT_CPU,
				v.ID,
				v.Model,
				v.Manufacturer,
				v.Frequency,
				v.Number_cores,
				v.Number_threads,
			)
			if err != nil {
				log.Println("Error write: cpu_INSERT_CPU ", err)
			}

			_, err = db.Exec(cpu_list_UPDATE_CPU_list_SET_CPUID, v.ID, data.Workstation.Hardware.CPU_list.ID)
			if err != nil {
				log.Println("Error write: cpu_list_UPDATE_CPU_list_SET_CPUID ", err)
			}
		}

	}

	if data.Workstation.Hardware.RAM_list.ID != "" {
		_, err = db.Exec(hardware_UPDATE_HARDWARE_SET_RAM_LIST, data.Workstation.Hardware.RAM_list.ID, data.Workstation.Hardware.ID)
		if err != nil {
			log.Println("Error write: hardware_UPDATE_HARDWARE_SET_RAM_LIST ", err)
		}

		_, err = db.Exec(ram_list_INSERT_RAM_LIST, data.Workstation.Hardware.RAM_list.ID)
		if err != nil {
			log.Println("Error write: ram_list_INSERT_RAM_LIST ", err)
		}

		for _, v := range data.Workstation.Hardware.RAM_list.RAMs {

			_, err = db.Exec(ram_INSERT_RAM,
				v.ID,
				v.Manufacturer,
				v.Size,
				v.Frequency,
				v.Serial_number,
			)
			if err != nil {
				log.Println("Error write: ram_INSERT_RAM ", err)
			}

			_, err = db.Exec(ram_list_UPDATE_RAM_list_SET_RAMID, v.ID, data.Workstation.Hardware.RAM_list.ID)
			if err != nil {
				log.Println("Error write: ram_list_UPDATE_RAM_list_SET_RAMID ", err)
			}
		}

	}

	if data.Workstation.Hardware.HDD_list.ID != "" {
		_, err = db.Exec(hardware_UPDATE_HARDWARE_SET_HDD_LIST, data.Workstation.Hardware.HDD_list.ID, data.Workstation.Hardware.ID)
		if err != nil {
			log.Println("Error write: hardware_UPDATE_HARDWARE_SET_HDD_LIST ", err)
		}

		_, err = db.Exec(hdd_list_INSERT_HDD_LIST, data.Workstation.Hardware.HDD_list.ID)
		if err != nil {
			log.Println("Error write: hdd_list_INSERT_HDD_LIST ", err)
		}

		for _, v := range data.Workstation.Hardware.HDD_list.HDDs {
			_, err = db.Exec(hdd_INSERT_HDD,
				v.ID,
				v.Model,
				v.Size,
				v.Type,
			)
			if err != nil {
				log.Println("Error write: hdd_INSERT_HDD ", err)
			}

			_, err = db.Exec(hdd_list_UPDATE_HDD_list_SET_HDDID, v.ID, data.Workstation.Hardware.HDD_list.ID)
			if err != nil {
				log.Println("Error write: hdd_list_UPDATE_HDD_list_SET_HDDID ", err)
			}
		}

	}

}
