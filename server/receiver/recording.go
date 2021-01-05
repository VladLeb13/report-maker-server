package receiver

import "log"

func addRecord(data *TestData) {
	//TODO: write to db report

	log.Printf("%+v", data)
}
