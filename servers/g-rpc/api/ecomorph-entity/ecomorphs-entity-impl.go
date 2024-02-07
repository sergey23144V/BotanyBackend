package ecomorph_entity

import (
	context "context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type EcomorphsEntityServetImpl struct {
	db *gorm.DB
}

func NewEcomorphsEntityServetImpl(db *gorm.DB) EcomorphsEntityServetImpl {
	return EcomorphsEntityServetImpl{db}
}

func (e EcomorphsEntityServetImpl) InsertEcomorphEntity(ctx context.Context, entity *InputEcomorphsEntity) (*EcomorphsEntity, error) {
	return DefaultCreateEcomorphsEntity(ctx, e.ToPB(ctx, entity), e.db)
}

func (e EcomorphsEntityServetImpl) UpdateEcomorphEntity(ctx context.Context, entity *InputEcomorphsEntity) (*EcomorphsEntity, error) {
	return DefaultStrictUpdateEcomorphsEntity(ctx, e.ToPB(ctx, entity), e.db)
}

func (e EcomorphsEntityServetImpl) GetEcomorphEntityByID(ctx context.Context, request *api.IdRequest) (*EcomorphsEntity, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return DefaultReadEcomorphsEntity(ctx, &EcomorphsEntity{Id: request.Id, UserId: userId}, e.db)
}

func (e EcomorphsEntityServetImpl) DeleteEcomorphEntityByID(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := DefaultDeleteEcomorphsEntity(ctx, &EcomorphsEntity{Id: request.Id, UserId: userId}, e.db)
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
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

func (e EcomorphsEntityServetImpl) ToPB(ctx context.Context, entity *InputEcomorphsEntity) *EcomorphsEntity {
	var id *resource.Identifier

	if entity.Id != nil {
		id = entity.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}
	userId := middlewares.GetUserIdFromContext(ctx)
	return &EcomorphsEntity{
		Id:          id,
		Title:       entity.Input.Title,
		Description: entity.Input.Description,
		Ecomorphs:   entity.Input.Ecomorphs,
		UserId:      userId,
	}
}
