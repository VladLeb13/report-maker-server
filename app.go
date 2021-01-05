package main

import (
	"log"

	"report-maker-server/server"
)

func main() {
	//TODO: make serve
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := server.Serve()
	if err != nil {
		log.Println(err)
	}

}
