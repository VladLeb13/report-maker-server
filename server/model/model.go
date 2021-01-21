package model

type (
	Workstation struct {
		ID              string          `json:"ID"`
		Name            string          `json:"Name"`
		Comment         string          `json:"Comment"`
		Allow_analysis  int             `json:"Allow_analysis"`
		Hardware        Hardware        `json:"Hardware"`
		Program_list    Program_list    `json:"Program_list"`
		Perfomance      Perfomance      `json:"Perfomance"`
		Fault_tolerance Fault_tolerance `json:"Fault_tolerance"`
	}
	Program_list struct {
		ID       string    `json:"ID"`
		Programs []Program `json:"Programs"`
	}
	Program struct {
		ID           string `json:"ID"`
		Manufacturer string `json:"Manufacturer"`
		Name         string `json:"Name"`
		Version      string `json:"Version"`
		Install_on   string `json:"Install_on"`
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

//Hardware
type (
	Hardware struct {
		ID          string      `json:"ID"`
		Matherboard Matherboard `json:"Matherboard"`
		CPU_list    CPU_list    `json:"CPU_list"`
		RAM_list    RAM_list    `json:"RAM_list"`
		HDD_list    HDD_list    `json:"HDD_list"`
	}

	Matherboard struct {
		ID      string `json:"ID"`
		Name    string `json:"Name"`
		Model   string `json:"Model"`
		Product string `json:"Product"`
	}

	CPU_list struct {
		ID   string `json:"ID"`
		CPUs []CPU  `json:"CPUs"`
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
		ID   string `json:"ID"`
		RAMs []RAM  `json:"RAMs"`
	}
	RAM struct {
		ID            string `json:"ID"`
		Manufacturer  string `json:"Manufacturer"`
		Size          int    `json:"Size"`
		Frequency     int    `json:"Frequency"`
		Serial_number string `json:"Serial_number"`
	}

	HDD_list struct {
		ID   string `json:"ID"`
		HDDs []HDD  `json:"HDDs"`
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
		ID     string            `json:"ID"`
		Date   string            `json:"Date"`
		Result Monitoring_result `json:"Result"`
	}

	Monitoring_result struct {
		ID                       string                   `json:"ID"`
		Upgrade_workstation_list Upgrade_workstation_list `json:"Upgrade_workstation_list"`
	}

	Upgrade_workstation_list struct {
		ID                 string              `json:"ID"`
		Date               string              `json:"Date"`
		Upgrade_list_items []Upgrade_list_item `json:"Upgrade_list_items"`
	}

	Upgrade_list_item struct {
		ID            string `json:"ID"`
		WorkstationID string `json:"WorkstationID"`
		Description   string `json:"Description"`
	}
)

type TO_WR struct {
	Workstation Workstation
	Monitoring  Monitoring
}
