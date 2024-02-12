package service

import (
	"context"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"

	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph"
	log "github.com/sirupsen/logrus"
)

type EcomorphServiceImpl struct {
	repository *repository.Repository
}

func NewEcomorphServiceImpl(repository *repository.Repository) EcomorphServiceImpl {
	return EcomorphServiceImpl{repository}
}

func (e EcomorphServiceImpl) CreateEcomorph(ctx context.Context, in *ecomorph.InputEcomorph) (*ecomorph.Ecomorph, error) {
	createEcomorph, err := e.repository.CreateEcomorph(ctx, ToPB(ctx, in))
	if err != nil {
		log.Error("Insert Ecomorph:", err)
		return nil, err
	}
	log.Debug("Insert Ecomorph: good")

	return createEcomorph, nil
}

func (e EcomorphServiceImpl) GetEcomorphById(ctx context.Context, request *api.IdRequest) (*ecomorph.Ecomorph, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	ecomorph, err := e.repository.GetEcomorphById(ctx, &ecomorph.Ecomorph{Id: request.Id, UserId: userId})
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
	err := e.repository.DeleteEcomorph(ctx, &ecomorph.Ecomorph{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}

func (e EcomorphServiceImpl) StrictUpdateEcomorph(ctx context.Context, in *ecomorph.InputEcomorph) (*ecomorph.Ecomorph, error) {
	createEcomorph, err := e.repository.StrictUpdateEcomorph(ctx, ToPB(ctx, in))
	if err != nil {
		log.Error("Update Ecomorph:", err)
		return nil, err
	}
	log.Debug("Update Ecomorph: good")

	return createEcomorph, nil
}

func (e EcomorphServiceImpl) UpdateEcomorph(ctx context.Context, in *ecomorph.InputEcomorph) (*ecomorph.Ecomorph, error) {
	//TODO implement me
	panic("implement me")
}

func (e EcomorphServiceImpl) GetListEcomorph(ctx context.Context, request *api.EmptyRequest) (*ecomorph.EcomorphsList, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	list, err := e.repository.GetListEcomorph(ctx, &ecomorph.Ecomorph{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &ecomorph.EcomorphsList{Ecomorph: list}, nil
}

func ToPB(ctx context.Context, in *ecomorph.InputEcomorph) *ecomorph.Ecomorph {
	var id *resource.Identifier

	if in.Id != nil {
		id = in.Id
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	userId := middlewares.GetUserIdFromContext(ctx)
	return &ecomorph.Ecomorph{
		Id:          id,
		Title:       in.Payload.Title,
		Description: in.Payload.Description,
		UserId:      userId,
	}
}
