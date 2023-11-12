package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var testQueries *Queries
var testStore *Store

func TestMain(m *testing.M) {
	dbSource := "postgresql://postgres:secretforgithubaction@localhost:5432/masterclass?sslmode=disable"

	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}

	testQueries = New(testDB)
	testStore = NewStore(testDB)

	os.Exit(m.Run())
}

func AccessSecretFile(filename string) string {
	path := fmt.Sprintf("../../secrets/%s", filename)
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf(">> secret: %s", err.Error())
	}

	return string(file)
}
