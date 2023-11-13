package db

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"masterclass/util"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	args        util.FlagArgs
	testQueries *Queries
	testStore   *Store
)

func TestMain(m *testing.M) {
	args = util.DeclareFlag()
	flag.Parse()

	config, err := util.LoadConfig("../..", *args.Env)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbSource := fmt.Sprintf("postgresql://postgres:%s@localhost:5432/masterclass?sslmode=disable", config.DBPassword)

	testDB, err := sql.Open(config.DBDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}

	testQueries = New(testDB)
	testStore = NewStore(testDB)

	os.Exit(m.Run())
}
