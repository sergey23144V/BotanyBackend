package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	EcomorphsEntityRepository
	EcomorphRepository
	TypePlantRepository
	TrialSiteRepository
	TransectRepository
	ImgRepositoryImpl
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		EcomorphsEntityRepository: NewEcomorphsEntityRepositoryImpl(db),
		EcomorphRepository:        NewEcomorphRepositoryImpl(db),
		TypePlantRepository:       NewTypePlantRepositoryImpl(db),
		TrialSiteRepository:       NewTrialSiteRepositoryImpl(db),
		TransectRepository:        NewTransectRepositoryImpl(db),
		ImgRepositoryImpl:         NewImgRepositoryImpl(db),
	}
}
