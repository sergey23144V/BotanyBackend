package service

import (
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
)

type EcomorphServiceImpl struct {
	repos repository.Repository
}

func (e EcomorphServiceImpl) Insert(input g_rpc.CreateEcomorphsRequest) (*g_rpc.Ecomorph, error) {
	def := g_rpc.Ecomorph{
		Id:          pkg.GenerateUUID(),
		Title:       input.Title,
		Description: input.Description,
	}

	return e.repos.EcomorphsRepository.Insert(def)
}

func (e EcomorphServiceImpl) GetById(inputId string) (*g_rpc.Ecomorph, error) {
	return e.repos.EcomorphsRepository.GetById(inputId)
}

func (e EcomorphServiceImpl) GetList() ([]*g_rpc.Ecomorph, error) {
	return e.repos.EcomorphsRepository.GetList()
}

func NewEcomorphServiceImpl(repos *repository.Repository) EcomorphServiceImpl {
	return EcomorphServiceImpl{*repos}
}
