package implementation

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type TypePlantServetImpl struct {
	service *service.Service
}

func (t TypePlantServetImpl) AddEcomorphsEntityToTypePlant(ctx context.Context, request *api.InputTypePlant_EcomorphsEntityRequest) (*api.TypePlant, error) {
	//TODO implement me
	panic("implement me")
}

func NewTypePlantServetImpl(service *service.Service) TypePlantServetImpl {
	return TypePlantServetImpl{service}
}

func (t TypePlantServetImpl) CreateTypePlant(ctx context.Context, request *api.InputTypePlantRequest) (*api.TypePlant, error) {
	return t.service.CreateTypePlant(ctx, request)
}

func (t TypePlantServetImpl) GetTypePlant(ctx context.Context, request *api.IdRequest) (*api.TypePlant, error) {
	return t.service.TypePlantService.GetTypePlantById(ctx, request)
}

func (t TypePlantServetImpl) UpdateTypePlant(ctx context.Context, request *api.InputTypePlantRequest) (*api.TypePlant, error) {
	return t.service.TypePlantService.UpdateTypePlant(ctx, request)
}

func (t TypePlantServetImpl) AddEcomorphsEntity(ctx context.Context, request *api.InputTypePlant_EcomorphsEntityRequest) (*api.TypePlant, error) {
	return t.service.TypePlantService.AddEcomorphsEntity(ctx, request)
}

func (t TypePlantServetImpl) DeleteTypePlant(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return t.service.TypePlantService.DeleteTypePlant(ctx, request)
}

func (t TypePlantServetImpl) GetAllTypePlant(ctx context.Context, request *api.PagesRequest) (*api.TypePlantList, error) {
	return t.service.TypePlantService.GetListTypePlant(ctx, request)
}

func (t TypePlantServetImpl) MustEmbedUnimplementedTypePlantServiceServer() {
	//TODO implement me
	panic("implement me")
}
