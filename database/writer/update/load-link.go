package update

import (
	"database/sql"

	"report-maker-server/server/model"
)

const select_links = `SELECT Workstation.ID
					  , Program_listID
					  , PerfomanceID
					  , Fault_toleranceID
					  , HardwareID
					  , Matherboard.ID
					  , CPU_list.ID
					  , RAM_list.ID
					  , HDD_list.ID
					  FROM Workstation
					  	INNER JOIN Hardware  ON Hardware.ID = Workstation.HardwareID
						INNER JOIN Matherboard ON Matherboard.ID = Hardware.MatherboardID
						INNER JOIN CPU_list ON CPU_list.ID = Hardware.CPU_listID
						INNER JOIN RAM_list ON RAM_list.ID = Hardware.RAM_listID
						INNER JOIN HDD_list ON HDD_list.ID = Hardware.HDD_listID
					  WHERE Workstation.Name = $1`

func LoadLink(db *sql.DB, data *model.TO_WR) (err error) {

	err = db.QueryRow(select_links, data.Workstation.Name).Scan(
		&data.Workstation.ID,
		&data.Workstation.Program_list.ID,
		&data.Workstation.Perfomance.ID,
		&data.Workstation.Fault_tolerance.ID,
		&data.Workstation.Hardware.ID,
		&data.Workstation.Hardware.Matherboard.ID,
		&data.Workstation.Hardware.CPU_list.ID,
		&data.Workstation.Hardware.RAM_list.ID,
		&data.Workstation.Hardware.HDD_list.ID,
	)
	return
}
