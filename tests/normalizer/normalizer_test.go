package normalizer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"report-maker-server/server/normalizer"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

func TestIdempotency(t *testing.T) {
	for i := 0; i < 40; i++ {
		f, err := os.Open("example.json")
		if err != nil {
			log.Fatalln(err)
		}

		defer f.Close()
		data, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatalln(err)
		}

		datalib := &datalib.Report{}
		err = json.Unmarshal(data, &datalib)
		if err != nil {
			log.Fatalln(err)
		}

		normalizer.Actions(datalib)

		f2, _ := os.Create("example2.json")
		//if err != nil {
		//	log.Fatalln(err)
		//}
		defer f2.Close()

		b, err := json.Marshal(datalib)
		if err != nil {
			log.Fatalln(err)
		}

		f2.WriteString(string(b))
	}
}

func TestGeneral(t *testing.T) {
	f, err := os.Open("example.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	datalib := &datalib.Report{}
	err = json.Unmarshal(data, &datalib)
	if err != nil {
		log.Fatalln(err)
	}

	normalizer.Actions(datalib)

	f2, _ := os.Create("example2.json")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	defer f2.Close()

	b, err := json.Marshal(datalib)
	if err != nil {
		log.Fatalln(err)
	}

	f2.WriteString(string(b))

}
