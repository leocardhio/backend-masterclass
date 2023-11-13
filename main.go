package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"masterclass/api"
	db "masterclass/db/sqlc"
	"masterclass/util"

	_ "github.com/lib/pq"
)

var (
	args util.FlagArgs
)

func main() {
	args = util.DeclareFlag()
	flag.Parse()

	fmt.Printf("Running on %s environment\n", *args.Env)
	config, err := util.LoadConfig(".", *args.Env)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbSource := "postgresql://postgres:verysecretpassword@localhost:5432/masterclass?sslmode=disable"

	conn, err := sql.Open(config.DBDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err.Error())
	}
}
