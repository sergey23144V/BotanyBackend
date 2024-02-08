package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph-entity"
	"github.com/sergey23144V/BotanyBackend/servers/grapgql/graph/model"
)

// InsertEcomorphEntity is the resolver for the insertEcomorphEntity field.
func (r *ecomorphsEntityMutationResolver) InsertEcomorphEntity(ctx context.Context, obj *model.EcomorphsEntityMutation, input *ecomorph_entity.InputFormEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
	panic(fmt.Errorf("not implemented: InsertEcomorphEntity - insertEcomorphEntity"))
}

// UpdateEcomorphEntity is the resolver for the updateEcomorphEntity field.
func (r *ecomorphsEntityMutationResolver) UpdateEcomorphEntity(ctx context.Context, obj *model.EcomorphsEntityMutation, input *ecomorph_entity.InputEcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
	panic(fmt.Errorf("not implemented: UpdateEcomorphEntity - updateEcomorphEntity"))
}

// DeleteEcomorphEntityByID is the resolver for the deleteEcomorphEntityByID field.
func (r *ecomorphsEntityMutationResolver) DeleteEcomorphEntityByID(ctx context.Context, obj *model.EcomorphsEntityMutation, id string) (*api.BoolResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteEcomorphEntityByID - deleteEcomorphEntityByID"))
}

// GetEcomorphEntityByID is the resolver for the getEcomorphEntityByID field.
func (r *ecomorphsEntityQueryResolver) GetEcomorphEntityByID(ctx context.Context, obj *model.EcomorphsEntityQuery, id string) (*ecomorph_entity.EcomorphsEntity, error) {
	panic(fmt.Errorf("not implemented: GetEcomorphEntityByID - getEcomorphEntityByID"))
}

// GetAllEcomorphEntity is the resolver for the getAllEcomorphEntity field.
func (r *ecomorphsEntityQueryResolver) GetAllEcomorphEntity(ctx context.Context, obj *model.EcomorphsEntityQuery) (*ecomorph_entity.EcomorphsEntityList, error) {
	list, err := ecomorph_entity.DefaultListEcomorphsEntity(ctx, r.Db)
	if err != nil {
		return nil, err
	}
	return &ecomorph_entity.EcomorphsEntityList{EcomorphsEntity: list}, nil
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
