package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/graphql/graph/model"
)

// TypeAnalysis is the resolver for the typeAnalysis field.
func (r *analysisResolver) TypeAnalysis(ctx context.Context, obj *api.Analysis) (model.TypeAnalysis, error) {
	if obj.TypeAnalysis == api.TypeAnalysis_TypeAnalysisTransect {
		return model.TypeAnalysisTypeAnalysisTransect, nil
	} else if obj.TypeAnalysis == api.TypeAnalysis_TypeAnalysisPlant {
		return model.TypeAnalysisTypeAnalysisPlant, nil
	}
	return "", nil
}

// CreatAnalysis is the resolver for the creatAnalysis field.
func (r *analysisMutationResolver) CreatAnalysis(ctx context.Context, obj *model.AnalysisMutation, input *api.InputCreateAnalysis) (*api.Analysis, error) {
	return r.service.AnalysisService.CreatAnalysis(ctx, input)
}

// RepeatedAnalysis is the resolver for the repeatedAnalysis field.
func (r *analysisMutationResolver) RepeatedAnalysis(ctx context.Context, obj *model.AnalysisMutation, input *api.InputUpdateAnalysis) (*api.Analysis, error) {
	panic(fmt.Errorf("not implemented: RepeatedAnalysis - repeatedAnalysis"))
}

// DeleteAnalysis is the resolver for the deleteAnalysis field.
func (r *analysisMutationResolver) DeleteAnalysis(ctx context.Context, obj *model.AnalysisMutation, id string) (*api.BoolResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteAnalysis - deleteAnalysis"))
}

// GetAnalysis is the resolver for the getAnalysis field.
func (r *analysisQueryResolver) GetAnalysis(ctx context.Context, obj *model.AnalysisQuery, id string) (*api.Analysis, error) {
	panic(fmt.Errorf("not implemented: GetAnalysis - getAnalysis"))
}

// GetListAnalysis is the resolver for the getListAnalysis field.
func (r *analysisQueryResolver) GetListAnalysis(ctx context.Context, obj *model.AnalysisQuery, pages *api.PagesRequest) (*api.AnalysisList, error) {
	panic(fmt.Errorf("not implemented: GetListAnalysis - getListAnalysis"))
}

// TypeAnalysis is the resolver for the typeAnalysis field.
func (r *analysisInputResolver) TypeAnalysis(ctx context.Context, obj *api.Analysis, data model.TypeAnalysis) error {
	if data == model.TypeAnalysisTypeAnalysisTransect {
		obj.TypeAnalysis = api.TypeAnalysis_TypeAnalysisTransect
		return nil
	} else if data == model.TypeAnalysisTypeAnalysisPlant {
		obj.TypeAnalysis = api.TypeAnalysis_TypeAnalysisPlant
		return nil
	}
	return nil
}

// TypeAnalysis is the resolver for the typeAnalysis field.
func (r *inputCreateAnalysisResolver) TypeAnalysis(ctx context.Context, obj *api.InputCreateAnalysis, data model.TypeAnalysis) error {
	if data == model.TypeAnalysisTypeAnalysisTransect {
		obj.TypeAnalysis = api.TypeAnalysis_TypeAnalysisTransect
		return nil
	} else if data == model.TypeAnalysisTypeAnalysisPlant {
		obj.TypeAnalysis = api.TypeAnalysis_TypeAnalysisPlant
		return nil
	}
	return nil
}

// Analysis returns AnalysisResolver implementation.
func (r *Resolver) Analysis() AnalysisResolver { return &analysisResolver{r} }

// AnalysisMutation returns AnalysisMutationResolver implementation.
func (r *Resolver) AnalysisMutation() AnalysisMutationResolver { return &analysisMutationResolver{r} }

// AnalysisQuery returns AnalysisQueryResolver implementation.
func (r *Resolver) AnalysisQuery() AnalysisQueryResolver { return &analysisQueryResolver{r} }

// AnalysisInput returns AnalysisInputResolver implementation.
func (r *Resolver) AnalysisInput() AnalysisInputResolver { return &analysisInputResolver{r} }

// InputCreateAnalysis returns InputCreateAnalysisResolver implementation.
func (r *Resolver) InputCreateAnalysis() InputCreateAnalysisResolver {
	return &inputCreateAnalysisResolver{r}
}

type analysisResolver struct{ *Resolver }
type analysisMutationResolver struct{ *Resolver }
type analysisQueryResolver struct{ *Resolver }
type analysisInputResolver struct{ *Resolver }
type inputCreateAnalysisResolver struct{ *Resolver }
