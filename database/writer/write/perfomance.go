package write

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

const perfomance_INSERT_PERFOMANCE = `INSERT INTO Perfomance(ID)
								VALUES($1)`

//perfomance
const (
	perfomance_UPDATE_PERFOMANCE_SET_FIELDS = `UPDATE Perfomance
								SET CPU = $1
								, RAM = $2
								, HDD = $3
								, Cluster = $4
								WHERE ID = $5`
)

func Perfomance(db *sql.DB, data model.TO_WR) {
	_, err := db.Exec(perfomance_INSERT_PERFOMANCE, data.Workstation.Perfomance.ID)
	if err != nil {
		log.Println("Error write: perfomance_INSERT_PERFOMANCE ", err)
	}

	_, err = db.Exec(perfomance_UPDATE_PERFOMANCE_SET_FIELDS, data.Workstation.Perfomance.CPU, data.Workstation.Perfomance.RAM,
		data.Workstation.Perfomance.HDD, data.Workstation.Perfomance.Cluster, data.Workstation.Perfomance.ID)
	if err != nil {
		log.Println("Error write: perfomance_UPDATE_PERFOMANCE_SET_FIELDS ", err)
	}

}
