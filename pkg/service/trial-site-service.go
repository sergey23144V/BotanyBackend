package service

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
)

type TrialSiteService interface {
	CreateTrialSite(ctx context.Context, ecomorph *api.InputTrialSiteRequest) (*api.TrialSite, error)
	GetTrialSiteById(ctx context.Context, request *api.IdRequest) (*api.TrialSite, error)
	DeleteTrialSite(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	AddPlant(context.Context, *api.AddPlantTrialSiteRequest) (*api.TrialSite, error)
	StrictUpdateTrialSite(ctx context.Context, in *api.InputTrialSiteRequest) (*api.TrialSite, error)
	UpdateTrialSite(ctx context.Context, in *api.InputTrialSiteRequest) (*api.TrialSite, error)
	GetListTrialSite(ctx context.Context, request *api.PagesRequest) (*api.TrialSiteList, error)
}

type TrialSiteServiceImpl struct {
	repository *repository.Repository
}

func NewTrialSiteServiceImpl(repository *repository.Repository) TrialSiteService {
	return TrialSiteServiceImpl{repository}
}

func (t TrialSiteServiceImpl) CreateTrialSite(ctx context.Context, ecomorph *api.InputTrialSiteRequest) (*api.TrialSite, error) {
	return t.repository.TrialSiteRepository.CreateTrialSite(ctx, t.ToPB(ctx, ecomorph))
}

func (t TrialSiteServiceImpl) GetTrialSiteById(ctx context.Context, request *api.IdRequest) (*api.TrialSite, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return t.repository.TrialSiteRepository.GetTrialSiteById(ctx, &api.TrialSite{Id: request.Id, UserId: userId})
}

func (t TrialSiteServiceImpl) DeleteTrialSite(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := t.repository.TrialSiteRepository.DeleteTrialSite(ctx, &api.TrialSite{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return nil, err
	}
	return result, nil
}

func (t TrialSiteServiceImpl) StrictUpdateTrialSite(ctx context.Context, in *api.InputTrialSiteRequest) (*api.TrialSite, error) {
	return t.repository.TrialSiteRepository.StrictUpdateTrialSite(ctx, t.ToPB(ctx, in))
}

func (t TrialSiteServiceImpl) AddPlant(ctx context.Context, in *api.AddPlantTrialSiteRequest) (*api.TrialSite, error) {
	panic("Not impl")
	//trialSite, err := t.GetTrialSiteById(ctx, &api.IdRequest{Id: in.IdTrialSite})
	//if err != nil {
	//	return nil, err
	//}
	//return t.repository.TrialSiteRepository.UpdateTrialSite(ctx, t.ToAddPlant(in, trialSite), &field_mask.FieldMask{Paths: []string{"TypePlant"}})
}

func (t TrialSiteServiceImpl) UpdateTrialSite(ctx context.Context, in *api.InputTrialSiteRequest) (*api.TrialSite, error) {
	return t.repository.TrialSiteRepository.UpdateTrialSite(ctx, t.ToPB(ctx, in), &field_mask.FieldMask{Paths: []string{"Title",
		"Description", "TypePlant", "SubDominant", "Dominant", "CountTypes", " Rating ", "Covered"}})
}

func (t TrialSiteServiceImpl) GetListTrialSite(ctx context.Context, request *api.PagesRequest) (*api.TrialSiteList, error) {
	var page *api.PagesResponse

	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := t.repository.TrialSiteRepository.GetListTrialSite(ctx, &api.TrialSite{UserId: userId}, request)
	if request != nil {
		page = &api.PagesResponse{Page: request.Page, Limit: request.Limit, Total: int32(len(getList))}
	}
	if err != nil {
		return nil, err
	}
	result := &api.TrialSiteList{List: getList, Page: page}
	return result, nil
}

func (t TrialSiteServiceImpl) ToPB(ctx context.Context, request *api.InputTrialSiteRequest) *api.TrialSite {
	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &api.TrialSite{
		Id:          id,
		Title:       request.Input.Title,
		Covered:     request.Input.Covered,
		Rating:      request.Input.Rating,
		CountTypes:  request.Input.CountTypes,
		Dominant:    request.Input.Dominant,
		SubDominant: request.Input.SubDominant,
		TransectId:  request.Input.TransectId,
		UserId:      userId,
	}
}

//func (t TrialSiteServiceImpl) ToAddPlant(request *api.AddPlantTrialSiteRequest, trialSite *api.TrialSite) *api.TrialSite {
//
//	plants := append(trialSite.TypePlant, &api.TypePlant{Id: request.IdPlant})
//	trialSite.TypePlant = plants
//	return trialSite
//}
