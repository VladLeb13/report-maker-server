package fill

import (
	"report-maker-server/server/model"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/google/uuid"
)

func Program(in datalib.Report) (program_list []model.Program, program model.Program_list) {
	var programs_list []model.Program
	for _, v := range in.Software.Programs {
		var program model.Program

		program.Manufacturer = v.Vendor
		program.Name = v.Name
		program.Install_on = v.InstallDate
		program.Version = v.Version

		programs_list = append(programs_list, program)
	}

	program.ID = uuid.New().String()
	for _, v := range programs_list {
		uuid := uuid.New().String()

		program.ProgramID = append(program.ProgramID, uuid)
		v.ID = uuid
	}

	return
}
