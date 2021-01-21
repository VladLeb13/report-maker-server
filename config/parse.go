package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Database_path   string `json:"database_path"`
	Template_path   string `json:"template_path"`
	Time_format     string `json:"time_format"`
	Scheduler_cycle int    `json:"scheduler_cycle"`
}

func Parse() (config Config, err error) {
	//TODO: Раскоментить при завершении тестирования
	//configfile, err := os.Open("config/config.json")

	configfile, err := os.Open("/home/lebedev/Документы/srv/src/report-maker-server/config/config.json")
	if err != nil {
		return Config{}, err
	}
	defer configfile.Close()

	data, err := ioutil.ReadAll(configfile)
	if err != nil {
		return Config{}, err
	}

	cnf := Config{}
	err = json.Unmarshal(data, &cnf)
	if err != nil {
		return Config{}, err
	}

	return cnf, nil
}
