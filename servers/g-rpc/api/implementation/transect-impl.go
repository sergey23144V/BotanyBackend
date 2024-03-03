package implementation

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/transect"
)

type TransectServetImpl struct {
	service *service.Service
}

func NewTransectServetImpl(service *service.Service) TransectServetImpl {
	return TransectServetImpl{service}
}

func (t TransectServetImpl) CreateTransect(ctx context.Context, request *transect.InputTransectRequest) (*transect.Transect, error) {
	return t.service.CreateTransect(ctx, request)
}

func (t TransectServetImpl) GetTransect(ctx context.Context, request *api.IdRequest) (*transect.Transect, error) {
	return t.service.TransectService.GetTransectById(ctx, request)
}

func (t TransectServetImpl) UpTransect(ctx context.Context, request *transect.InputTransectRequest) (*transect.Transect, error) {
	return t.service.TransectService.UpdateTransect(ctx, request)
}

func (t TransectServetImpl) DeleteTransect(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return t.service.TransectService.DeleteTransect(ctx, request)
}

func (t TransectServetImpl) GetAllTransect(ctx context.Context, request *api.PagesRequest) (*transect.TransectList, error) {
	return t.service.TransectService.GetListTransect(ctx, request)
}

func (t TransectServetImpl) MustEmbedUnimplementedTransectServiceServer() {
	//TODO implement me
	panic("implement me")
}
