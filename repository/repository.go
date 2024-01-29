package repository

import (
	"BotanyBackend/servers/g-rpc/Type"
	"gorm.io/gorm"
)

type EcoTypeRepository interface {
	Insert(input Type.EcoEntity) (*Type.EcoEntity, error)
	GetById(inputId string) (Type.EcoEntity, error)
	GetList() ([]Type.EcoEntity, error)
}

type Repository struct {
	EcoTypeRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		EcoTypeRepository: NewEcoTypeRepositoryImpl(db),
	}
}
