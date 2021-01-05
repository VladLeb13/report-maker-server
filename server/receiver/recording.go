package receiver

import (
	"log"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

func addRecord(data *datalib.Report) {
	//TODO: write to db report

	log.Printf("%+v", data)
}
