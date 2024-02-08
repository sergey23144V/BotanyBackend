package main

import (
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"github.com/sergey23144V/BotanyBackend/servers/grapgql"
	"log"
)

func main() {

	cfgDB := repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		DBName:   "botanynew",
		SSLMode:  "disable",
		Password: "123321",
	}

	db, err := repository.ConnectDB(cfgDB)
	if err != nil {
		log.Fatal(err)
	}

	g_rpc.StartGrpc(db)
	grapgql.StartGraphQl(db)

	select {}
}
