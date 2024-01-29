package service

import (
	repository "BotanyBackend/repository"
	"BotanyBackend/servers/g-rpc/Type"
)

type EcoTypeService interface {
	Insert(input Type.CreateEcoEntityRequest) (*Type.EcoEntity, error)
	GetById(inputId string) (Type.EcoEntity, error)
	GetList() (Type.EcoEntityList, error)
}

type Service struct {
	EcoTypeService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		EcoTypeService: NewEcoTypeServiceImpl(repos),
	}
}
