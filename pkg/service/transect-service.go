package service

import (
	"context"
	"errors"
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
	AddTrialSiteToTransect(context.Context, *api.InputTransectRequest) (*api.Transect, error)
	UpdateTransect(ctx context.Context, in *api.InputTransectRequest) (*api.Transect, error)
	GetListTransect(ctx context.Context, request *api.TransectListRequest) (*api.TransectList, error)
	DetectionPlant(ctx context.Context, in *api.Transect) (*api.Transect, error)
}

type TransectServiceImpl struct {
	repository *repository.Repository
}

func (t TransectServiceImpl) AddTrialSiteToTransect(ctx context.Context, request *api.InputTransectRequest) (*api.Transect, error) {
	transect, err := t.GetTransectById(ctx, &api.IdRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}
	if request.Input.TrialSite != nil {
		for _, plant := range request.Input.TrialSite {
			transect.TrialSite = append(transect.TrialSite, plant)
		}
		return t.repository.StrictUpdateTransect(ctx, transect)
	} else {
		return nil, errors.New("not have plant")
	}
}

func NewTransectServiceImpl(repository *repository.Repository) TransectService {
	return TransectServiceImpl{repository}
}

func (t TransectServiceImpl) CreateTransect(ctx context.Context, input *api.InputTransectRequest) (*api.Transect, error) {
	transect, err := t.repository.TransectRepository.CreateTransect(ctx, t.ToPB(ctx, input))
	if err != nil {
		return nil, err
	}
	return t.DetectionPlant(ctx, transect)
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

func (t TransectServiceImpl) DetectionPlant(ctx context.Context, in *api.Transect) (*api.Transect, error) {
	var (
		typePlant []*api.Plant
	)
	output, err := t.repository.TransectRepository.GetTransectByIdForAnalysis(ctx, in)
	if err != nil {
		return in, err
	}
	trialSites := output.TrialSite
	countsDominants := make(map[string]int)

	for _, item := range trialSites {
		for _, plant := range item.Plant {
			if plant.TypePlant != nil {
				typePlant = append(typePlant, plant)
				countsDominants[plant.TypePlant.Id.ResourceId] += int(plant.Coverage)
			}
		}
	}
	var maxCoverage, secondMaxCoverage int
	var maxResourceId, secondMaxResourceId string

	for resourceId, coverage := range countsDominants {
		if coverage > maxCoverage {
			secondMaxCoverage = maxCoverage
			secondMaxResourceId = maxResourceId

			maxCoverage = coverage
			maxResourceId = resourceId
		} else if coverage > secondMaxCoverage {
			secondMaxCoverage = coverage
			secondMaxResourceId = resourceId
		}
	}
	if maxResourceId != "" {
		output.Dominant = &api.TypePlant{Id: &resource.Identifier{ResourceId: maxResourceId}}
	}
	if secondMaxResourceId != "" {
		output.SubDominant = &api.TypePlant{Id: &resource.Identifier{ResourceId: secondMaxResourceId}}
	}
	output.Rating = RevealBallNumber(int(output.Covered))
	output.CountTypes = int32(t.repository.CountPlant(typePlant))

	return t.repository.StrictUpdateTransect(ctx, output)
}

func (t TransectServiceImpl) StrictUpdateTransect(ctx context.Context, in *api.InputTransectRequest) (*api.Transect, error) {
	return t.repository.TransectRepository.StrictUpdateTransect(ctx, t.ToPB(ctx, in))
}

func (t TransectServiceImpl) UpdateTransect(ctx context.Context, in *api.InputTransectRequest) (*api.Transect, error) {
	transect, err := t.repository.TransectRepository.UpdateTransect(ctx, t.ToPB(ctx, in), &field_mask.FieldMask{Paths: t.getMask(in.Input)})
	if err != nil {
		return nil, err
	}
	return t.DetectionPlant(ctx, transect)
}

func (t TransectServiceImpl) GetListTransect(ctx context.Context, request *api.TransectListRequest) (*api.TransectList, error) {
	var page *api.PagesResponse
	if request == nil {
		request = &api.TransectListRequest{}
	}
	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := t.repository.TransectRepository.GetListTransect(ctx, &api.Transect{UserId: userId}, request)
	if request.Page != nil {
		page = &api.PagesResponse{Page: request.Page.Page, Limit: request.Page.Limit, Total: int32(len(getList))}
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
		Img:             request.Input.Img,
		UserId:          userId,
	}
}

func (t TransectServiceImpl) getMask(in *api.InputFormTransectRequest) []string {
	fieldMask := []string{}
	if in.Title != "" {
		fieldMask = append(fieldMask, "Title")
	}
	if in.Covered != 0 {
		fieldMask = append(fieldMask, "Covered")
	}
	if in.SquareTrialSite != 0 {
		fieldMask = append(fieldMask, "SquareTrialSite")
	}
	if in.Square != 0 {
		fieldMask = append(fieldMask, "Square")
	}
	if in.Rating != 0 {
		fieldMask = append(fieldMask, "Rating")
	}
	if in.CountTypes != 0 {
		fieldMask = append(fieldMask, "CountTypes")
	}
	if in.SubDominant != nil {
		fieldMask = append(fieldMask, "SubDominant")
	}
	if in.Dominant != nil {
		fieldMask = append(fieldMask, "Dominant")
	}
	if in.TrialSite != nil {
		fieldMask = append(fieldMask, "TrialSite")
	}
	return fieldMask
}
