package receiver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"report-maker-server/server/controller"
)

type TestData struct {
	Field1 string `json:"f1"`
	Field2 string `json:"f2"`
}

func Upload(w http.ResponseWriter, r *http.Request) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	login, pass, ok := r.BasicAuth()
	if !ok {
		log.Println("Auth error")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authResult, err := controller.FindInBase(login, pass)
	if err != nil {
		log.Println("Auth error: ", err)
	}

	if !authResult {
		log.Println("Auth error")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Body error: ", err)
	}

	datalib := &TestData{}
	err = json.Unmarshal(content, datalib)
	if err != nil {
		log.Println("Unmarshal error: ", err)
	}

	normalized(datalib)

	addRecord(datalib)

}
