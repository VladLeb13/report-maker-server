package update

import (
	"database/sql"
	"log"

	"report-maker-server/server/model"
)

const workstation_UPDATE_WORKSTATION = `UPDATE Workstation
										SET Comment  = $1
										, Allow_analysis = $2
										WHERE ID = $3`

const workstation_SELECT_WORKSTATION = `SELECT Comment, Allow_analysis,  
										FROM Workstation
										WHERE ID = 1$`

func Workstation(db *sql.DB, new model.TO_WR) {

	var old model.Workstation

	err := db.QueryRow(workstation_SELECT_WORKSTATION, new.Workstation.ID).Scan(
		&old.Comment,
		&old.Allow_analysis,
	)

	if err != nil {
		log.Println("Error of query: workstation_SELECT_WORKSTATION", err)
	}

	var update model.Workstation
	if new.Workstation.Comment != old.Comment {
		update.Comment = new.Workstation.Comment
	} else {
		update.Comment = old.Comment
	}
	if new.Workstation.Allow_analysis != old.Allow_analysis {
		update.Allow_analysis = new.Workstation.Allow_analysis
	} else {
		update.Comment = old.Comment
	}

	_, err = db.Exec(workstation_UPDATE_WORKSTATION,
		update.Comment,
		update.Allow_analysis,
		new.Workstation.ID,
	)
	if err != nil {
		log.Println("Error write: workstation_UPDATE_WORKSTATION ", err)
	}
}
