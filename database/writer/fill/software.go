package fill

import (
	"report-maker-server/server/model"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/google/uuid"
)

func Program(report datalib.Report) (list model.Program_list) {
	list.ID = uuid.New().String()

	for _, v := range report.Software.Programs {
		var program model.Program

		program.ID = uuid.New().String()
		program.Manufacturer = v.Vendor
		program.Name = v.Name
		program.Install_on = v.InstallDate
		program.Version = v.Version

		list.Programs = append(list.Programs, program)
	}

	return
}
