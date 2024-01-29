package service

import (
	repository "BotanyBackend/repository"
	"BotanyBackend/servers/g-rpc/Type"
)

type EcoTypeServiceImpl struct {
	repos repository.Repository
}

func NewEcoTypeServiceImpl(repos *repository.Repository) EcoTypeService {
	return EcoTypeServiceImpl{*repos}
}

func (e EcoTypeServiceImpl) Insert(input Type.CreateEcoEntityRequest) (*Type.EcoEntity, error) {

	def := Type.EcoEntity{
		Id:          "2",
		Title:       input.Title,
		Description: input.Description,
		EcoType:     input.EcoType,
	}

	return e.repos.Insert(def)
}

func (e EcoTypeServiceImpl) GetById(inputId string) (Type.EcoEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (e EcoTypeServiceImpl) GetList() (Type.EcoEntityList, error) {
	//TODO implement me
	panic("implement me")
}
