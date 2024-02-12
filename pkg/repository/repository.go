package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	EcomorphsEntityRepository
	EcomorphRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		EcomorphsEntityRepository: NewEcomorphsEntityRepositoryImpl(db),
		EcomorphRepository:        NewEcomorphRepositoryImpl(db),
	}
}
