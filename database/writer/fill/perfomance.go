package fill

import (
	"strconv"

	"report-maker-server/server/model"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/google/uuid"
)

func Perfomance(report datalib.Report) (perf model.Perfomance) {
	if len(report.Perfomance.HDD) != 0 {
		perf.HDD = strconv.Itoa(int(report.Perfomance.HDD[0].PercentDiskTime))
	}

	if len(report.Perfomance.RAM) != 0 {
		perf.RAM = strconv.Itoa(int(report.Perfomance.RAM[0].AvailableMBytes))
	}

	if len(report.Perfomance.CPU) != 0 {
		perf.CPU = strconv.Itoa(int(report.Perfomance.CPU[0].PercentProcessorUtility))
	}

	perf.Cluster = 0
	perf.ID = uuid.New().String()

	return
}
