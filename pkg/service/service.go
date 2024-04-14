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
	AnalysisService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		EcomorphsEntityService: NewEcomorphsEntityServiceImpl(repos),
		EcomorphService:        NewEcomorphServiceImpl(repos),
		TypePlantService:       NewTypePlantServiceImpl(repos),
		TrialSiteService:       NewTrialSiteServiceImpl(repos),
		TransectService:        NewTransectServiceImpl(repos),
		ImgService:             NewImgServiceImpl(repos),
		AnalysisService:        NewAnalysisServiceImpl(repos),
	}
}

func RevealBallNumber(coverage int) int32 {
	switch {
	case coverage < 1:
		return 0
	case coverage < 5:
		return 1
	case coverage < 25:
		return 2
	case coverage < 50:
		return 3
	case coverage < 75:
		return 4
	default:
		return 5
	}
}
