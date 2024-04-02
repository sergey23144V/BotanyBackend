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
	DetectionPlant(ctx context.Context, in *api.Transect) (*api.Transect, error)
}

type TransectServiceImpl struct {
	repository *repository.Repository
}

func NewTransectServiceImpl(repository *repository.Repository) TransectService {
	return TransectServiceImpl{repository}
}

func (t TransectServiceImpl) CreateTransect(ctx context.Context, ecomorph *api.InputTransectRequest) (*api.Transect, error) {
	transect, err := t.repository.TransectRepository.CreateTransect(ctx, t.ToPB(ctx, ecomorph))
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
		dominants    *api.TypePlant
		subDominants *api.TypePlant
	)
	output, err := t.GetTransectById(ctx, &api.IdRequest{Id: in.Id})
	if err != nil {
		return in, err
	}
	trialSites := output.TrialSite
	countsDominants := make(map[string]int)
	countsSubDominants := make(map[string]int)

	for _, item := range trialSites {

		countsDominants[item.Dominant.Id.ResourceId]++
		countsSubDominants[item.SubDominant.Id.ResourceId]++
	}

	var maxCount int
	for item, count := range countsDominants {
		if count > maxCount {
			typePlant, err := t.repository.TypePlantRepository.GetTypePlantById(ctx, &api.TypePlant{Id: &resource.Identifier{ResourceId: item}})
			if err != nil {
				return nil, err
			}
			dominants = typePlant
			maxCount = count
		}
	}
	maxCount = 0
	for item, count := range countsSubDominants {
		if count > maxCount {
			typePlant, err := t.repository.TypePlantRepository.GetTypePlantById(ctx, &api.TypePlant{Id: &resource.Identifier{ResourceId: item}})
			if err != nil {
				return nil, err
			}
			subDominants = typePlant
			maxCount = count
		}
	}

	output.Dominant = dominants
	output.SubDominant = subDominants

	return t.repository.StrictUpdateTransect(ctx, output)
}

func (t TransectServiceImpl) StrictUpdateTransect(ctx context.Context, in *api.InputTransectRequest) (*api.Transect, error) {
	return t.repository.TransectRepository.StrictUpdateTransect(ctx, t.ToPB(ctx, in))
}

func (t TransectServiceImpl) UpdateTransect(ctx context.Context, in *api.InputTransectRequest) (*api.Transect, error) {
	return t.repository.TransectRepository.UpdateTransect(ctx, t.ToPB(ctx, in), &field_mask.FieldMask{Paths: t.getMask(in.Input)})
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
