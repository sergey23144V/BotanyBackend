package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/graphql/graph/model"
)

// CreateTransect is the resolver for the createTransect field.
func (r *transectMutationResolver) CreateTransect(ctx context.Context, obj *model.TransectMutation, input *api.InputFormTransectRequest) (*api.Transect, error) {
	return r.service.CreateTransect(ctx, &api.InputTransectRequest{Input: input})
}

// UpTransect is the resolver for the upTransect field.
func (r *transectMutationResolver) UpTransect(ctx context.Context, obj *model.TransectMutation, input *api.InputTransectRequest) (*api.Transect, error) {
	return r.service.UpdateTransect(ctx, input)
}

// AddTrialSiteToTransect is the resolver for the addTrialSiteToTransect field.
func (r *transectMutationResolver) AddTrialSiteToTransect(ctx context.Context, obj *model.TransectMutation, input *api.InputTransectRequest) (*api.Transect, error) {
	return r.service.AddTrialSiteToTransect(ctx, input)
}

// DeleteTransect is the resolver for the deleteTransect field.
func (r *transectMutationResolver) DeleteTransect(ctx context.Context, obj *model.TransectMutation, id string) (*api.BoolResponse, error) {
	return r.service.DeleteTransect(ctx, ToIdRequest(id))
}

// GetTransect is the resolver for the getTransect field.
func (r *transectQueryResolver) GetTransect(ctx context.Context, obj *model.TransectQuery, id string) (*api.Transect, error) {
	return r.service.GetTransectById(ctx, ToIdRequest(id))
}

// GetAllTransect is the resolver for the getAllTransect field.
func (r *transectQueryResolver) GetAllTransect(ctx context.Context, obj *model.TransectQuery, pages *api.TransectListRequest) (*api.TransectList, error) {
	return r.service.GetListTransect(ctx, pages)
}

// TransectMutation returns TransectMutationResolver implementation.
func (r *Resolver) TransectMutation() TransectMutationResolver { return &transectMutationResolver{r} }

// TransectQuery returns TransectQueryResolver implementation.
func (r *Resolver) TransectQuery() TransectQueryResolver { return &transectQueryResolver{r} }

type transectMutationResolver struct{ *Resolver }
type transectQueryResolver struct{ *Resolver }
