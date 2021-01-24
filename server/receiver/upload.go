package receiver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"report-maker-server/database/writer"
	"report-maker-server/server/controller"
	"report-maker-server/server/normalizer"
	"report-maker-server/tools"

	"github.com/VladLeb13/report-maker-lib/datalib"
	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	app_ctx := ctx.MustGet("app-context").(*tools.AppContex)

	w := ctx.Writer
	r := ctx.Request

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

	datalib := &datalib.Report{}
	err = json.Unmarshal(content, datalib)
	if err != nil {
		log.Println("Unmarshal error: ", err)
	}

	normalizer.Actions(datalib)

	writer.AddRecord(app_ctx, datalib)

}
