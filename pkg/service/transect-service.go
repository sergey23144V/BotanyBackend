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

type TransectService interface {
	CreateTransect(ctx context.Context, ecomorph *api.InputTransectRequest) (*api.Transect, error)
	GetTransectById(ctx context.Context, request *api.IdRequest) (*api.Transect, error)
	DeleteTransect(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateTransect(ctx context.Context, in *api.InputTransectRequest) (*api.Transect, error)
	UpdateTransect(ctx context.Context, in *api.InputTransectRequest) (*api.Transect, error)
	GetListTransect(ctx context.Context, request *api.PagesRequest) (*api.TransectList, error)
}

type TransectServiceImpl struct {
	repository *repository.Repository
}

func NewTransectServiceImpl(repository *repository.Repository) TransectService {
	return TransectServiceImpl{repository}
}

func (t TransectServiceImpl) CreateTransect(ctx context.Context, ecomorph *api.InputTransectRequest) (*api.Transect, error) {
	return t.repository.TransectRepository.CreateTransect(ctx, t.ToPB(ctx, ecomorph))
}

func (t TransectServiceImpl) GetTransectById(ctx context.Context, request *api.IdRequest) (*api.Transect, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return t.repository.TransectRepository.GetTransectById(ctx, &api.Transect{Id: request.Id, UserId: userId})
}

func (t TransectServiceImpl) DeleteTransect(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := t.repository.TransectRepository.DeleteTransect(ctx, &api.Transect{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return nil, err
	}
	return result, nil
}

func (t TransectServiceImpl) StrictUpdateTransect(ctx context.Context, in *api.InputTransectRequest) (*api.Transect, error) {
	return t.repository.TransectRepository.StrictUpdateTransect(ctx, t.ToPB(ctx, in))
}

func (t TransectServiceImpl) UpdateTransect(ctx context.Context, in *api.InputTransectRequest) (*api.Transect, error) {
	return t.repository.TransectRepository.UpdateTransect(ctx, t.ToPB(ctx, in), &field_mask.FieldMask{Paths: []string{"Title",
		"TrialSite", "SubDominant", "Dominant", "CountTypes", "SquareTrialSite", "Square", "Rating", "Covered"}})
}

func (t TransectServiceImpl) GetListTransect(ctx context.Context, request *api.PagesRequest) (*api.TransectList, error) {
	var page *api.PagesResponse

	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := t.repository.TransectRepository.GetListTransect(ctx, &api.Transect{UserId: userId}, request)
	if request != nil {
		page = &api.PagesResponse{Page: request.Page, Limit: request.Limit, Total: int32(len(getList))}
	}
	if err != nil {
		return nil, err
	}
	result := &api.TransectList{List: getList, Page: page}
	return result, nil
}

func (t TransectServiceImpl) ToPB(ctx context.Context, request *api.InputTransectRequest) *api.Transect {

	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)

	return &api.Transect{
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
