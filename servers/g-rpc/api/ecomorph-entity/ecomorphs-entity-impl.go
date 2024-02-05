package ecomorph_entity

import (
	context "context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type EcomorphsEntityServetImpl struct {
	db *gorm.DB
}

func (e EcomorphsEntityServetImpl) InsertEcomorphEntity(ctx context.Context, request *CreateEcomorphsEntityRequest) (*EcomorphsEntity, error) {
	return DefaultCreateEcomorphsEntity(ctx, &EcomorphsEntity{
		Id:          &resource.Identifier{ResourceId: pkg.GenerateUUID()},
		Title:       request.Title,
		Description: request.Description,
		Ecomorphs:   request.EcomorphsId,
	}, e.db)
}

func (e EcomorphsEntityServetImpl) GetEcomorphEntityByID(ctx context.Context, request *api.IdRequest) (*EcomorphsEntity, error) {
	return DefaultReadEcomorphsEntity(ctx, &EcomorphsEntity{Id: request.Id}, e.db)
}

func (e EcomorphsEntityServetImpl) UpEcomorphEntityByID(ctx context.Context, request *UpdateEcomorphsEntityRequest) (*EcomorphsEntity, error) {
	return DefaultStrictUpdateEcomorphsEntity(ctx, &EcomorphsEntity{
		Title:       request.Title,
		Description: request.Description,
		Ecomorphs:   request.EcomorphsId,
	}, e.db)
}

func (e EcomorphsEntityServetImpl) DeleteEcomorphEntityByID(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e EcomorphsEntityServetImpl) GetAllEcomorphEntity(ctx context.Context, request *api.EmptyRequest) (*EcomorphsEntityList, error) {
	list, err := DefaultListEcomorphsEntity(ctx, e.db)
	if err != nil {
		return nil, err
	}
	return &EcomorphsEntityList{EcomorphsEntity: list}, nil
}

func (e EcomorphsEntityServetImpl) mustEmbedUnimplementedEcomorphEntityServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewEcomorphsEntityServetImpl(db *gorm.DB) EcomorphsEntityServetImpl {
	return EcomorphsEntityServetImpl{db}
}
