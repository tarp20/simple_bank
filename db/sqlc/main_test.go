package db

import (
	    "testing"
		"database/sql"
		"log"

		_ "github.com/lib/pq"
	)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot open database", err)
	}

	testQueries = New(testDB)

	m.Run()
}