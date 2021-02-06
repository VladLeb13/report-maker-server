package update

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

//hardware-board
const (
	matherboard_UPDATE_BOARD = `UPDATE Matherboard
								SET Name  = $1
								, Model = $2
								, Product = $3
								WHERE ID = $4`

	matherboard_SELECT_BOARD = `SELECT Name, Model, Product
								FROM Matherboard
								WHERE ID = $1`
)

//hardware-cpu
const (
	cpu_list_SELECT_CPU_LIST = `SELECT  CPU.ID AS "ID"
								, CPU.Model AS "Model"
								, CPU.Manufacturer AS "Manufacturer"
								, CPU.Frequency AS "Frequency"
								, CPU.Number_cores AS "Number_cores"
								, CPU.Number_threads AS "Number_threads"
								FROM CPU_list
								INNER JOIN CPU ON CPU.ID = CPU_list.CPUID
								WHERE CPU_list.ID = $1`

	cpu_list_DELETE_CPU_LIST = `DELETE FROM CPU_list WHERE ID = $1 AND CPUID = $2`
	cpu_DELETE_CPU           = `DELETE FROM CPU WHERE ID = $1`

	cpu_list_UPDATE_CPU_list_SET_CPUID = `UPDATE CPU_list
								SET CPUID = $1
								WHERE ID = $2`

	cpu_INSERT_CPU = `INSERT INTO CPU(ID, Model, Manufacturer, Frequency, Number_cores, Number_threads)
								VALUES($1, $2, $3, $4, $5, $6)`
)

//hardware-ram
const (
	ram_list_SELECT_RAM_LIST = `SELECT  RAM.ID AS "ID"
								, RAM.Manufacturer AS "Manufacturer"
								, RAM.Size AS "Size"
								, RAM.Frequency AS "Frequency"
								, RAM.Serial_number AS "Serial_number"
								FROM RAM_list
								INNER JOIN RAM ON RAM.ID = RAM_list.RAMID
								WHERE RAM_list.ID = $1`

	ram_list_DELETE_RAM_LIST = `DELETE FROM RAM_list WHERE ID = $1 AND RAMID = $2`
	ram_DELETE_RAM           = `DELETE FROM RAM WHERE ID = $1`

	ram_list_UPDATE_RAM_list_SET_RAMID = `UPDATE RAM_list
								SET RAMID = $1
								WHERE ID = $2`

	ram_INSERT_RAM = `INSERT INTO RAM(ID, Manufacturer,Size, Frequency, Serial_number)
								VALUES($1, $2, $3, $4, $5)`
)

//hardware-hdd
const (
	hdd_list_SELECT_HDD_LIST = `SELECT  HDD.ID AS "ID"
								, HDD.Model AS "Model"
								, HDD.Size AS "Size"
								, HDD.Type AS "Type"
								FROM HDD_list
								INNER JOIN HDD ON HDD.ID = HDD_list.HDDID
								WHERE HDD_list.ID = $1`

	hdd_list_DELETE_HDD_LIST = `DELETE FROM HDD_list WHERE ID = $1 AND HDDID = $2`
	hdd_DELETE_HDD           = `DELETE FROM HDD WHERE ID = $1`

	hdd_list_UPDATE_HDD_list_SET_HDDID = `UPDATE HDD_list
								SET HDDID = $1
								WHERE ID = $2`

	hdd_INSERT_HDD = `INSERT INTO HDD(ID, Model,Size, Type)
								VALUES($1, $2, $3, $4)`
)

