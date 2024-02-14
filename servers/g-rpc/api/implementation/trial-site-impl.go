package implementation

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/trial-site"
)

type TrialSiteServetImpl struct {
	service *service.Service
}

func NewTrialSiteServetImpl(service *service.Service) TrialSiteServetImpl {
	return TrialSiteServetImpl{service}
}

func (t TrialSiteServetImpl) CreateTrialSite(ctx context.Context, request *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error) {
	return t.service.CreateTrialSite(ctx, request)
}

func (t TrialSiteServetImpl) GetTrialSite(ctx context.Context, request *api.IdRequest) (*trial_site.TrialSite, error) {
	return t.service.TrialSiteService.GetTrialSiteById(ctx, request)
}

func (t TrialSiteServetImpl) UpTrialSite(ctx context.Context, request *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error) {
	return t.service.TrialSiteService.UpdateTrialSite(ctx, request)
}

func (t TrialSiteServetImpl) DeleteTrialSite(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	return t.service.TrialSiteService.DeleteTrialSite(ctx, request)
}

func (t TrialSiteServetImpl) GetAllTrialSite(ctx context.Context, request *api.EmptyRequest) (*trial_site.TrialSiteList, error) {
	return t.service.TrialSiteService.GetListTrialSite(ctx, request)
}

func (t TrialSiteServetImpl) MustEmbedUnimplementedTrialSiteServiceServer() {
	//TODO implement me
	panic("implement me")
}
