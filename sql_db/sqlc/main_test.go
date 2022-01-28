package sql_db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:example@localhost:5432/sm_db?sslmode=disable"
)

func TestMain(m *testing.M) {
	log.Println("Testing")
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
