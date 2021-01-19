package fill

import (
	"strconv"

	"report-maker-server/server/model"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/google/uuid"
)

func Perfomance(in datalib.Perfomance) (perfomance model.Perfomance) {
	if len(in.HDD) != 0 {
		perfomance.HDD = strconv.Itoa(int(in.HDD[0].PercentDiskTime))
	}

	if len(in.RAM) != 0 {
		perfomance.RAM = strconv.Itoa(int(in.RAM[0].AvailableMBytes))
	}

	if len(in.CPU) != 0 {
		perfomance.CPU = strconv.Itoa(int(in.CPU[0].PercentProcessorUtility))
	}

	perfomance.Cluster = 0
	perfomance.ID = uuid.New().String()

	return
}
