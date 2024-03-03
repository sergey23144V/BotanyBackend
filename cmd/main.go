package main

import (
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/auth"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph-entity"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/transect"
	trial_site "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/trial-site"
	type_plant "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/type-plant"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/user"
	"github.com/sergey23144V/BotanyBackend/servers/grapgql"
	"log"
	"os"
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

	db.LogMode(true)
	db.AutoMigrate(ecomorph.EcomorphORM{}, ecomorph_entity.EcomorphsEntityORM{}, user.UserORM{}, type_plant.TypePlantORM{}, trial_site.TrialSiteORM{}, transect.TransectORM{})
	log.Println("migrant")

	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	authServet := auth.NewAuthServer(db)

	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)

	g_rpc.StartGrpc(db, &authServet, newService)
	grapgql.StartGraphQl(db, &authServet, newService)

	select {}
}
