package g_rpc

import (
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/implementation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpc(authServet *api.AuthServerImpl, newRepository *service.Service) {
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		s := grpc.NewServer(
			grpc.UnaryInterceptor(middlewares.AuthInterceptor))

		//Создание Сервера
		ecomorphsEntityServet := implementation.NewEcomorphsEntityServetImpl(newRepository)
		ecomorphsServet := implementation.NewEcomorphsServetImplImpl(newRepository)
		typePlantServet := implementation.NewTypePlantServetImpl(newRepository)
		trialSiteServet := implementation.NewTrialSiteServetImpl(newRepository)
		transectServet := implementation.NewTransectServetImpl(newRepository)
		imgServet := implementation.NewImgServetImpl(newRepository)
		analysisServet := implementation.NewAnalysisServetImplServetImpl(newRepository)

		//Регистрация Сервера
		api.RegisterEcomorphEntityServiceServer(s, ecomorphsEntityServet)
		api.RegisterEcomorphServiceServer(s, ecomorphsServet)
		api.RegisterAuthServiceServer(s, authServet)
		api.RegisterTypePlantServiceServer(s, typePlantServet)
		api.RegisterTrialSiteServiceServer(s, trialSiteServet)
		api.RegisterTransectServiceServer(s, transectServet)
		api.RegisterImgServiceServer(s, imgServet)
		api.RegisterAnalysisServiceServer(s, analysisServet)

		reflection.Register(s)

		log.Println("Starting server on :50051")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}

	}()

}