func Hardware(db *sql.DB, new model.TO_WR) {

	var old model.Hardware

	err := db.QueryRow(matherboard_SELECT_BOARD, new.Workstation.Hardware.Matherboard.ID).Scan(
		&old.Matherboard.Name,
		&old.Matherboard.Model,
		&old.Matherboard.Product,
	)
	if err != nil {
		log.Println("Error write: matherboard_SELECT_BOARD ", err)
	}

	rows, err := db.Query(cpu_list_SELECT_CPU_LIST, new.Workstation.Hardware.CPU_list.ID)
	if err != nil {
		log.Println("Error write: cpu_list_SELECT_CPU_LIST ", err)
	}
	for rows.Next() {
		var cpu model.CPU
		rows.Scan(&cpu.ID,
			&cpu.Model,
			&cpu.Manufacturer,
			&cpu.Frequency,
			&cpu.Number_cores,
			&cpu.Number_threads,
		)

		old.CPU_list.CPUs = append(old.CPU_list.CPUs, cpu)
	}
	rows.Close()

	rows, err = db.Query(ram_list_SELECT_RAM_LIST, new.Workstation.Hardware.RAM_list.ID)
	if err != nil {
		log.Println("Error write: ram_list_SELECT_RAM_LIST ", err)
	}
	for rows.Next() {
		var ram model.RAM
		rows.Scan(&ram.ID,
			&ram.Manufacturer,
			&ram.Size,
			&ram.Frequency,
			&ram.Serial_number,
		)

		old.RAM_list.RAMs = append(old.RAM_list.RAMs, ram)
	}
	rows.Close()

	rows, err = db.Query(hdd_list_SELECT_HDD_LIST, new.Workstation.Hardware.HDD_list.ID)
	if err != nil {
		log.Println("Error write: hdd_list_SELECT_HDD_LIST ", err)
	}
	for rows.Next() {
		var hdd model.HDD
		rows.Scan(&hdd.ID,
			&hdd.Model,
			&hdd.Size,
			&hdd.Type,
		)

		old.HDD_list.HDDs = append(old.HDD_list.HDDs, hdd)
	}
	rows.Close()

	var update model.Hardware
	if new.Workstation.Hardware.Matherboard.Name != old.Matherboard.Name {
		update.Matherboard.Name = new.Workstation.Hardware.Matherboard.Name
	} else {
		update.Matherboard.Name = old.Matherboard.Name
	}

	if new.Workstation.Hardware.Matherboard.Model != old.Matherboard.Model {
		update.Matherboard.Model = new.Workstation.Hardware.Matherboard.Model
	} else {
		update.Matherboard.Model = old.Matherboard.Model
	}

	if new.Workstation.Hardware.Matherboard.Product != old.Matherboard.Product {
		update.Matherboard.Product = new.Workstation.Hardware.Matherboard.Product
	} else {
		update.Matherboard.Product = old.Matherboard.Product
	}

	len_new := len(new.Workstation.Hardware.CPU_list.CPUs)
	len_old := len(old.CPU_list.CPUs)
	if len_new != len_old {
		for _, new_val := range new.Workstation.Hardware.CPU_list.CPUs {
			update.CPU_list.CPUs = append(update.CPU_list.CPUs, new_val)
		}
	} else {
		for i, new_val := range new.Workstation.Hardware.CPU_list.CPUs {
			if new_val != old.CPU_list.CPUs[i] {
				update.CPU_list.CPUs = append(update.CPU_list.CPUs, new_val)
			} else {
				update.CPU_list.CPUs = append(update.CPU_list.CPUs, old.CPU_list.CPUs[i])
			}
		}
	}
	for _, v := range old.CPU_list.CPUs {
		db.Exec(cpu_list_DELETE_CPU_LIST, old.CPU_list.ID, v.ID)
		db.Exec(cpu_DELETE_CPU, v.ID)
	}

	len_new = len(new.Workstation.Hardware.RAM_list.RAMs)
	len_old = len(old.RAM_list.RAMs)
	if len_new != len_old {
		for _, new_val := range new.Workstation.Hardware.RAM_list.RAMs {
			update.RAM_list.RAMs = append(update.RAM_list.RAMs, new_val)
		}
	} else {
		for i, new_val := range new.Workstation.Hardware.RAM_list.RAMs {
			if new_val != old.RAM_list.RAMs[i] {
				update.RAM_list.RAMs = append(update.RAM_list.RAMs, new_val)
			} else {
				update.RAM_list.RAMs = append(update.RAM_list.RAMs, old.RAM_list.RAMs[i])
			}
		}
	}
	for _, v := range old.RAM_list.RAMs {
		db.Exec(ram_list_DELETE_RAM_LIST, old.RAM_list.ID, v.ID)
		db.Exec(ram_DELETE_RAM, v.ID)
	}

	len_new = len(new.Workstation.Hardware.HDD_list.HDDs)
	len_old = len(old.HDD_list.HDDs)
	if len_new != len_old {
		for _, new_val := range new.Workstation.Hardware.HDD_list.HDDs {
			update.HDD_list.HDDs = append(update.HDD_list.HDDs, new_val)
		}
	} else {
		for i, new_val := range new.Workstation.Hardware.HDD_list.HDDs {
			if new_val != old.HDD_list.HDDs[i] {
				update.HDD_list.HDDs = append(update.HDD_list.HDDs, new_val)
			} else {
				update.HDD_list.HDDs = append(update.HDD_list.HDDs, old.HDD_list.HDDs[i])
			}
		}
	}
	for _, v := range old.HDD_list.HDDs {
		db.Exec(hdd_list_DELETE_HDD_LIST, old.HDD_list.ID, v.ID)
		db.Exec(hdd_DELETE_HDD, v.ID)
	}

	_, err = db.Exec(matherboard_UPDATE_BOARD,
		update.Matherboard.Name,
		update.Matherboard.Model,
		update.Matherboard.Product,
		new.Workstation.Hardware.Matherboard.ID)
	if err != nil {
		log.Println("Error write: matherboard_UPDATE_BOARD ", err)
	}

	for _, v := range update.CPU_list.CPUs {

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

		_, err = db.Exec(cpu_list_UPDATE_CPU_list_SET_CPUID, v.ID, old.CPU_list.ID)
		if err != nil {
			log.Println("Error write: cpu_list_UPDATE_CPU_list_SET_CPUID ", err)
		}
	}

	for _, v := range update.RAM_list.RAMs {

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

		_, err = db.Exec(ram_list_UPDATE_RAM_list_SET_RAMID, v.ID, old.RAM_list.ID)
		if err != nil {
			log.Println("Error write: ram_list_UPDATE_RAM_list_SET_RAMID ", err)
		}
	}

	for _, v := range update.HDD_list.HDDs {
		_, err = db.Exec(hdd_INSERT_HDD,
			v.ID,
			v.Model,
			v.Size,
			v.Type,
		)
		if err != nil {
			log.Println("Error write: hdd_INSERT_HDD ", err)
		}

		_, err = db.Exec(hdd_list_UPDATE_HDD_list_SET_HDDID, v.ID, old.HDD_list.ID)
		if err != nil {
			log.Println("Error write: hdd_list_UPDATE_HDD_list_SET_HDDID ", err)
		}
	}

}
