package server

import (
	"errors"
	"log"
	"net/http"

	"report-maker-server/server/controller"

	"github.com/gorilla/mux"
)

func Serve() (err error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	router := mux.NewRouter()

	log.Println("Server start in localhost:8080 ....")

	router.HandleFunc("/home", controller.Home)
	router.HandleFunc("/reports", controller.Reports)
	router.HandleFunc("/login", controller.Login)

	router.HandleFunc("/logining", controller.Logining).Methods("POST")

	router.Use(controller.BaseAuth)

	http.ListenAndServe("localhost:8080", router)

	return errors.New("Server shutdown")

}
