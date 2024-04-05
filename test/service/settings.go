package service

import (
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

func CreateService() *service.Service {

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
		return nil
	}

	err = db.AutoMigrate(api.ImgORM{}, api.EcomorphORM{}, api.EcomorphsEntityORM{}, api.UserORM{}, api.TypePlantORM{}, api.PlantORM{}, api.TransectORM{}, api.TrialSiteORM{})

	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)

	return newService
}
