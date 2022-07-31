package main

import (
	"database/sql"
	"log"

	"github.com/WeslleyRibeiro-1999/deejay-playlist/api"
	db "github.com/WeslleyRibeiro-1999/deejay-playlist/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/deejay?sslmode=disable")
	if err != nil {
		log.Fatal("Não foi possivel conectar: ", err)
	}

	store := db.NewSQLStore(conn)
	server := api.InstanceServer(store)

	err = server.Start("0.0.0.0:8000")
	if err != nil {
		log.Fatal("Api iniciada com erro: ", err)
	}
}
