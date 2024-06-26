package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/graphql/graph/model"
)

// CreateTypePlant is the resolver for the createTypePlant field.
func (r *typePlantMutationResolver) CreateTypePlant(ctx context.Context, obj *model.TypePlantMutation, input *api.InputFormTypePlantRequest) (*api.TypePlant, error) {
	return r.service.TypePlantService.CreateTypePlant(ctx, &api.InputTypePlantRequest{Input: input})
}

// UpdateTypePlant is the resolver for the updateTypePlant field.
func (r *typePlantMutationResolver) UpdateTypePlant(ctx context.Context, obj *model.TypePlantMutation, input *api.InputTypePlantRequest) (*api.TypePlant, error) {
	return r.service.TypePlantService.UpdateTypePlant(ctx, input)
}

// AddEcomorphsEntityToTypePlant is the resolver for the addEcomorphsEntityToTypePlant field.
func (r *typePlantMutationResolver) AddEcomorphsEntityToTypePlant(ctx context.Context, obj *model.TypePlantMutation, input *api.InputTypePlant_EcomorphsEntityRequest) (*api.TypePlant, error) {
	return r.service.TypePlantService.AddEcomorphsEntityToTypePlant(ctx, input)
}

// DeleteTypePlant is the resolver for the deleteTypePlant field.
func (r *typePlantMutationResolver) DeleteTypePlant(ctx context.Context, obj *model.TypePlantMutation, id string) (*api.BoolResponse, error) {
	return r.service.TypePlantService.DeleteTypePlant(ctx, ToIdRequest(id))
}

// GetTypePlant is the resolver for the getTypePlant field.
func (r *typePlantQueryResolver) GetTypePlant(ctx context.Context, obj *model.TypePlantQuery, id string) (*api.TypePlant, error) {
	return r.service.TypePlantService.GetTypePlantById(ctx, ToIdRequest(id))
}

// GetAllTypePlant is the resolver for the getAllTypePlant field.
func (r *typePlantQueryResolver) GetAllTypePlant(ctx context.Context, obj *model.TypePlantQuery, pages *api.TypePlantListRequest) (*api.TypePlantList, error) {
	return r.service.TypePlantService.GetListTypePlant(ctx, pages)
}

// TypePlantMutation returns TypePlantMutationResolver implementation.
func (r *Resolver) TypePlantMutation() TypePlantMutationResolver {
	return &typePlantMutationResolver{r}
}

// TypePlantQuery returns TypePlantQueryResolver implementation.
func (r *Resolver) TypePlantQuery() TypePlantQueryResolver { return &typePlantQueryResolver{r} }

type typePlantMutationResolver struct{ *Resolver }
type typePlantQueryResolver struct{ *Resolver }
