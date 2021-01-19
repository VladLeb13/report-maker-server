package fill

import (
	"report-maker-server/server/model"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/google/uuid"
)

func Fault_tolerance(report datalib.Report) (flt model.Fault_tolerance) {
	flt.Cluster = 0
	flt.Number_of_error = len(report.Events.List)
	flt.Backup = 1
	flt.Commissioning_date = "2020-06-22T00:00:00Z"
	flt.ID = uuid.New().String()

	return
}
