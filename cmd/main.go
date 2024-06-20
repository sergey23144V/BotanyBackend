package main

import (
	"github.com/joho/godotenv"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/implementation"
	"github.com/sergey23144V/BotanyBackend/servers/graphql"
	"github.com/sergey23144V/BotanyBackend/servers/rest"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	cfgDB := repository.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("USER"),
		DBName:   os.Getenv("DB"),
		SSLMode:  os.Getenv("SSLMode"),
		Password: os.Getenv("PASSWORD"),
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
