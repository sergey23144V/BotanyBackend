package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/graphql/graph/model"
)

// InsertEcomorphEntity is the resolver for the insertEcomorphEntity field.
func (r *ecomorphsEntityMutationResolver) InsertEcomorphEntity(ctx context.Context, obj *model.EcomorphsEntityMutation, input *api.InputFormEcomorphsEntity) (*api.EcomorphsEntity, error) {
	return r.service.EcomorphsEntityService.CreateEcomorphsEntity(ctx, &api.InputEcomorphsEntity{Input: input})
}

// UpdateEcomorphEntity is the resolver for the updateEcomorphEntity field.
func (r *ecomorphsEntityMutationResolver) UpdateEcomorphEntity(ctx context.Context, obj *model.EcomorphsEntityMutation, input *api.InputEcomorphsEntity) (*api.EcomorphsEntity, error) {
	return r.service.EcomorphsEntityService.StrictUpdateEcomorphsEntity(ctx, input)
}

// DeleteEcomorphEntityByID is the resolver for the deleteEcomorphEntityByID field.
func (r *ecomorphsEntityMutationResolver) DeleteEcomorphEntityByID(ctx context.Context, obj *model.EcomorphsEntityMutation, id string) (*api.BoolResponse, error) {
	return r.service.EcomorphsEntityService.DeleteEcomorphsEntity(ctx, ToIdRequest(id))
}

// GetEcomorphEntityByID is the resolver for the getEcomorphEntityByID field.
func (r *ecomorphsEntityQueryResolver) GetEcomorphEntityByID(ctx context.Context, obj *model.EcomorphsEntityQuery, id string) (*api.EcomorphsEntity, error) {
	return r.service.EcomorphsEntityService.GetEcomorphsEntityById(ctx, ToIdRequest(id))
}

// GetAllEcomorphEntity is the resolver for the getAllEcomorphEntity field.
func (r *ecomorphsEntityQueryResolver) GetAllEcomorphEntity(ctx context.Context, obj *model.EcomorphsEntityQuery, pages *api.EcomorphsEntityListRequest) (*api.EcomorphsEntityList, error) {
	return r.service.EcomorphsEntityService.GetListEcomorphsEntity(ctx, pages)
}

// EcomorphsEntityMutation returns EcomorphsEntityMutationResolver implementation.
func (r *Resolver) EcomorphsEntityMutation() EcomorphsEntityMutationResolver {
	return &ecomorphsEntityMutationResolver{r}
}

// EcomorphsEntityQuery returns EcomorphsEntityQueryResolver implementation.
func (r *Resolver) EcomorphsEntityQuery() EcomorphsEntityQueryResolver {
	return &ecomorphsEntityQueryResolver{r}
}

type ecomorphsEntityMutationResolver struct{ *Resolver }
type ecomorphsEntityQueryResolver struct{ *Resolver }
