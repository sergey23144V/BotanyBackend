package g_rpc

import (
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/implementation"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpc(ser service.Service) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	//Создание Сервера
	ecomorphsEntityServet := implementation.NewEcomorphsEntityServetImpl(ser)
	ecomorphsServet := implementation.NewEcomorphsServetImplImpl(ser)

	//Регистрация Сервера
	RegisterEcomorphEntityServiceServer(s, ecomorphsEntityServet)
	RegisterEcomorphServiceServer(s, ecomorphsServet)

	reflection.Register(s)

	log.Println("Starting server on :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
