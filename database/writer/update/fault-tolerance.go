package update

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

//fault-tolerance
const (
	fault_tolerance_UPDATE_FLT_TOLERANCE = `UPDATE Fault_tolerance
								SET Commissioning_date = $1
								, Backup = $2
								, Number_of_error = $3
								, Cluster = $4
								WHERE ID = $5`
)

const fault_tolerance_SELECT__FLT_TOLERANCE = `SELECT Commissioning_date, Backup, Number_of_error, Cluster
										FROM Fault_tolerance
										WHERE ID = 1$`

func Fault_tolerance(db *sql.DB, new model.TO_WR) {

	var old model.Fault_tolerance

	err := db.QueryRow(fault_tolerance_SELECT__FLT_TOLERANCE, new.Workstation.Fault_tolerance.ID).Scan(
		&old.Commissioning_date,
		&old.Backup,
		&old.Number_of_error,
		&old.Cluster,
	)

	if err != nil {
		log.Println("Error of query: fault_tolerance_SELECT__FLT_TOLERANCE", err)
	}

	var update model.Fault_tolerance
	if new.Workstation.Fault_tolerance.Commissioning_date != old.Commissioning_date {
		update.Commissioning_date = new.Workstation.Fault_tolerance.Commissioning_date
	} else {
		update.Commissioning_date = old.Commissioning_date
	}

	if new.Workstation.Fault_tolerance.Backup != old.Backup {
		update.Backup = new.Workstation.Fault_tolerance.Backup
	} else {
		update.Backup = old.Backup
	}

	if new.Workstation.Fault_tolerance.Number_of_error != old.Number_of_error {
		update.Number_of_error = new.Workstation.Fault_tolerance.Number_of_error
	} else {
		update.Number_of_error = old.Number_of_error
	}

	if new.Workstation.Fault_tolerance.Cluster != old.Cluster {
		update.Cluster = new.Workstation.Fault_tolerance.Cluster
	} else {
		update.Cluster = old.Cluster
	}

	_, err = db.Exec(fault_tolerance_UPDATE_FLT_TOLERANCE, update.Commissioning_date, update.Backup,
		update.Number_of_error, update.Cluster, new.Workstation.Fault_tolerance.ID)
	if err != nil {
		log.Println("Error write: fault_tolerance_UPDATE_FLT_TOLERANCE ", err)
	}

}
