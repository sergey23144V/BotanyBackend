package implementation

import (
	"context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type AnalysisServetImpl struct {
	service *service.Service
}

func NewAnalysisServetImplServetImpl(repository *service.Service) api.AnalysisServiceServer {
	return AnalysisServetImpl{repository}
}

func (a AnalysisServetImpl) CreatAnalysis(ctx context.Context, analysis *api.InputCreateAnalysis) (*api.Analysis, error) {
	return a.service.AnalysisService.CreatAnalysis(ctx, analysis)
}

func (a AnalysisServetImpl) RepeatedAnalysis(ctx context.Context, analysis *api.InputUpdateAnalysis) (*api.Analysis, error) {
	//TODO implement me
	panic("implement me")
}

func (a AnalysisServetImpl) GetAnalysis(ctx context.Context, request *api.IdRequest) (*api.Analysis, error) {
	//TODO implement me
	panic("implement me")
}

func (a AnalysisServetImpl) GetListAnalysis(ctx context.Context, request *api.PagesRequest) (*api.AnalysisList, error) {
	//TODO implement me
	panic("implement me")
}

func (a AnalysisServetImpl) DeleteAnalysis(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AnalysisServetImpl) MustEmbedUnimplementedAnalysisServiceServer() {
	//TODO implement me
	panic("implement me")
}
