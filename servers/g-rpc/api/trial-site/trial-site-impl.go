package trial_site

import (
	context "context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	log "github.com/sirupsen/logrus"
)

type TrialSiteServetImpl struct {
	db *gorm.DB
}

func NewTrialSiteServetImpl(db *gorm.DB) TrialSiteServetImpl {
	return TrialSiteServetImpl{db}
}

func (t TrialSiteServetImpl) CreateTrialSite(ctx context.Context, request *InputTrialSiteRequest) (*TrialSite, error) {
	createEcomorph, err := DefaultCreateTrialSite(ctx, t.ToPB(ctx, request), t.db)
	if err != nil {
		log.Error("Insert TrialSite:", err)
		return nil, err
	}
	log.Debug("Insert TrialSite: good")

	return createEcomorph, nil
}

func (t TrialSiteServetImpl) GetTrialSite(ctx context.Context, request *api.IdRequest) (*TrialSite, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	typePlant, err := DefaultReadTrialSite(ctx, &TrialSite{Id: request.Id, UserId: userId}, t.db)
	if err != nil {
		log.Error("Get TrialSite:", err)
		return nil, err
	}
	log.Debug("Get TrialSite: good")

	return typePlant, nil
}

func (t TrialSiteServetImpl) UpTrialSite(ctx context.Context, request *InputTrialSiteRequest) (*TrialSite, error) {
	return DefaultStrictUpdateTrialSite(ctx, t.ToPB(ctx, request), t.db)
}

func (t TrialSiteServetImpl) DeleteTrialSite(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := DefaultDeleteTrialSite(ctx, &TrialSite{Id: request.Id, UserId: userId}, t.db)
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (t TrialSiteServetImpl) GetAllTrialSite(ctx context.Context, request *api.EmptyRequest) (*TrialSiteList, error) {
	list, err := DefaultListTrialSite(ctx, t.db)
	if err != nil {
		return nil, err
	}
	return &TrialSiteList{TrialSite: list}, nil
}

func (t TrialSiteServetImpl) mustEmbedUnimplementedTrialSiteServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (t TrialSiteServetImpl) ToPB(ctx context.Context, request *InputTrialSiteRequest) *TrialSite {
	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &TrialSite{
		Id:          id,
		Title:       request.Input.Title,
		Covered:     request.Input.Covered,
		Rating:      request.Input.Rating,
		CountTypes:  request.Input.CountTypes,
		Dominant:    request.Input.Dominant,
		SubDominant: request.Input.SubDominant,
		TypePlant:   request.Input.TypePlant,
		UserId:      userId,
	}
}
