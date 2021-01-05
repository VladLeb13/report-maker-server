package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Mu *sync.Mutex
	db *sql.DB
}

func (d *Database) SynchInsert(query string, args ...interface{}) (err error) {
	//todo: create before context
	d.Mu.Lock()
	defer d.Mu.Unlock()

	d.db = getConn()
	defer d.db.Close()

	_, err = d.db.Exec(query, args...)
	if err != nil {
		log.Println("Query error: ", err)
	}

	return
}

func (d *Database) SynchSelect(query string, args ...interface{}) (rows *sql.Rows, err error) {
	//todo: create before context
	d.Mu.Lock()
	defer d.Mu.Unlock()

	d.db = getConn()
	defer d.db.Close()

	rows, err = d.db.Query(query, args...)
	if err != nil {
		log.Println("Query error: ", err)
	}

	return rows, err
}

func getConn() *sql.DB {

	//todo: context: config pathDb
	path := "/home/worker/Studing/report-maker-server/src/database/"

	db, err := sql.Open("sqlite3", path+"main.db")
	if err != nil {
		log.Println("Open database error: ", err)
	}

	return db
}
