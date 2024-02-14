package service

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	trial_site "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/trial-site"
	"google.golang.org/genproto/protobuf/field_mask"
)

type TrialSiteService interface {
	CreateTrialSite(ctx context.Context, ecomorph *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error)
	GetTrialSiteById(ctx context.Context, request *api.IdRequest) (*trial_site.TrialSite, error)
	DeleteTrialSite(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateTrialSite(ctx context.Context, in *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error)
	UpdateTrialSite(ctx context.Context, in *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error)
	GetListTrialSite(ctx context.Context, request *api.EmptyRequest) (*trial_site.TrialSiteList, error)
}

type TrialSiteServiceImpl struct {
	repository *repository.Repository
}

func NewTrialSiteServiceImpl(repository *repository.Repository) TrialSiteService {
	return TrialSiteServiceImpl{repository}
}

func (t TrialSiteServiceImpl) CreateTrialSite(ctx context.Context, ecomorph *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error) {
	return t.repository.TrialSiteRepository.CreateTrialSite(ctx, t.ToPB(ctx, ecomorph))
}

func (t TrialSiteServiceImpl) GetTrialSiteById(ctx context.Context, request *api.IdRequest) (*trial_site.TrialSite, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return t.repository.TrialSiteRepository.GetTrialSiteById(ctx, &trial_site.TrialSite{Id: request.Id, UserId: userId})
}

func (t TrialSiteServiceImpl) DeleteTrialSite(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := t.repository.TrialSiteRepository.DeleteTrialSite(ctx, &trial_site.TrialSite{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return nil, err
	}
	return result, nil
}

func (t TrialSiteServiceImpl) StrictUpdateTrialSite(ctx context.Context, in *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error) {
	return t.repository.TrialSiteRepository.StrictUpdateTrialSite(ctx, t.ToPB(ctx, in))
}

func (t TrialSiteServiceImpl) UpdateTrialSite(ctx context.Context, in *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error) {
	return t.repository.TrialSiteRepository.UpdateTrialSite(ctx, t.ToPB(ctx, in), &field_mask.FieldMask{Paths: []string{"Title", "Description", "Ecomorphs"}})
}

func (t TrialSiteServiceImpl) GetListTrialSite(ctx context.Context, request *api.EmptyRequest) (*trial_site.TrialSiteList, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := t.repository.TrialSiteRepository.GetListTrialSite(ctx, &trial_site.TrialSite{UserId: userId})
	if err != nil {
		return nil, err
	}
	result := &trial_site.TrialSiteList{TrialSite: getList}
	return result, nil
}

func (t TrialSiteServiceImpl) ToPB(ctx context.Context, request *trial_site.InputTrialSiteRequest) *trial_site.TrialSite {
	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &trial_site.TrialSite{
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
