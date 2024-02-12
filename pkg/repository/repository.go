package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	EcomorphsEntityRepository
	EcomorphRepository
	TypePlantRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		EcomorphsEntityRepository: NewEcomorphsEntityRepositoryImpl(db),
		EcomorphRepository:        NewEcomorphRepositoryImpl(db),
		TypePlantRepository:       NewTypePlantRepositoryImpl(db),
	}
}
