package service

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	trial_site "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/trial-site"
	type_plant "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/type-plant"
	"google.golang.org/genproto/protobuf/field_mask"
)

type TrialSiteService interface {
	CreateTrialSite(ctx context.Context, ecomorph *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error)
	GetTrialSiteById(ctx context.Context, request *api.IdRequest) (*trial_site.TrialSite, error)
	DeleteTrialSite(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	AddPlant(context.Context, *trial_site.AddPlantTrialSiteRequest) (*trial_site.TrialSite, error)
	StrictUpdateTrialSite(ctx context.Context, in *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error)
	UpdateTrialSite(ctx context.Context, in *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error)
	GetListTrialSite(ctx context.Context, request *api.PagesRequest) (*trial_site.TrialSiteList, error)
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

func (t TrialSiteServiceImpl) AddPlant(ctx context.Context, in *trial_site.AddPlantTrialSiteRequest) (*trial_site.TrialSite, error) {
	trialSite, err := t.GetTrialSiteById(ctx, &api.IdRequest{Id: in.Id})
	if err != nil {
		return nil, err
	}
	return t.repository.TrialSiteRepository.UpdateTrialSite(ctx, t.ToAddPlant(in, trialSite), &field_mask.FieldMask{Paths: []string{"TypePlant"}})
}

func (t TrialSiteServiceImpl) UpdateTrialSite(ctx context.Context, in *trial_site.InputTrialSiteRequest) (*trial_site.TrialSite, error) {
	return t.repository.TrialSiteRepository.UpdateTrialSite(ctx, t.ToPB(ctx, in), &field_mask.FieldMask{Paths: []string{"Title",
		"Description", "TypePlant", "SubDominant", "Dominant", "CountTypes", " Rating ", "Covered"}})
}

func (t TrialSiteServiceImpl) GetListTrialSite(ctx context.Context, request *api.PagesRequest) (*trial_site.TrialSiteList, error) {
	var page *api.PagesResponse

	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := t.repository.TrialSiteRepository.GetListTrialSite(ctx, &trial_site.TrialSite{UserId: userId}, request)
	if request != nil {
		page = &api.PagesResponse{Page: request.Page, Limit: request.Limit, Total: int32(len(getList))}
	}
	if err != nil {
		return nil, err
	}
	result := &trial_site.TrialSiteList{List: getList, Page: page}
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
		TransectId:  request.Input.TransectId,
		UserId:      userId,
	}
}

func (t TrialSiteServiceImpl) ToAddPlant(request *trial_site.AddPlantTrialSiteRequest, trialSite *trial_site.TrialSite) *trial_site.TrialSite {

	plants := append(trialSite.TypePlant, &type_plant.TypePlant{Id: request.IdPlant})
	trialSite.TypePlant = plants
	return trialSite
}
