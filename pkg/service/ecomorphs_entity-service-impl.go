package service

import (
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
)

type EcomorphsEntityServiceImpl struct {
	repos repository.Repository
}

//func (e EcomorphsEntityServiceImpl) Insert(input g_rpc.CreateEcomorphsEntityRequest) (*g_rpc.EcomorphsEntity, error) {
//
//	def := g_rpc.EcomorphsEntity{
//		Id:          pkg.GenerateUUID(),
//		Title:       input.Title,
//		Description: input.Description,
//	}
//
//	return e.repos.EcomorphsEntityRepository.Insert(def)
//}
//
//func (e EcomorphsEntityServiceImpl) GetById(inputId string) (*g_rpc.EcomorphsEntity, error) {
//	return e.repos.EcomorphsEntityRepository.GetById(inputId)
//}
//
//func (e EcomorphsEntityServiceImpl) GetList() ([]*g_rpc.EcomorphsEntity, error) {
//	return e.repos.EcomorphsEntityRepository.GetList()
//}
//
//func NewEcomorphsEntityServiceImpl(repos *repository.Repository) EcomorphsEntityServiceImpl {
//	return EcomorphsEntityServiceImpl{*repos}
//}
