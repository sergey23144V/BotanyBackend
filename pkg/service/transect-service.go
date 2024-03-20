package service

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/transect"
	"google.golang.org/genproto/protobuf/field_mask"
)

type TransectService interface {
	CreateTransect(ctx context.Context, ecomorph *transect.InputTransectRequest) (*transect.Transect, error)
	GetTransectById(ctx context.Context, request *api.IdRequest) (*transect.Transect, error)
	DeleteTransect(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateTransect(ctx context.Context, in *transect.InputTransectRequest) (*transect.Transect, error)
	UpdateTransect(ctx context.Context, in *transect.InputTransectRequest) (*transect.Transect, error)
	GetListTransect(ctx context.Context, request *api.PagesRequest) (*transect.TransectList, error)
}

type TransectServiceImpl struct {
	repository *repository.Repository
}

func NewTransectServiceImpl(repository *repository.Repository) TransectService {
	return TransectServiceImpl{repository}
}

func (t TransectServiceImpl) CreateTransect(ctx context.Context, ecomorph *transect.InputTransectRequest) (*transect.Transect, error) {
	return t.repository.TransectRepository.CreateTransect(ctx, t.ToPB(ctx, ecomorph))
}

func (t TransectServiceImpl) GetTransectById(ctx context.Context, request *api.IdRequest) (*transect.Transect, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return t.repository.TransectRepository.GetTransectById(ctx, &transect.Transect{Id: request.Id, UserId: userId})
}

func (t TransectServiceImpl) DeleteTransect(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := t.repository.TransectRepository.DeleteTransect(ctx, &transect.Transect{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return nil, err
	}
	return result, nil
}

func (t TransectServiceImpl) StrictUpdateTransect(ctx context.Context, in *transect.InputTransectRequest) (*transect.Transect, error) {
	return t.repository.TransectRepository.StrictUpdateTransect(ctx, t.ToPB(ctx, in))
}

func (t TransectServiceImpl) UpdateTransect(ctx context.Context, in *transect.InputTransectRequest) (*transect.Transect, error) {
	return t.repository.TransectRepository.UpdateTransect(ctx, t.ToPB(ctx, in), &field_mask.FieldMask{Paths: []string{"Title",
		"TrialSite", "SubDominant", "Dominant", "CountTypes", "SquareTrialSite", "Square", "Rating", "Covered"}})
}

func (t TransectServiceImpl) GetListTransect(ctx context.Context, request *api.PagesRequest) (*transect.TransectList, error) {
	var page *api.PagesResponse

	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := t.repository.TransectRepository.GetListTransect(ctx, &transect.Transect{UserId: userId}, request)
	if request != nil {
		page = &api.PagesResponse{Page: request.Page, Limit: request.Limit, Total: int32(len(getList))}
	}
	if err != nil {
		return nil, err
	}
	result := &transect.TransectList{List: getList, Page: page}
	return result, nil
}

func (t TransectServiceImpl) ToPB(ctx context.Context, request *transect.InputTransectRequest) *transect.Transect {

	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)

	return &transect.Transect{
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
