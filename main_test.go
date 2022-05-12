package main

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"tutorial.sqlc.dev/app/tutorial"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const DBUrl = "postgres://root:secret@localhost:5432/hazel"

var testQueries *tutorial.Queries

func TestMain(m *testing.M) {

	db, err := sql.Open("pgx", DBUrl)
	if err != nil {
		log.Fatalf("couldn't connect to database, %v", err)
	}

	testQueries = tutorial.New(db)

	os.Exit(m.Run())
}
