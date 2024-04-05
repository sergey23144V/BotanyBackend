package repository

import (
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"gorm.io/gorm"
)

type Repository struct {
	EcomorphsEntityRepository
	EcomorphRepository
	TypePlantRepository
	TrialSiteRepository
	TransectRepository
	ImgRepositoryImpl
	AnalysisRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		EcomorphsEntityRepository: NewEcomorphsEntityRepositoryImpl(db),
		EcomorphRepository:        NewEcomorphRepositoryImpl(db),
		TypePlantRepository:       NewTypePlantRepositoryImpl(db),
		TrialSiteRepository:       NewTrialSiteRepositoryImpl(db),
		TransectRepository:        NewTransectRepositoryImpl(db),
		ImgRepositoryImpl:         NewImgRepositoryImpl(db),
		AnalysisRepository:        NewAnalysisRepositoryImpl(db),
	}
}

func (t Repository) CountPlant(plantList []*api.Plant) int {
	countPlant := 0
	seen := make(map[string]bool)
	if len(plantList) > 0 {
		for _, item := range plantList {
			if seen[item.TypePlant.Id.ResourceId] {
				continue
			}
			seen[item.TypePlant.Id.ResourceId] = true
			countPlant++
		}
		return countPlant
	}
	return 0
}
