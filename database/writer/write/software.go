package write

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

const program_list_INSERT_PROG_LIST = `INSERT INTO Program_list(ID)
								VALUES($1)`

//program list
const (
	program_list_UPDATE_Program_list_SET_ProgramID = `UPDATE Program_list
								SET ProgramID = $1
								WHERE ID = $2`

	program_INSERT_Program = `INSERT INTO Program(ID, Manufacturer, Name, Version, Install_on)
								VALUES($1, $2, $3, $4, $5)`
)

func Software(db *sql.DB, data model.TO_WR) {
	_, err := db.Exec(program_list_INSERT_PROG_LIST, data.Workstation.Program_list.ID)
	if err != nil {
		log.Println("Error write: program_list_INSERT_PROG_LIST ", err)
	}
	if len(data.Workstation.Program_list.Programs) > 0 {
		for _, v := range data.Workstation.Program_list.Programs {
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

			_, err = db.Exec(program_list_UPDATE_Program_list_SET_ProgramID, v.ID, data.Workstation.Program_list.ID)
			if err != nil {
				log.Println("Error write: program_list_UPDATE_Program_list_SET_ProgramID ", err)
			}
		}
	}

}
