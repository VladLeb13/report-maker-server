package tests

import (
	"log"
	"sync"
	"testing"

	"report-maker-server/database"
)

const INS = `INSERT 
INTO test_table (id,name)
VALUES ($1, $2)
`

const SEL = `SELECT name, id
FROM test_table
WHERE test_table.id = $1
`

func TestGeneral(t *testing.T) {
	mutex := &sync.Mutex{}

	db := &database.Database{
		Mu: mutex,
	}

	//db.SynchInsert(INS,1, "First")
	//db.SynchInsert(INS,2, "Second")

	rows, err := db.SynchSelect(SEL, 1)
	if err != nil {
		log.Println("Query error: ", err)
	}

	var (
		id   int
		name string
	)
	for rows.Next() {
		err := rows.Scan(&name, &id)
		if err != nil {
			log.Fatalln("Error rows: ", err)
		}
	}

	rows.Close()

}
