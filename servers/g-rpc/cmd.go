package g_rpc

import (
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/auth"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/transect"
	trial_site "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/trial-site"
	type_plant "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/type-plant"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph-entity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpc(db *gorm.DB) {
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		s := grpc.NewServer(
			grpc.UnaryInterceptor(middlewares.AuthInterceptor))

		//Создание Сервера
		ecomorphsEntityServet := ecomorph_entity.NewEcomorphsEntityServetImpl(db)
		ecomorphsServet := ecomorph.NewEcomorphsServetImplImpl(db)
		authServet := auth.NewAuthServer(db)
		typePlantServet := type_plant.NewTypePlantServetImpl(db)
		trialSiteServet := trial_site.NewTrialSiteServetImpl(db)
		transectServet := transect.NewTransectServetImpl(db)

		//Регистрация Сервера
		ecomorph_entity.RegisterEcomorphEntityServiceServer(s, ecomorphsEntityServet)
		ecomorph.RegisterEcomorphServiceServer(s, ecomorphsServet)
		auth.RegisterAuthServiceServer(s, authServet)
		type_plant.RegisterTypePlantServiceServer(s, typePlantServet)
		trial_site.RegisterTrialSiteServiceServer(s, trialSiteServet)
		transect.RegisterTransectServiceServer(s, transectServet)

		reflection.Register(s)

		log.Println("Starting server on :50051")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}
