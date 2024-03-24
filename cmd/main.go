package main

import (
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/graphql"
	"log"
)

func main() {

	cfgDB := repository.Config{
		Host:     "localhost",
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

	db.AutoMigrate(api.EcomorphORM{}, api.EcomorphsEntityORM{}, api.UserORM{}, api.TypePlantORM{}, api.TransectORM{}, api.TrialSiteORM{})
	log.Println("migrant")

	authServet := api.NewAuthServer(db)

	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)

	g_rpc.StartGrpc(&authServet, newService)
	graphql.StartGraphQl(db, &authServet, newService)

	select {}
}
