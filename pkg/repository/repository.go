package repository

import (
	"gorm.io/gorm"
)

//type EcomorphsEntityRepository interface {
//	Insert(input g_rpc.EcomorphsEntity) (*g_rpc.EcomorphsEntity, error)
//	GetById(inputId string) (*g_rpc.EcomorphsEntity, error)
//	GetList() ([]*g_rpc.EcomorphsEntity, error)
//}
//
//type EcomorphsRepository interface {
//	Insert(input g_rpc.Ecomorph) (*g_rpc.Ecomorph, error)
//	GetById(inputId string) (*g_rpc.Ecomorph, error)
//	GetList() ([]*g_rpc.Ecomorph, error)
//}

type Repository struct {
	//EcomorphsEntityRepository
	//EcomorphsRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		//EcomorphsEntityRepository: NewEcomorphsEntityRepositoryImpl(db),
		//EcomorphsRepository:       NewEcomorphsRepositoryImpl(db),
	}
}
