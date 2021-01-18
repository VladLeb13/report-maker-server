package model

type (
	Workstation struct {
		ID                string `json:"ID"`
		Name              string `json:"Name"`
		Comment           string `json:"Comment"`
		Allow_analysis    int    `json:"Allow_analysis"`
		HardwareID        string `json:"HardwareID"`
		Program_listID    string `json:"Program_listID"`
		PerfomanceID      string `json:"PerfomanceID"`
		Fault_toleranceID string `json:"Fault_toleranceID"`
	}
	Hardware struct {
		ID            string `json:"ID"`
		CPU_listID    string `json:"CPU_listID"`
		MatherboardID string `json:"MatherboardID"`
		RAM_listID    string `json:"RAM_listID"`
		HDD_listID    string `json:"HDD_listID"`
	}
	Program_list struct {
		ID        string `json:"ID"`
		ProgramID string `json:"ProgramID"`
	}

	Perfomance struct {
		ID      string `json:"ID"`
		CPU     string `json:"CPU"`
		RAM     string `json:"RAM"`
		HDD     string `json:"HDD"`
		Cluster int    `json:"Cluster"`
	}

	Fault_tolerance struct {
		ID                 string `json:"ID"`
		Commissioning_date string `json:"Commissioning_date"`
		Backup             int    `json:"Backup"`
		Number_of_error    int    `json:"Number_of_error"`
		Cluster            int    `json:"Cluster"`
	}
)

type Program struct {
	ID           string `json:"ID"`
	Manufacturer string `json:"Manufacturer"`
	Name         string `json:"Name"`
	Version      string `json:"Version"`
	Install_on   string `json:"Install_on"`
}

//Hardware
type (
	CPU_list struct {
		ID    string `json:"ID"`
		CPUID string `json:"CPUID"`
	}
	CPU struct {
		ID             string `json:"ID"`
		Model          string `json:"Model"`
		Manufacturer   string `json:"Manufacturer"`
		Frequency      int    `json:"Frequency"`
		Number_cores   int    `json:"Number_cores"`
		Number_threads int    `json:"Number_threads"`
	}
	RAM_list struct {
		ID    string `json:"ID"`
		RAMID string `json:"RAMID"`
	}
	RAM struct {
		ID            string `json:"ID"`
		Manufacturer  string `json:"Manufacturer"`
		Size          int    `json:"Size"`
		Frequency     int    `json:"Frequency"`
		Serial_number string `json:"Serial_number"`
	}
	HDD_list struct {
		ID    string `json:"ID"`
		HDDID string `json:"HDDID"`
	}
	HDD struct {
		ID    string `json:"ID"`
		Model string `json:"Model"`
		Size  int    `json:"Size"`
		Type  int    `json:"Type"`
	}
)

type (
	Monitoring struct {
		ID       string `json:"ID"`
		Date     string `json:"Date"`
		ResultID string `json:"ResultID"`
	}
	Monitoring_result struct {
		ID                         string `json:"ID"`
		Upgrade_workstation_listID string `json:"Upgrade_workstation_listID"`
	}
	Upgrade_list_item struct {
		ID            string `json:"ID"`
		WorkstationID string `json:"WorkstationID"`
		Description   string `json:"Description"`
	}
	Upgrade_workstation_list struct {
		ID                  string `json:"ID"`
		Date                string `json:"Date"`
		Upgrade_list_itemID string `json:"Upgrade_list_itemID"`
	}
)
