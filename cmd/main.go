package main

import (
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/auth"
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

	authServet := auth.NewAuthServer(db)

	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)

	g_rpc.StartGrpc(db, &authServet, newService)
	grapgql.StartGraphQl(db, &authServet, newService)

	select {}
}
