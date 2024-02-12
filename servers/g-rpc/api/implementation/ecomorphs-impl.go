package implementation

import (
	context "context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph"
)

// EcomorphsServetImpl
type EcomorphsServetImpl struct {
	repository *service.Service
}

func NewEcomorphsServetImplImpl(repository *service.Service) EcomorphsServetImpl {
	return EcomorphsServetImpl{repository}
}

func (e EcomorphsServetImpl) UpdateEcomorph(ctx context.Context, in *ecomorph.InputEcomorph) (*ecomorph.Ecomorph, error) {
	return e.repository.StrictUpdateEcomorph(ctx, in)

}

func (e EcomorphsServetImpl) GetListEcomorph(ctx context.Context, request *api.EmptyRequest) (*ecomorph.EcomorphsList, error) {
	return e.repository.GetListEcomorph(ctx, request)
}

func (e EcomorphsServetImpl) InsertEcomorph(ctx context.Context, in *ecomorph.InputEcomorph) (*ecomorph.Ecomorph, error) {
	return e.repository.CreateEcomorph(ctx, in)

}

func (e EcomorphsServetImpl) GetEcomorphById(ctx context.Context, request *api.IdRequest) (*ecomorph.Ecomorph, error) {
	return e.repository.GetEcomorphById(ctx, request)
}

func (e EcomorphsServetImpl) DeleteEcomorphById(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return e.repository.EcomorphService.DeleteEcomorph(ctx, request)

}

func (e EcomorphsServetImpl) MustEmbedUnimplementedEcomorphServiceServer() {
	//TODO implement me
	panic("implement me")
}

func ToPB(ctx context.Context, in *ecomorph.InputEcomorph) *ecomorph.Ecomorph {
	var id *resource.Identifier

	if in.Id != nil {
		id = in.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &ecomorph.Ecomorph{
		Id:          id,
		Title:       in.Payload.Title,
		Description: in.Payload.Description,
		UserId:      userId,
	}
}
