package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Parse() (err error) {
	configfile, err := os.Open("config/config.json")
	if err != nil {
		return err
	}
	defer configfile.Close()

	data, err := ioutil.ReadAll(configfile)
	if err != nil {
		return err
	}

	cnf := &Config{}
	err = json.Unmarshal(data, &cnf)
	if err != nil {
		return err
	}

	//TODO: add to context
	//ctx.Context = context.WithValue(ctx.Context, "configuration", cnf)

	return nil
}
