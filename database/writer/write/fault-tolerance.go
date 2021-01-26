package write

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

const fault_tolerance_INSERT_FAULT_T = `INSERT INTO Fault_tolerance(ID)
								VALUES($1)`

//fault-tolerance
const (
	Fault_tolerance_UPDATE_FLT_TOLERANCE_SET_FIELDS = `UPDATE Fault_tolerance
								SET Commissioning_date = $1
								, Backup = $2
								, Number_of_error = $3
								, Cluster = $4
								WHERE ID = $5`
)

func Fault_tolerance(db *sql.DB, data model.TO_WR) {
	_, err := db.Exec(fault_tolerance_INSERT_FAULT_T, data.Workstation.Fault_tolerance.ID)
	if err != nil {
		log.Println("Error write: fault_tolerance_INSERT_FAULT_T ", err)
	}

	_, err = db.Exec(Fault_tolerance_UPDATE_FLT_TOLERANCE_SET_FIELDS, data.Workstation.Fault_tolerance.Commissioning_date, data.Workstation.Fault_tolerance.Backup,
		data.Workstation.Fault_tolerance.Number_of_error, data.Workstation.Fault_tolerance.Cluster, data.Workstation.Fault_tolerance.ID)
	if err != nil {
		log.Println("Error write: Fault_tolerance_UPDATE_FLT_TOLERANCE_SET_FIELDS ", err)
	}
}
