package update

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

//program list
const (
	program_list_SELECT_PROGRAM_LIST = `SELECT  Program.ID AS "ID"
								, Program.Manufacturer AS "Manufacturer"
								, Program.Name AS "Name"
								, Program.Version AS "Version"
								, Program.Install_on AS "Install_on"
								FROM Program_list
								INNER JOIN Program ON Program.ID = Program_list.ProgramID
								WHERE Program_list.ID = $1`

	program_list_DELETE_PROGRAM_LIST = `DELETE FROM Program_list WHERE ID = $1 AND ProgramID = $2`
	program_DELETE_PROGRAM           = `DELETE FROM Program WHERE ID = $1`

	program_list_UPDATE_Program_list_SET_ProgramID = `UPDATE Program_list
								SET ProgramID = $1
								WHERE ID = $2`

	program_INSERT_Program = `INSERT INTO Program(ID, Manufacturer, Name, Version, Install_on)
								VALUES($1, $2, $3, $4, $5)`
)

func Software(db *sql.DB, new model.TO_WR) {
	var old model.Program_list

	rows, err := db.Query(program_list_SELECT_PROGRAM_LIST, new.Workstation.Program_list.ID)
	if err != nil {
		log.Println("Error write: program_list_SELECT_PROGRAM_LIST ", err)
	}
	for rows.Next() {
		var program model.Program
		rows.Scan(&program.ID,
			&program.Manufacturer,
			&program.Name,
			&program.Version,
			&program.Install_on,
		)

		old.Programs = append(old.Programs, program)
	}
	rows.Close()

	var update model.Program_list
	len_new := len(new.Workstation.Program_list.Programs)
	len_old := len(old.Programs)
	if len_new != len_old {
		for _, new_val := range new.Workstation.Program_list.Programs {
			update.Programs = append(update.Programs, new_val)
		}
	} else {
		for i, new_val := range new.Workstation.Program_list.Programs {
			if new_val != old.Programs[i] {
				update.Programs = append(update.Programs, new_val)
			} else {
				update.Programs = append(update.Programs, old.Programs[i])
			}
		}
	}
	for _, v := range old.Programs {
		db.Exec(program_list_DELETE_PROGRAM_LIST, old.ID, v.ID)
		db.Exec(program_DELETE_PROGRAM, v.ID)
	}

	for _, v := range update.Programs {

		_, err = db.Exec(program_INSERT_Program,
			v.ID,
			v.Manufacturer,
			v.Name,
			v.Version,
			v.Install_on,
		)
		if err != nil {
			log.Println("Error write: program_INSERT_Program ", err)
		}

		_, err = db.Exec(program_list_UPDATE_Program_list_SET_ProgramID, v.ID, new.Workstation.Program_list.ID)
		if err != nil {
			log.Println("Error write: program_list_UPDATE_Program_list_SET_ProgramID ", err)
		}
	}
}
