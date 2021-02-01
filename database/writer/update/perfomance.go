package update

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

//perfomance
const (
	perfomance_UPDATE_PERFOMANCE = `UPDATE Perfomance
								SET CPU = $1
								, RAM = $2
								, HDD = $3
								, Cluster = $4
								WHERE ID = $5`
)

const perfomance_SELECT_PERFOMANCE = `SELECT CPU, RAM, HDD, Cluster
										FROM Perfomance
										WHERE ID = 1$`

func Perfomance(db *sql.DB, new model.TO_WR) {

	var old model.Perfomance

	err := db.QueryRow(perfomance_SELECT_PERFOMANCE, new.Workstation.Perfomance.ID).Scan(
		&old.CPU,
		&old.RAM,
		&old.HDD,
		&old.Cluster,
	)

	if err != nil {
		log.Println("Error of query: perfomance_SELECT_PERFOMANCE", err)
	}

	var update model.Perfomance
	if new.Workstation.Perfomance.CPU != old.CPU {
		update.CPU = new.Workstation.Perfomance.CPU
	} else {
		update.CPU = old.CPU
	}

	if new.Workstation.Perfomance.RAM != old.RAM {
		update.RAM = new.Workstation.Perfomance.RAM
	} else {
		update.RAM = old.RAM
	}

	if new.Workstation.Perfomance.HDD != old.HDD {
		update.HDD = new.Workstation.Perfomance.HDD
	} else {
		update.HDD = old.HDD
	}

	if new.Workstation.Perfomance.Cluster != old.Cluster {
		update.Cluster = new.Workstation.Perfomance.Cluster
	} else {
		update.Cluster = old.Cluster
	}

	_, err = db.Exec(perfomance_UPDATE_PERFOMANCE, update.CPU, update.RAM,
		update.HDD, update.Cluster, new.Workstation.Perfomance.ID)
	if err != nil {
		log.Println("Error write: perfomance_UPDATE_PERFOMANCE ", err)
	}

}
