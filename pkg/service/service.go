package service

import (
	"context"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph-entity"
)

type EcomorphService interface {
	CreateEcomorph(ctx context.Context, in *ecomorph.InputEcomorph) (*ecomorph.Ecomorph, error)
	GetEcomorphById(ctx context.Context, request *api.IdRequest) (*ecomorph.Ecomorph, error)
	DeleteEcomorph(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateEcomorph(ctx context.Context, in *ecomorph.InputEcomorph) (*ecomorph.Ecomorph, error)
	UpdateEcomorph(ctx context.Context, in *ecomorph.InputEcomorph) (*ecomorph.Ecomorph, error)
	GetListEcomorph(ctx context.Context, request *api.EmptyRequest) (*ecomorph.EcomorphsList, error)
}

type EcomorphsEntityService interface {
	CreateEcomorphsEntity(ctx context.Context, entity *ecomorph_entity.InputEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error)
	GetEcomorphsEntityById(ctx context.Context, request *api.IdRequest) (*ecomorph_entity.EcomorphsEntity, error)
	DeleteEcomorphsEntity(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateEcomorphsEntity(ctx context.Context, entity *ecomorph_entity.InputEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error)
	UpdateEcomorphsEntity(ctx context.Context, entity *ecomorph_entity.InputEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error)
	GetListEcomorphsEntity(ctx context.Context, request *api.EmptyRequest) (*ecomorph_entity.EcomorphsEntityList, error)
}

type Service struct {
	EcomorphService
	EcomorphsEntityService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		EcomorphsEntityService: NewEcomorphsEntityServiceImpl(repos),
		EcomorphService:        NewEcomorphServiceImpl(repos),
	}
}
