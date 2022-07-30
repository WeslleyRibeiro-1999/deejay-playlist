package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/deejay?sslmode=disable")
	if err != nil {
		log.Fatal("could not connect with db: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
