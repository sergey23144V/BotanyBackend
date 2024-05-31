package service

import (
	context "context"
	er "errors"
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
	GetListEcomorphsEntity(ctx context.Context, request *api.EcomorphsEntityListRequest) (*api.EcomorphsEntityList, error)
}

type EcomorphsEntityServiceImpl struct {
	repository *repository.Repository
}

func NewEcomorphsEntityServiceImpl(repository *repository.Repository) EcomorphsEntityServiceImpl {
	return EcomorphsEntityServiceImpl{repository}
}

func (e EcomorphsEntityServiceImpl) CreateEcomorphsEntity(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error) {
	pb, err := e.ToPB(ctx, entity)
	if err != nil {
		return nil, err
	}
	return e.repository.CreateEcomorphsEntity(ctx, pb)
}

func (e EcomorphsEntityServiceImpl) GetEcomorphsEntityById(ctx context.Context, request *api.IdRequest) (*api.EcomorphsEntity, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return e.repository.EcomorphsEntityRepository.GetEcomorphsEntityById(ctx, &api.EcomorphsEntity{Id: request.Id, UserId: userId})
}

func (e EcomorphsEntityServiceImpl) DeleteEcomorphsEntity(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	role := middlewares.GetUserRoleFromContext(ctx)
	err := e.repository.EcomorphsEntityRepository.DeleteEcomorphsEntity(ctx, &api.EcomorphsEntity{Id: request.Id, UserId: userId}, *role)
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (e EcomorphsEntityServiceImpl) StrictUpdateEcomorphsEntity(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error) {
	pb, err := e.ToPB(ctx, entity)
	if err != nil {
		return nil, err
	}
	return e.repository.EcomorphsEntityRepository.StrictUpdateEcomorphsEntity(ctx, pb)
}

func (e EcomorphsEntityServiceImpl) UpdateEcomorphsEntity(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error) {
	fieldMask := []string{}
	if entity.Input.Title != "" {
		fieldMask = append(fieldMask, "Title")
	}
	if entity.Input.Description != "" {
		fieldMask = append(fieldMask, "Description")
	}
	pb, err := e.ToPB(ctx, entity)
	if err != nil {
		return nil, err
	}
	role := middlewares.GetUserRoleFromContext(ctx)
	return e.repository.EcomorphsEntityRepository.UpdateEcomorphsEntity(ctx, pb, &field_mask.FieldMask{Paths: fieldMask}, *role)
}

func (e EcomorphsEntityServiceImpl) GetListEcomorphsEntity(ctx context.Context, request *api.EcomorphsEntityListRequest) (*api.EcomorphsEntityList, error) {
	var page *api.PagesResponse
	list, err := e.repository.EcomorphsEntityRepository.GetListEcomorphsEntity(ctx, &api.EcomorphsEntity{}, request)
	if err != nil {
		return nil, err
	}
	if request.Page != nil {
		page = &api.PagesResponse{Page: request.Page.Page, Limit: request.Page.Limit, Total: int32(len(list))}
	}
	return &api.EcomorphsEntityList{List: list, Page: page}, nil
}

func (e EcomorphsEntityServiceImpl) ToPB(ctx context.Context, entity *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error) {
	var (
		id             *resource.Identifier
		ecomorphEntity *api.EcomorphsEntity
	)

	if entity.Id != nil {
		id = entity.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}
	userId := middlewares.GetUserIdFromContext(ctx)
	role := middlewares.GetUserRoleFromContext(ctx)
	if *role == api.RoleType_SuperUser {
		ecomorphEntity = &api.EcomorphsEntity{
			Id:           id,
			Title:        entity.Input.Title,
			Description:  entity.Input.Description,
			Ecomorphs:    entity.Input.Ecomorphs,
			Score:        entity.Input.Score,
			DisplayTable: entity.Input.DisplayTable,
		}
	} else if *role != api.RoleType_SuperUser {
		return nil, er.New("has no rights")
	} else {
		ecomorphEntity = &api.EcomorphsEntity{
			Id:           id,
			Title:        entity.Input.Title,
			Description:  entity.Input.Description,
			Ecomorphs:    entity.Input.Ecomorphs,
			Score:        entity.Input.Score,
			DisplayTable: entity.Input.DisplayTable,
			UserId:       userId,
		}
	}
	return ecomorphEntity, nil
}
