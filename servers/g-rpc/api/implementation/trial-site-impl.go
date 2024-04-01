package implementation

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type TrialSiteServetImpl struct {
	service *service.Service
}

func NewTrialSiteServetImpl(service *service.Service) TrialSiteServetImpl {
	return TrialSiteServetImpl{service}
}

func (t TrialSiteServetImpl) CreateTrialSite(ctx context.Context, request *api.InputTrialSiteRequest) (*api.TrialSite, error) {
	return t.service.CreateTrialSite(ctx, request)
}

func (t TrialSiteServetImpl) GetTrialSite(ctx context.Context, request *api.IdRequest) (*api.TrialSite, error) {
	return t.service.TrialSiteService.GetTrialSiteById(ctx, request)
}

func (t TrialSiteServetImpl) UpTrialSite(ctx context.Context, request *api.InputTrialSiteRequest) (*api.TrialSite, error) {
	return t.service.TrialSiteService.UpdateTrialSite(ctx, request)
}

func (t TrialSiteServetImpl) DeleteTrialSite(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return t.service.TrialSiteService.DeleteTrialSite(ctx, request)
}

func (t TrialSiteServetImpl) GetAllTrialSite(ctx context.Context, request *api.PagesRequest) (*api.TrialSiteList, error) {
	return t.service.TrialSiteService.GetListTrialSite(ctx, request)
}

func (t TrialSiteServetImpl) CreatePlant(ctx context.Context, request *api.InputPlantRequest) (*api.Plant, error) {
	return t.service.CreatePlant(ctx, request)
}
func (t TrialSiteServetImpl) GetPlant(ctx context.Context, request *api.IdRequest) (*api.Plant, error) {
	return t.service.TrialSiteService.GetPlant(ctx, request)
}
func (t TrialSiteServetImpl) UpdatePlant(ctx context.Context, request *api.InputPlantRequest) (*api.Plant, error) {
	return t.service.TrialSiteService.UpdatePlant(ctx, request)
}
func (t TrialSiteServetImpl) DeletePlant(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return t.service.TrialSiteService.DeletePlant(ctx, request)
}
func (t TrialSiteServetImpl) GetAllPlant(ctx context.Context, request *api.PagesRequest) (*api.PlantList, error) {
	return t.service.TrialSiteService.GetAllPlant(ctx, request)
}

func (t TrialSiteServetImpl) MustEmbedUnimplementedTrialSiteServiceServer() {
	//TODO implement me
	panic("implement me")
}
