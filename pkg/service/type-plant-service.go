package service

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	type_plant "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/type-plant"
	"google.golang.org/genproto/protobuf/field_mask"

	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
)

type TypePlantService interface {
	CreateTypePlant(ctx context.Context, ecomorph *type_plant.InputTypePlantRequest) (*type_plant.TypePlant, error)
	GetTypePlantById(ctx context.Context, request *api.IdRequest) (*type_plant.TypePlant, error)
	DeleteTypePlant(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateTypePlant(ctx context.Context, in *type_plant.InputTypePlantRequest) (*type_plant.TypePlant, error)
	UpdateTypePlant(ctx context.Context, in *type_plant.InputTypePlantRequest) (*type_plant.TypePlant, error)
	GetListTypePlant(ctx context.Context, request *api.EmptyRequest) (*type_plant.TypePlantList, error)
}

type TypePlantServiceImpl struct {
	repository *repository.Repository
}

func NewTypePlantServiceImpl(repository *repository.Repository) TypePlantServiceImpl {
	return TypePlantServiceImpl{repository}
}

func (t TypePlantServiceImpl) CreateTypePlant(ctx context.Context, ecomorph *type_plant.InputTypePlantRequest) (*type_plant.TypePlant, error) {
	return t.repository.TypePlantRepository.CreateTypePlant(ctx, t.ToPB(ctx, ecomorph))
}

func (t TypePlantServiceImpl) GetTypePlantById(ctx context.Context, request *api.IdRequest) (*type_plant.TypePlant, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return t.repository.TypePlantRepository.GetTypePlantById(ctx, &type_plant.TypePlant{Id: request.Id, UserId: userId})
}

func (t TypePlantServiceImpl) DeleteTypePlant(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := t.repository.TypePlantRepository.DeleteTypePlant(ctx, &type_plant.TypePlant{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return nil, err
	}
	return result, nil
}

func (t TypePlantServiceImpl) StrictUpdateTypePlant(ctx context.Context, in *type_plant.InputTypePlantRequest) (*type_plant.TypePlant, error) {
	return t.repository.TypePlantRepository.StrictUpdateTypePlant(ctx, t.ToPB(ctx, in))
}

func (t TypePlantServiceImpl) UpdateTypePlant(ctx context.Context, in *type_plant.InputTypePlantRequest) (*type_plant.TypePlant, error) {
	return t.repository.TypePlantRepository.UpdateTypePlant(ctx, t.ToPB(ctx, in), &field_mask.FieldMask{Paths: []string{"Title", "Description", "Ecomorphs"}})
}

func (t TypePlantServiceImpl) GetListTypePlant(ctx context.Context, request *api.EmptyRequest) (*type_plant.TypePlantList, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := t.repository.TypePlantRepository.GetListTypePlant(ctx, &type_plant.TypePlant{UserId: userId})
	if err != nil {
		return nil, err
	}
	result := &type_plant.TypePlantList{TypePlant: getList}
	return result, nil
}

func (t TypePlantServiceImpl) ToPB(ctx context.Context, request *type_plant.InputTypePlantRequest) *type_plant.TypePlant {
	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &type_plant.TypePlant{
		Id:              id,
		Title:           request.Input.Title,
		Subtitle:        request.Input.Subtitle,
		EcomorphsEntity: request.Input.EcomorphsEntity,
		UserId:          userId,
	}
}
