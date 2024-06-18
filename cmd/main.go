package main

import (
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/implementation"
	"github.com/sergey23144V/BotanyBackend/servers/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/rest"
	"log"
)

func main() {

	cfgDB := repository.Config{
		Host:     "db",
		Port:     "5432",
		Username: "postgres",
		DBName:   "botany",
		SSLMode:  "disable",
		Password: "123321",
	}

	db, err := repository.ConnectDB(cfgDB)
	if err != nil {
		log.Fatal(err)
	}

	repository.Migrate(db)

	authServet := implementation.NewAuthServer(db)

	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)

	g_rpc.StartGrpc(&authServet, newService)
	rest.StartRest(newService)
	graphql.StartGraphQl(db, &authServet, newService)

	select {}
}
