package type_plant

import (
	context "context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	log "github.com/sirupsen/logrus"
)

type TypePlantServetImpl struct {
	db *gorm.DB
}

func NewTypePlantServetImpl(db *gorm.DB) TypePlantServetImpl {
	return TypePlantServetImpl{db}
}

func (t TypePlantServetImpl) CreateTypePlant(ctx context.Context, request *InputTypePlantRequest) (*TypePlant, error) {
	createEcomorph, err := DefaultCreateTypePlant(ctx, t.ToPB(ctx, request), t.db)
	if err != nil {
		log.Error("Insert Ecomorph:", err)
		return nil, err
	}
	log.Debug("Insert Ecomorph: good")

	return createEcomorph, nil
}

func (t TypePlantServetImpl) GetTypePlant(ctx context.Context, request *api.IdRequest) (*TypePlant, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	typePlant, err := DefaultReadTypePlant(ctx, &TypePlant{Id: request.Id, UserId: userId}, t.db)
	if err != nil {
		log.Error("Get Ecomorph:", err)
		return nil, err
	}
	log.Debug("Get Ecomorph: good")

	return typePlant, nil
}

func (t TypePlantServetImpl) UpdateTypePlant(ctx context.Context, request *InputTypePlantRequest) (*TypePlant, error) {
	return DefaultStrictUpdateTypePlant(ctx, t.ToPB(ctx, request), t.db)
}

func (t TypePlantServetImpl) DeleteTypePlant(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := DefaultDeleteTypePlant(ctx, &TypePlant{Id: request.Id, UserId: userId}, t.db)
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (t TypePlantServetImpl) GetAllTypePlant(ctx context.Context, request *api.EmptyRequest) (*TypePlantList, error) {
	list, err := DefaultListTypePlant(ctx, t.db)
	if err != nil {
		return nil, err
	}
	return &TypePlantList{TypePlant: list}, nil
}

func (t TypePlantServetImpl) mustEmbedUnimplementedTypePlantServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (t TypePlantServetImpl) ToPB(ctx context.Context, request *InputTypePlantRequest) *TypePlant {
	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id.ResourceId = pkg.GenerateUUID()
	}
	userId := middlewares.GetUserIdFromContext(ctx)
	return &TypePlant{
		Id:              id,
		Title:           request.Input.Title,
		Subtitle:        request.Input.Subtitle,
		EcomorphsEntity: request.Input.EcomorphsEntity,
		UserId:          userId,
	}
}
