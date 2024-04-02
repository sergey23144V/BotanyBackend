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
	"sort"
)

type TrialSiteService interface {
	DetectionPlant(ctx context.Context, request *api.TrialSite) (*api.TrialSite, error)
	CreateTrialSite(ctx context.Context, ecomorph *api.InputTrialSiteRequest) (*api.TrialSite, error)
	GetTrialSiteById(ctx context.Context, request *api.IdRequest) (*api.TrialSite, error)
	DeleteTrialSite(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateTrialSite(ctx context.Context, in *api.InputTrialSiteRequest) (*api.TrialSite, error)
	UpdateTrialSite(ctx context.Context, in *api.InputTrialSiteRequest) (*api.TrialSite, error)
	GetListTrialSite(ctx context.Context, request *api.PagesRequest) (*api.TrialSiteList, error)

	CreatePlant(context.Context, *api.InputPlantRequest) (*api.Plant, error)

	// Получение сущности по id
	GetPlant(context.Context, *api.IdRequest) (*api.Plant, error)
	// Обновление сущности по id
	UpdatePlant(context.Context, *api.InputPlantRequest) (*api.Plant, error)
	// Удаление сущности по заголовку
	DeletePlant(context.Context, *api.IdRequest) (*api.BoolResponse, error)
	// Получение списка всех сущностей
	GetAllPlant(context.Context, *api.PagesRequest) (*api.PlantList, error)
}

type TrialSiteServiceImpl struct {
	repository *repository.Repository
}

func NewTrialSiteServiceImpl(repository *repository.Repository) TrialSiteService {
	return TrialSiteServiceImpl{repository}
}

func (t TrialSiteServiceImpl) CreateTrialSite(ctx context.Context, in *api.InputTrialSiteRequest) (*api.TrialSite, error) {
	pb, err := t.ToPB(ctx, in)
	if err != nil {
		return nil, err
	}
	site, err := t.repository.TrialSiteRepository.CreateTrialSite(ctx, pb)
	if err != nil {
		return nil, err
	}
	return t.DetectionPlant(ctx, site)
}

func (t TrialSiteServiceImpl) DetectionPlant(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error) {
	plantList, err := t.repository.TrialSiteRepository.GetAllPlantORM(ctx, &api.PlantORM{TrialSiteId: &in.Id.ResourceId, UserId: &in.UserId.ResourceId}, nil)
	if err != nil {
		return nil, err
	}
	dtn := in
	if len(plantList) > 0 {
		sort.Slice(plantList, func(i, j int) bool {
			return plantList[i].Coverage < plantList[j].Coverage
		})

		if plantList[0] != nil {
			dtn.Dominant = plantList[0].TypePlant
			if plantList[1] != nil {
				dtn.SubDominant = plantList[1].TypePlant
			}
		}
		dtn.CountTypes = int32(t.repository.CountPlant(plantList))
		return t.repository.TrialSiteRepository.StrictUpdateTrialSite(ctx, dtn)
	}
	return in, err
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
	pb, err := t.ToPB(ctx, in)
	if err != nil {
		return nil, err
	}
	site, err := t.repository.TrialSiteRepository.StrictUpdateTrialSite(ctx, pb)
	if err != nil {
		return nil, err
	}
	return t.DetectionPlant(ctx, site)
}

func (t TrialSiteServiceImpl) UpdateTrialSite(ctx context.Context, in *api.InputTrialSiteRequest) (*api.TrialSite, error) {
	pb, err := t.ToPB(ctx, in)
	if err != nil {
		return nil, err
	}
	site, err := t.repository.TrialSiteRepository.UpdateTrialSite(ctx, pb, &field_mask.FieldMask{Paths: getMask(in)})
	if err != nil {
		return nil, err
	}
	return t.DetectionPlant(ctx, site)
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

func (t TrialSiteServiceImpl) CreatePlant(ctx context.Context, in *api.InputPlantRequest) (*api.Plant, error) {
	return t.repository.TrialSiteRepository.CreatePlant(ctx, t.ToPBPlant(ctx, in))
}
func (t TrialSiteServiceImpl) GetPlant(ctx context.Context, request *api.IdRequest) (*api.Plant, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return t.repository.TrialSiteRepository.GetPlant(ctx, &api.Plant{Id: request.Id, UserId: userId})
}
func (t TrialSiteServiceImpl) UpdatePlant(ctx context.Context, in *api.InputPlantRequest) (*api.Plant, error) {
	return t.repository.TrialSiteRepository.UpdatePlant(ctx, t.ToPBPlant(ctx, in), &field_mask.FieldMask{Paths: getMaskPlan(in)})
}
func (t TrialSiteServiceImpl) DeletePlant(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := t.repository.TrialSiteRepository.DeletePlant(ctx, &api.Plant{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return nil, err
	}
	return result, nil
}
func (t TrialSiteServiceImpl) GetAllPlant(ctx context.Context, request *api.PagesRequest) (*api.PlantList, error) {
	var page *api.PagesResponse

	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := t.repository.TrialSiteRepository.GetAllPlant(ctx, &api.Plant{UserId: userId}, request)
	if request != nil {
		page = &api.PagesResponse{Page: request.Page, Limit: request.Limit, Total: int32(len(getList))}
	}
	if err != nil {
		return nil, err
	}
	result := &api.PlantList{List: getList, Page: page}
	return result, nil
}

func (t TrialSiteServiceImpl) validTrialSite(ctx context.Context, in *api.InputTrialSiteRequest) error {
	var (
		sum       int32
		plantList []*api.Plant
		err       error
	)

	if in.Id != nil {
		plantList, err = t.repository.TrialSiteRepository.GetAllPlantORM(ctx, &api.PlantORM{TrialSiteId: &in.Id.ResourceId}, nil)
		if err != nil {
			return err
		}
	} else {
		inputPlantList := in.Input.Plant

		for _, item := range inputPlantList {
			result, err := t.GetPlant(ctx, &api.IdRequest{Id: item.Id})
			if err != nil {
				return err
			}
			plantList = append(plantList, result)
		}

	}

	for _, item := range plantList {
		sum += item.Coverage
	}
	if sum < 100 {
		return nil
	} else {
		return errors.New("error invalid TrialSite")
	}
}

func (t TrialSiteServiceImpl) ToPB(ctx context.Context, request *api.InputTrialSiteRequest) (*api.TrialSite, error) {
	var id *resource.Identifier

	err := t.validTrialSite(ctx, request)
	if err != nil {
		return nil, err
	}

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
		Plant:       request.Input.Plant,
		Img:         request.Input.Img,
		UserId:      userId,
	}, err
}

func (t TrialSiteServiceImpl) ToPBPlant(ctx context.Context, request *api.InputPlantRequest) *api.Plant {
	var id *resource.Identifier

	if request.Id != nil {
		id = request.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &api.Plant{
		Id:        id,
		TypePlant: request.Input.TypePlantId,
		Count:     request.Input.Count,
		Coverage:  request.Input.Coverage,
		UserId:    userId,
	}
}

func getMaskPlan(in *api.InputPlantRequest) []string {
	fieldMask := []string{}
	if in.Input.Count != 0 {
		fieldMask = append(fieldMask, "Count")
	}
	if in.Input.Coverage != 0 {
		fieldMask = append(fieldMask, "Coverage")
	}
	if in.Input.TypePlantId != nil {
		fieldMask = append(fieldMask, "TypePlantId")
	}

	return fieldMask
}

func getMask(in *api.InputTrialSiteRequest) []string {
	fieldMask := []string{}
	if in.Input.Title != "" {
		fieldMask = append(fieldMask, "Title")
	}
	if in.Input.Covered != 0 {
		fieldMask = append(fieldMask, "Covered")
	}
	if in.Input.Rating != 0 {
		fieldMask = append(fieldMask, "Rating")
	}
	if in.Input.CountTypes != 0 {
		fieldMask = append(fieldMask, "CountTypes")
	}
	if in.Input.SubDominant != nil {
		fieldMask = append(fieldMask, "SubDominant")
	}
	if in.Input.Dominant != nil {
		fieldMask = append(fieldMask, "Dominant")
	}
	if in.Input.Plant != nil {
		fieldMask = append(fieldMask, "Plant")
	}

	return fieldMask
}

//func (t TrialSiteServiceImpl) ToAddPlant(request *api.AddPlantTrialSiteRequest, trialSite *api.TrialSite) *api.TrialSite {
//
//	plants := append(trialSite.TypePlant, &api.TypePlant{Id: request.IdPlant})
//	trialSite.TypePlant = plants
//	return trialSite
//}
