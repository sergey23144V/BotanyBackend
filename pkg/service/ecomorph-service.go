package service

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"

	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	log "github.com/sirupsen/logrus"
)

type EcomorphService interface {
	CreateEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error)
	GetEcomorphById(ctx context.Context, request *api.IdRequest) (*api.Ecomorph, error)
	DeleteEcomorph(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
	StrictUpdateEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error)
	UpdateEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error)
	GetListEcomorph(ctx context.Context, request *api.PagesRequest) (*api.EcomorphsList, error)
}

type EcomorphServiceImpl struct {
	repository *repository.Repository
}

func NewEcomorphServiceImpl(repository *repository.Repository) EcomorphServiceImpl {
	return EcomorphServiceImpl{repository}
}

func (e EcomorphServiceImpl) CreateEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error) {
	createEcomorph, err := e.repository.CreateEcomorph(ctx, e.ToPB(ctx, in))
	if err != nil {
		log.Error("Insert Ecomorph:", err)
		return nil, err
	}
	log.Debug("Insert Ecomorph: good")

	return createEcomorph, nil
}

func (e EcomorphServiceImpl) GetEcomorphById(ctx context.Context, request *api.IdRequest) (*api.Ecomorph, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	ecomorph, err := e.repository.GetEcomorphById(ctx, &api.Ecomorph{Id: request.Id, UserId: userId})
	if err != nil {
		log.Error("Get Ecomorph:", err)
		return nil, err
	}
	log.Debug("Get Ecomorph: good")

	return ecomorph, nil
}

func (e EcomorphServiceImpl) DeleteEcomorph(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := e.repository.DeleteEcomorph(ctx, &api.Ecomorph{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (e EcomorphServiceImpl) StrictUpdateEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error) {
	createEcomorph, err := e.repository.StrictUpdateEcomorph(ctx, e.ToPB(ctx, in))
	if err != nil {
		log.Error("Update Ecomorph:", err)
		return nil, err
	}
	log.Debug("Update Ecomorph: good")

	return createEcomorph, nil
}

func (e EcomorphServiceImpl) UpdateEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error) {
	fieldMask := []string{}
	if in.Payload.Title != "" {
		fieldMask = append(fieldMask, "Title")
	}
	if in.Payload.Description != "" {
		fieldMask = append(fieldMask, "Description")
	}
	return e.repository.UpdateEcomorph(ctx, e.ToPB(ctx, in), &field_mask.FieldMask{Paths: fieldMask})
}

func (e EcomorphServiceImpl) GetListEcomorph(ctx context.Context, request *api.PagesRequest) (*api.EcomorphsList, error) {
	var page *api.PagesResponse
	userId := middlewares.GetUserIdFromContext(ctx)
	list, err := e.repository.GetListEcomorph(ctx, &api.Ecomorph{UserId: userId}, request)
	if err != nil {
		return nil, err
	}
	if request != nil {
		page = &api.PagesResponse{Page: request.Page, Limit: request.Limit, Total: int32(len(list))}
	}
	return &api.EcomorphsList{List: list, Page: page}, nil
}

func (e EcomorphServiceImpl) ToPB(ctx context.Context, in *api.InputEcomorph) *api.Ecomorph {
	var id *resource.Identifier

	if in.Id != nil {
		id = in.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &api.Ecomorph{
		Id:          id,
		Title:       in.Payload.Title,
		Description: in.Payload.Description,
		UserId:      userId,
	}
}
