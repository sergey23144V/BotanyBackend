package implementation

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/type-plant"
)

type TypePlantServetImpl struct {
	service *service.Service
}

func NewTypePlantServetImpl(service *service.Service) TypePlantServetImpl {
	return TypePlantServetImpl{service}
}

func (t TypePlantServetImpl) CreateTypePlant(ctx context.Context, request *type_plant.InputTypePlantRequest) (*type_plant.TypePlant, error) {
	return t.service.CreateTypePlant(ctx, request)
}

func (t TypePlantServetImpl) GetTypePlant(ctx context.Context, request *api.IdRequest) (*type_plant.TypePlant, error) {
	return t.service.TypePlantService.GetTypePlantById(ctx, request)
}

func (t TypePlantServetImpl) UpdateTypePlant(ctx context.Context, request *type_plant.InputTypePlantRequest) (*type_plant.TypePlant, error) {
	return t.service.TypePlantService.UpdateTypePlant(ctx, request)
}

func (t TypePlantServetImpl) DeleteTypePlant(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return t.service.TypePlantService.DeleteTypePlant(ctx, request)
}

func (t TypePlantServetImpl) GetAllTypePlant(ctx context.Context, request *api.PagesRequest) (*type_plant.TypePlantList, error) {
	return t.service.TypePlantService.GetListTypePlant(ctx, request)
}

func (t TypePlantServetImpl) MustEmbedUnimplementedTypePlantServiceServer() {
	//TODO implement me
	panic("implement me")
}
