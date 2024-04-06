package implementation

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

// EcomorphsServetImpl
type EcomorphsServetImpl struct {
	service *service.Service
}

func NewEcomorphsServetImplImpl(repository *service.Service) api.EcomorphServiceServer {
	return EcomorphsServetImpl{repository}
}

func (e EcomorphsServetImpl) UpdateEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error) {
	return e.service.UpdateEcomorph(ctx, in)

}

func (e EcomorphsServetImpl) GetListEcomorph(ctx context.Context, request *api.EcomorphListRequest) (*api.EcomorphsList, error) {
	return e.service.GetListEcomorph(ctx, request)
}

func (e EcomorphsServetImpl) InsertEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error) {
	return e.service.CreateEcomorph(ctx, in)

}

func (e EcomorphsServetImpl) GetEcomorphById(ctx context.Context, request *api.IdRequest) (*api.Ecomorph, error) {
	return e.service.GetEcomorphById(ctx, request)
}

func (e EcomorphsServetImpl) DeleteEcomorphById(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return e.service.EcomorphService.DeleteEcomorph(ctx, request)

}

func (e EcomorphsServetImpl) MustEmbedUnimplementedEcomorphServiceServer() {
	//TODO implement me
	panic("implement me")
}
