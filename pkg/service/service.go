package service

import (
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
)

type Service struct {
	EcomorphService
	EcomorphsEntityService
	TypePlantService
	TrialSiteService
	TransectService
	ImgService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		EcomorphsEntityService: NewEcomorphsEntityServiceImpl(repos),
		EcomorphService:        NewEcomorphServiceImpl(repos),
		TypePlantService:       NewTypePlantServiceImpl(repos),
		TrialSiteService:       NewTrialSiteServiceImpl(repos),
		TransectService:        NewTransectServiceImpl(repos),
		ImgService:             NewImgServiceImpl(repos),
	}
}
