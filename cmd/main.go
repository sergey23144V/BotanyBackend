package main

import (
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
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

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	g_rpc.StartGrpc(*service)

}
