package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/graphql/graph/model"
)

// InsertEcomorph is the resolver for the insertEcomorph field.
func (r *ecomorphMutationResolver) InsertEcomorph(ctx context.Context, obj *model.EcomorphMutation, input *api.InputFormEcomorph) (*api.Ecomorph, error) {
	ecomorph, err := r.service.EcomorphService.CreateEcomorph(ctx, &api.InputEcomorph{Payload: input})
	return ecomorph, err
}

// UpdateEcomorph is the resolver for the updateEcomorph field.
func (r *ecomorphMutationResolver) UpdateEcomorph(ctx context.Context, obj *model.EcomorphMutation, input *api.InputEcomorph) (*api.Ecomorph, error) {
	return r.service.EcomorphService.UpdateEcomorph(ctx, input)
}

// DeleteEcomorphByID is the resolver for the deleteEcomorphById field.
func (r *ecomorphMutationResolver) DeleteEcomorphByID(ctx context.Context, obj *model.EcomorphMutation, id string) (*api.BoolResponse, error) {
	return r.service.EcomorphService.DeleteEcomorph(ctx, ToIdRequest(id))
}

// GetEcomorphByID is the resolver for the getEcomorphById field.
func (r *ecomorphQueryResolver) GetEcomorphByID(ctx context.Context, obj *model.EcomorphQuery, id string) (*api.Ecomorph, error) {
	return r.service.EcomorphService.GetEcomorphById(ctx, ToIdRequest(id))
}

// GetListEcomorph is the resolver for the getListEcomorph field.
func (r *ecomorphQueryResolver) GetListEcomorph(ctx context.Context, obj *model.EcomorphQuery, pages *api.EcomorphListRequest) (*api.EcomorphsList, error) {
	return r.service.EcomorphService.GetListEcomorph(ctx, pages)
}

// EcomorphMutation returns EcomorphMutationResolver implementation.
func (r *Resolver) EcomorphMutation() EcomorphMutationResolver { return &ecomorphMutationResolver{r} }

// EcomorphQuery returns EcomorphQueryResolver implementation.
func (r *Resolver) EcomorphQuery() EcomorphQueryResolver { return &ecomorphQueryResolver{r} }

type ecomorphMutationResolver struct{ *Resolver }
type ecomorphQueryResolver struct{ *Resolver }
