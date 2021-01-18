package database

import (
	"database/sql"
	"log"
	"sync"

	"report-maker-server/config"
	"report-maker-server/tools"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Mu *sync.Mutex
	db *sql.DB
}

func (d *Database) synchInsert(query string, args ...interface{}) (err error) {
	d.Mu.Lock()
	defer d.Mu.Unlock()

	//d.db = Get()
	defer d.db.Close()

	_, err = d.db.Exec(query, args...)
	if err != nil {
		log.Println("Query error: ", err)
	}

	return
}

func (d *Database) synchSelect(query string, args ...interface{}) (rows *sql.Rows, err error) {
	d.Mu.Lock()
	defer d.Mu.Unlock()

	//d.db = Get()
	defer d.db.Close()

	rows, err = d.db.Query(query, args...)
	if err != nil {
		log.Println("Query error: ", err)
	}

	return rows, err
}

func Get(ctx *tools.AppContex) *sql.DB {
	cnf := ctx.Context.Value("config").(config.Config)

	path := cnf.Database_path
	db, err := sql.Open("sqlite3", path+"main.db")
	if err != nil {
		log.Println("Open database error: ", err)
	}

	return db
}
