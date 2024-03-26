package service

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"

	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
)

type TypePlantService interface {
	CreateTypePlant(ctx context.Context, ecomorph *api.InputTypePlantRequest) (*api.TypePlant, error)
	GetTypePlantById(ctx context.Context, request *api.IdRequest) (*api.TypePlant, error)
	DeleteTypePlant(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateTypePlant(ctx context.Context, in *api.InputTypePlantRequest) (*api.TypePlant, error)
	UpdateTypePlant(ctx context.Context, in *api.InputTypePlantRequest) (*api.TypePlant, error)
	AddEcomorphsEntity(ctx context.Context, request *api.InputTypePlant_EcomorphsEntityRequest) (*api.TypePlant, error)
	GetListTypePlant(ctx context.Context, request *api.PagesRequest) (*api.TypePlantList, error)
}

type TypePlantServiceImpl struct {
	repository *repository.Repository
}

func NewTypePlantServiceImpl(repository *repository.Repository) TypePlantServiceImpl {
	return TypePlantServiceImpl{repository}
}

func (t TypePlantServiceImpl) CreateTypePlant(ctx context.Context, ecomorph *api.InputTypePlantRequest) (*api.TypePlant, error) {
	return t.repository.TypePlantRepository.CreateTypePlant(ctx, t.ToPB(ctx, ecomorph))
}

func (t TypePlantServiceImpl) GetTypePlantById(ctx context.Context, request *api.IdRequest) (*api.TypePlant, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return t.repository.TypePlantRepository.GetTypePlantById(ctx, &api.TypePlant{Id: request.Id, UserId: userId})
}

func (t TypePlantServiceImpl) DeleteTypePlant(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := t.repository.TypePlantRepository.DeleteTypePlant(ctx, &api.TypePlant{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return nil, err
	}
	return result, nil
}

func (t TypePlantServiceImpl) StrictUpdateTypePlant(ctx context.Context, in *api.InputTypePlantRequest) (*api.TypePlant, error) {
	return t.repository.TypePlantRepository.StrictUpdateTypePlant(ctx, t.ToPB(ctx, in))
}

func (t TypePlantServiceImpl) UpdateTypePlant(ctx context.Context, in *api.InputTypePlantRequest) (*api.TypePlant, error) {
	fieldMask := []string{}
	if in.Input.Title != "" {
		fieldMask = append(fieldMask, "Title")
	}
	if in.Input.Subtitle != "" {
		fieldMask = append(fieldMask, "Subtitle")
	}
	return t.repository.TypePlantRepository.UpdateTypePlant(ctx, t.ToPB(ctx, in), &field_mask.FieldMask{Paths: fieldMask})
}

func (t TypePlantServiceImpl) AddEcomorphsEntity(ctx context.Context, in *api.InputTypePlant_EcomorphsEntityRequest) (*api.TypePlant, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return t.repository.TypePlantRepository.UpdateTypePlant(ctx, &api.TypePlant{Id: in.Id, EcomorphsEntity: in.EcomorphsEntity, UserId: userId}, &field_mask.FieldMask{Paths: []string{}})
}

func (t TypePlantServiceImpl) GetListTypePlant(ctx context.Context, request *api.PagesRequest) (*api.TypePlantList, error) {
	var page *api.PagesResponse
	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := t.repository.TypePlantRepository.GetListTypePlant(ctx, &api.TypePlant{UserId: userId}, request)
	if request != nil {
		page = &api.PagesResponse{Page: request.Page, Limit: request.Limit, Total: int32(len(getList))}
	}
	if err != nil {
		return nil, err
	}
	result := &api.TypePlantList{List: getList, Page: page}
	return result, nil
}

func (t TypePlantServiceImpl) ToPB(ctx context.Context, request *api.InputTypePlantRequest) *api.TypePlant {
	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &api.TypePlant{
		Id:              id,
		Title:           request.Input.Title,
		Subtitle:        request.Input.Subtitle,
		EcomorphsEntity: request.Input.EcomorphsEntity,
		UserId:          userId,
	}
}
