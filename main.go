package main

import (
	"context"
	"github.com/CharlieAlphaQA/simplebank/api"
	db "github.com/CharlieAlphaQA/simplebank/db/sqlc"
	"github.com/CharlieAlphaQA/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err2 := util.LoadConfig(".")
	if err2 != nil {
		log.Fatal("cannot connect to config:", err2)
	}
	//conn, err := sql.Open(config.DBDriver, config.DBSource)
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(connPool)
	server, err1 := api.NewServer(config, store)
	if err1 != nil {
		log.Fatal("cannot create server:", err1)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
