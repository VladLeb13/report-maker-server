package write

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

const workstation_INSERT_WORKSTATION = `INSERT INTO Workstation(ID, Name, Comment, Allow_analysis,  
									HardwareID, Program_listID, PerfomanceID, Fault_toleranceID)
								VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

func Workstation(db *sql.DB, data model.TO_WR) {

	_, err := db.Exec(workstation_INSERT_WORKSTATION,
		data.Workstation.ID,
		data.Workstation.Name,
		data.Workstation.Comment,
		data.Workstation.Allow_analysis,
		data.Workstation.Hardware.ID,
		data.Workstation.Program_list.ID,
		data.Workstation.Perfomance.ID,
		data.Workstation.Fault_tolerance.ID,
	)
	if err != nil {
		log.Println("Error write: workstation_INSERT_WORKSTATION ", err)
	}
}
