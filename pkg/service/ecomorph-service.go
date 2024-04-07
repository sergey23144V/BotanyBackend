package service

import (
	"context"
	er "errors"
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
	GetListEcomorph(ctx context.Context, request *api.EcomorphListRequest) (*api.EcomorphsList, error)
}

type EcomorphServiceImpl struct {
	repository *repository.Repository
}

func NewEcomorphServiceImpl(repository *repository.Repository) EcomorphServiceImpl {
	return EcomorphServiceImpl{repository}
}

func (e EcomorphServiceImpl) CreateEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error) {
	pb, err := e.ToPB(ctx, in)
	if err != nil {
		return nil, err
	}
	createEcomorph, err := e.repository.CreateEcomorph(ctx, pb)
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
	userRole := middlewares.GetUserRoleFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := e.repository.DeleteEcomorph(ctx, &api.Ecomorph{Id: request.Id, UserId: userId}, *userRole)
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (e EcomorphServiceImpl) StrictUpdateEcomorph(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error) {
	pb, err := e.ToPB(ctx, in)
	if err != nil {
		return nil, err
	}
	createEcomorph, err := e.repository.StrictUpdateEcomorph(ctx, pb)
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
	userRole := middlewares.GetUserRoleFromContext(ctx)
	pb, err := e.ToPB(ctx, in)
	if err != nil {
		return nil, err
	}
	return e.repository.UpdateEcomorph(ctx, pb, &field_mask.FieldMask{Paths: fieldMask}, *userRole)
}

func (e EcomorphServiceImpl) GetListEcomorph(ctx context.Context, request *api.EcomorphListRequest) (*api.EcomorphsList, error) {
	var page *api.PagesResponse
	userId := middlewares.GetUserIdFromContext(ctx)

	list, err := e.repository.GetListEcomorph(ctx, &api.Ecomorph{UserId: userId}, request)
	if err != nil {
		return nil, err
	}
	if request.Page != nil {
		page = &api.PagesResponse{Page: request.Page.Page, Limit: request.Page.Limit, Total: int32(len(list))}
	}
	return &api.EcomorphsList{List: list, Page: page}, nil
}

func (e EcomorphServiceImpl) ToPB(ctx context.Context, in *api.InputEcomorph) (*api.Ecomorph, error) {
	var (
		id       *resource.Identifier
		ecomorph *api.Ecomorph
	)

	if in.Id != nil {
		id = in.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}
	userId := middlewares.GetUserIdFromContext(ctx)
	role := middlewares.GetUserRoleFromContext(ctx)

	if *role == api.RoleType_SuperUser && in.Publicly {
		ecomorph = &api.Ecomorph{
			Id:          id,
			Title:       in.Payload.Title,
			Description: in.Payload.Description,
		}
	} else if *role != api.RoleType_SuperUser && in.Publicly {
		return nil, er.New("has no rights")
	} else {
		ecomorph = &api.Ecomorph{
			Id:          id,
			Title:       in.Payload.Title,
			Description: in.Payload.Description,
			UserId:      userId,
		}
	}

	return ecomorph, nil
}
