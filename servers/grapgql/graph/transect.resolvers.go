package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/transect"
	"github.com/sergey23144V/BotanyBackend/servers/grapgql/graph/model"
)

// CreateTransect is the resolver for the createTransect field.
func (r *transectMutationResolver) CreateTransect(ctx context.Context, obj *model.TransectMutation, input *transect.InputTransectRequest) (*transect.Transect, error) {
	panic(fmt.Errorf("not implemented: CreateTransect - createTransect"))
}

// UpTransect is the resolver for the upTransect field.
func (r *transectMutationResolver) UpTransect(ctx context.Context, obj *model.TransectMutation, input *transect.InputTransectRequest) (*transect.Transect, error) {
	panic(fmt.Errorf("not implemented: UpTransect - upTransect"))
}

// DeleteTransect is the resolver for the deleteTransect field.
func (r *transectMutationResolver) DeleteTransect(ctx context.Context, obj *model.TransectMutation, id string) (*api.BoolResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteTransect - deleteTransect"))
}

// GetTransect is the resolver for the getTransect field.
func (r *transectQueryResolver) GetTransect(ctx context.Context, obj *model.TransectQuery, id string) (*transect.Transect, error) {
	panic(fmt.Errorf("not implemented: GetTransect - getTransect"))
}

// GetAllTransect is the resolver for the getAllTransect field.
func (r *transectQueryResolver) GetAllTransect(ctx context.Context, obj *model.TransectQuery) (*transect.TransectList, error) {
	list, err := transect.DefaultListTransect(ctx, r.Db)
	if err != nil {
		return nil, err
	}
	return &transect.TransectList{Transect: list}, nil
}

// TransectMutation returns TransectMutationResolver implementation.
func (r *Resolver) TransectMutation() TransectMutationResolver { return &transectMutationResolver{r} }

// TransectQuery returns TransectQueryResolver implementation.
func (r *Resolver) TransectQuery() TransectQueryResolver { return &transectQueryResolver{r} }

type transectMutationResolver struct{ *Resolver }
type transectQueryResolver struct{ *Resolver }
