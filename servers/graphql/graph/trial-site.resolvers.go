package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"

	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/graphql/graph/model"
)

// CreateTrialSite is the resolver for the createTrialSite field.
func (r *trialSiteMutationResolver) CreateTrialSite(ctx context.Context, obj *model.TrialSiteMutation, input *api.InputFormTrialSiteRequest) (*api.TrialSite, error) {
	if !middlewares.ValidToken(ctx) {
		return nil, errors.NotAuthorization
	}
	return r.service.CreateTrialSite(ctx, &api.InputTrialSiteRequest{Input: input})
}

// UpTrialSite is the resolver for the upTrialSite field.
func (r *trialSiteMutationResolver) UpTrialSite(ctx context.Context, obj *model.TrialSiteMutation, input *api.InputTrialSiteRequest) (*api.TrialSite, error) {
	if !middlewares.ValidToken(ctx) {
		return nil, errors.NotAuthorization
	}
	return r.service.UpdateTrialSite(ctx, input)
}

// DeleteTrialSite is the resolver for the deleteTrialSite field.
func (r *trialSiteMutationResolver) DeleteTrialSite(ctx context.Context, obj *model.TrialSiteMutation, id string) (*api.BoolResponse, error) {
	if !middlewares.ValidToken(ctx) {
		return nil, errors.NotAuthorization
	}
	return r.service.DeleteTrialSite(ctx, ToIdRequest(id))
}

// CreatePlant is the resolver for the createPlant field.
func (r *trialSiteMutationResolver) CreatePlant(ctx context.Context, obj *model.TrialSiteMutation, input *api.InputPlantRequest) (*api.Plant, error) {
	if !middlewares.ValidToken(ctx) {
		return nil, errors.NotAuthorization
	}
	return r.service.CreatePlant(ctx, input)
}

// UpdatePlant is the resolver for the updatePlant field.
func (r *trialSiteMutationResolver) UpdatePlant(ctx context.Context, obj *model.TrialSiteMutation, input *api.InputPlantRequest) (*api.Plant, error) {
	if !middlewares.ValidToken(ctx) {
		return nil, errors.NotAuthorization
	}
	return r.service.UpdatePlant(ctx, input)
}

// DeletePlant is the resolver for the deletePlant field.
func (r *trialSiteMutationResolver) DeletePlant(ctx context.Context, obj *model.TrialSiteMutation, id string) (*api.BoolResponse, error) {
	if !middlewares.ValidToken(ctx) {
		return nil, errors.NotAuthorization
	}
	return r.service.DeletePlant(ctx, ToIdRequest(id))
}

// GetTrialSite is the resolver for the getTrialSite field.
func (r *trialSiteQueryResolver) GetTrialSite(ctx context.Context, obj *model.TrialSiteQuery, id string) (*api.TrialSite, error) {
	if !middlewares.ValidToken(ctx) {
		return nil, errors.NotAuthorization
	}
	return r.service.GetTrialSiteById(ctx, ToIdRequest(id))
}

// GetAllTrialSite is the resolver for the getAllTrialSite field.
func (r *trialSiteQueryResolver) GetAllTrialSite(ctx context.Context, obj *model.TrialSiteQuery, pages *api.PagesRequest) (*api.TrialSiteList, error) {
	if !middlewares.ValidToken(ctx) {
		return nil, errors.NotAuthorization
	}
	return r.service.GetListTrialSite(ctx, pages)
}

// GetPlant is the resolver for the getPlant field.
func (r *trialSiteQueryResolver) GetPlant(ctx context.Context, obj *model.TrialSiteQuery, id string) (*api.Plant, error) {
	if !middlewares.ValidToken(ctx) {
		return nil, errors.NotAuthorization
	}
	return r.service.GetPlant(ctx, ToIdRequest(id))
}

// GetAllPlant is the resolver for the getAllPlant field.
func (r *trialSiteQueryResolver) GetAllPlant(ctx context.Context, obj *model.TrialSiteQuery, pages *api.PagesRequest) (*api.PlantList, error) {
	return r.service.GetAllPlant(ctx, pages)
}

// TrialSiteMutation returns TrialSiteMutationResolver implementation.
func (r *Resolver) TrialSiteMutation() TrialSiteMutationResolver {
	return &trialSiteMutationResolver{r}
}

// TrialSiteQuery returns TrialSiteQueryResolver implementation.
func (r *Resolver) TrialSiteQuery() TrialSiteQueryResolver { return &trialSiteQueryResolver{r} }

type trialSiteMutationResolver struct{ *Resolver }
type trialSiteQueryResolver struct{ *Resolver }
