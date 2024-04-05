package implementation

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type TransectServetImpl struct {
	service *service.Service
}

func (t TransectServetImpl) AddTrialSiteToTransect(ctx context.Context, request *api.InputTransectRequest) (*api.Transect, error) {
	//TODO implement me
	panic("implement me")
}

func NewTransectServetImpl(service *service.Service) TransectServetImpl {
	return TransectServetImpl{service}
}

func (t TransectServetImpl) CreateTransect(ctx context.Context, request *api.InputTransectRequest) (*api.Transect, error) {
	return t.service.CreateTransect(ctx, request)
}

func (t TransectServetImpl) GetTransect(ctx context.Context, request *api.IdRequest) (*api.Transect, error) {
	return t.service.TransectService.GetTransectById(ctx, request)
}

func (t TransectServetImpl) UpTransect(ctx context.Context, request *api.InputTransectRequest) (*api.Transect, error) {
	return t.service.TransectService.UpdateTransect(ctx, request)
}

func (t TransectServetImpl) DeleteTransect(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return t.service.TransectService.DeleteTransect(ctx, request)
}

func (t TransectServetImpl) GetAllTransect(ctx context.Context, request *api.PagesRequest) (*api.TransectList, error) {
	return t.service.TransectService.GetListTransect(ctx, request)
}

func (t TransectServetImpl) MustEmbedUnimplementedTransectServiceServer() {
	//TODO implement me
	panic("implement me")
}
