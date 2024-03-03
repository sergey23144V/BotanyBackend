package implementation

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph-entity"
)

type EcomorphsEntityServetImpl struct {
	service *service.Service
}

func NewEcomorphsEntityServetImpl(repository *service.Service) EcomorphsEntityServetImpl {
	return EcomorphsEntityServetImpl{repository}
}

func (e EcomorphsEntityServetImpl) InsertEcomorphEntity(ctx context.Context, entity *ecomorph_entity.InputEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
	return e.service.CreateEcomorphsEntity(ctx, entity)
}

func (e EcomorphsEntityServetImpl) UpdateEcomorphEntity(ctx context.Context, entity *ecomorph_entity.InputEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
	return e.service.EcomorphsEntityService.UpdateEcomorphsEntity(ctx, entity)
}

func (e EcomorphsEntityServetImpl) GetEcomorphEntityByID(ctx context.Context, request *api.IdRequest) (*ecomorph_entity.EcomorphsEntity, error) {
	return e.service.EcomorphsEntityService.GetEcomorphsEntityById(ctx, request)
}

func (e EcomorphsEntityServetImpl) DeleteEcomorphEntityByID(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return e.service.EcomorphsEntityService.DeleteEcomorphsEntity(ctx, request)
}

func (e EcomorphsEntityServetImpl) GetAllEcomorphEntity(ctx context.Context, request *api.PagesRequest) (*ecomorph_entity.EcomorphsEntityList, error) {
	return e.service.EcomorphsEntityService.GetListEcomorphsEntity(ctx, request)
}

func (e EcomorphsEntityServetImpl) MustEmbedUnimplementedEcomorphEntityServiceServer() {
	//TODO implement me
	panic("implement me")
}
