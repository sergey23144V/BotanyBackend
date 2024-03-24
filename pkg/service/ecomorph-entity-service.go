package service

import (
	context "context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
)

type EcomorphsEntityService interface {
	CreateEcomorphsEntity(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error)
	GetEcomorphsEntityById(ctx context.Context, request *api.IdRequest) (*api.EcomorphsEntity, error)
	DeleteEcomorphsEntity(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateEcomorphsEntity(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error)
	UpdateEcomorphsEntity(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error)
	GetListEcomorphsEntity(ctx context.Context, request *api.PagesRequest) (*api.EcomorphsEntityList, error)
}

type EcomorphsEntityServiceImpl struct {
	repository *repository.Repository
}

func NewEcomorphsEntityServiceImpl(repository *repository.Repository) EcomorphsEntityServiceImpl {
	return EcomorphsEntityServiceImpl{repository}
}

func (e EcomorphsEntityServiceImpl) CreateEcomorphsEntity(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error) {
	return e.repository.CreateEcomorphsEntity(ctx, e.ToPB(ctx, entity))
}

func (e EcomorphsEntityServiceImpl) GetEcomorphsEntityById(ctx context.Context, request *api.IdRequest) (*api.EcomorphsEntity, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return e.repository.EcomorphsEntityRepository.GetEcomorphsEntityById(ctx, &api.EcomorphsEntity{Id: request.Id, UserId: userId})
}

func (e EcomorphsEntityServiceImpl) DeleteEcomorphsEntity(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := e.repository.EcomorphsEntityRepository.DeleteEcomorphsEntity(ctx, &api.EcomorphsEntity{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (e EcomorphsEntityServiceImpl) StrictUpdateEcomorphsEntity(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error) {
	return e.repository.EcomorphsEntityRepository.StrictUpdateEcomorphsEntity(ctx, e.ToPB(ctx, entity))
}

func (e EcomorphsEntityServiceImpl) UpdateEcomorphsEntity(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error) {
	fieldMask := []string{}
	if entity.Input.Title != "" {
		fieldMask = append(fieldMask, "Title")
	}
	if entity.Input.Description != "" {
		fieldMask = append(fieldMask, "Description")
	}
	return e.repository.EcomorphsEntityRepository.UpdateEcomorphsEntity(ctx, e.ToPB(ctx, entity), &field_mask.FieldMask{Paths: fieldMask})
}

func (e EcomorphsEntityServiceImpl) GetListEcomorphsEntity(ctx context.Context, request *api.PagesRequest) (*api.EcomorphsEntityList, error) {
	var page *api.PagesResponse
	list, err := e.repository.EcomorphsEntityRepository.GetListEcomorphsEntity(ctx, &api.EcomorphsEntity{UserId: middlewares.GetUserIdFromContext(ctx)}, request)
	if err != nil {
		return nil, err
	}
	if request != nil {
		page = &api.PagesResponse{Page: request.Page, Limit: request.Limit, Total: int32(len(list))}
	}
	return &api.EcomorphsEntityList{List: list, Page: page}, nil
}

func (e EcomorphsEntityServiceImpl) ToPB(ctx context.Context, entity *api.InputEcomorphsEntity) *api.EcomorphsEntity {
	var id *resource.Identifier

	if entity.Id != nil {
		id = entity.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}
	userId := middlewares.GetUserIdFromContext(ctx)
	return &api.EcomorphsEntity{
		Id:          id,
		Title:       entity.Input.Title,
		Description: entity.Input.Description,
		Ecomorphs:   entity.Input.Ecomorphs,
		UserId:      userId,
	}
}
