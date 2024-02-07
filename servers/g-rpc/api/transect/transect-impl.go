package transect

import (
	context "context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	log "github.com/sirupsen/logrus"
)

type TransectServetImpl struct {
	db *gorm.DB
}

func NewTransectServetImpl(db *gorm.DB) TransectServetImpl {
	return TransectServetImpl{db}
}

func (t TransectServetImpl) CreateTransect(ctx context.Context, request *InputTransectRequest) (*Transect, error) {
	createEcomorph, err := DefaultCreateTransect(ctx, t.ToPB(ctx, request), t.db)
	if err != nil {
		log.Error("Insert Transect:", err)
		return nil, err
	}
	log.Debug("Insert Transect: good")

	return createEcomorph, nil
}

func (t TransectServetImpl) GetTransect(ctx context.Context, request *api.IdRequest) (*Transect, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	typePlant, err := DefaultReadTransect(ctx, &Transect{Id: request.Id, UserId: userId}, t.db)
	if err != nil {
		log.Error("Get Transect:", err)
		return nil, err
	}
	log.Debug("Get Transect: good")

	return typePlant, nil
}

func (t TransectServetImpl) UpTransect(ctx context.Context, request *InputTransectRequest) (*Transect, error) {
	return DefaultStrictUpdateTransect(ctx, t.ToPB(ctx, request), t.db)
}

func (t TransectServetImpl) DeleteTransect(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := DefaultDeleteTransect(ctx, &Transect{Id: request.Id, UserId: userId}, t.db)
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (t TransectServetImpl) GetAllTransect(ctx context.Context, request *api.EmptyRequest) (*TransectList, error) {
	list, err := DefaultListTransect(ctx, t.db)
	if err != nil {
		return nil, err
	}
	return &TransectList{Transect: list}, nil
}

func (t TransectServetImpl) mustEmbedUnimplementedTransectServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (t TransectServetImpl) ToPB(ctx context.Context, request *InputTransectRequest) *Transect {

	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)

	return &Transect{
		Id:              id,
		Title:           request.Input.Title,
		Covered:         request.Input.Covered,
		Rating:          request.Input.Rating,
		Square:          request.Input.Square,
		SquareTrialSite: request.Input.SquareTrialSite,
		CountTypes:      request.Input.CountTypes,
		Dominant:        request.Input.Dominant,
		SubDominant:     request.Input.SubDominant,
		TrialSite:       request.Input.TrialSite,
		UserId:          userId,
	}
}
