package main

import (
	repository "BotanyBackend/repository"
	"BotanyBackend/servers/g-rpc/Type"
	"BotanyBackend/servers/g-rpc/TypeImpl"
	"BotanyBackend/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	//Создание Сервера
	srv := TypeImpl.EcoCRUDServiceServerImpl{Sevice: service}

	//Регистрация Сервера
	Type.RegisterEcoCRUDServiceServer(s, srv)

	reflection.Register(s)

	log.Println("Starting server on :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
