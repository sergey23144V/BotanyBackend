package service

import (
	context "context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph-entity"
)

type EcomorphsEntityServiceImpl struct {
	repository *repository.Repository
}

func NewEcomorphsEntityServiceImpl(repository *repository.Repository) EcomorphsEntityServiceImpl {
	return EcomorphsEntityServiceImpl{repository}
}

func (e EcomorphsEntityServiceImpl) CreateEcomorphsEntity(ctx context.Context, entity *ecomorph_entity.InputEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
	return e.repository.CreateEcomorphsEntity(ctx, e.ToPB(ctx, entity))
}

func (e EcomorphsEntityServiceImpl) GetEcomorphsEntityById(ctx context.Context, request *api.IdRequest) (*ecomorph_entity.EcomorphsEntity, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return e.repository.EcomorphsEntityRepository.GetEcomorphsEntityById(ctx, &ecomorph_entity.EcomorphsEntity{Id: request.Id, UserId: userId})
}

func (e EcomorphsEntityServiceImpl) DeleteEcomorphsEntity(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := e.repository.EcomorphsEntityRepository.DeleteEcomorphsEntity(ctx, &ecomorph_entity.EcomorphsEntity{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (e EcomorphsEntityServiceImpl) StrictUpdateEcomorphsEntity(ctx context.Context, entity *ecomorph_entity.InputEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
	return e.repository.EcomorphsEntityRepository.UpdateEcomorphsEntity(ctx, e.ToPB(ctx, entity))
}

func (e EcomorphsEntityServiceImpl) UpdateEcomorphsEntity(ctx context.Context, entity *ecomorph_entity.InputEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (e EcomorphsEntityServiceImpl) GetListEcomorphsEntity(ctx context.Context, request *api.EmptyRequest) (*ecomorph_entity.EcomorphsEntityList, error) {
	list, err := e.repository.EcomorphsEntityRepository.GetListEcomorphsEntity(ctx)
	if err != nil {
		return nil, err
	}
	return &ecomorph_entity.EcomorphsEntityList{EcomorphsEntity: list}, nil
}

func (e EcomorphsEntityServiceImpl) ToPB(ctx context.Context, entity *ecomorph_entity.InputEcomorphsEntity) *ecomorph_entity.EcomorphsEntity {
	var id *resource.Identifier

	if entity.Id != nil {
		id = entity.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}
	userId := middlewares.GetUserIdFromContext(ctx)
	return &ecomorph_entity.EcomorphsEntity{
		Id:          id,
		Title:       entity.Input.Title,
		Description: entity.Input.Description,
		Ecomorphs:   entity.Input.Ecomorphs,
		UserId:      userId,
	}
}
