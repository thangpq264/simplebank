package db

import (
	"database/sql"
	"testing"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mylove764@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries = Queries 

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect db: ", err)
	}

	testQueries = New(conn)
}
