package implementation

import (
	"context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type ImgServetImpl struct {
	service *service.Service
}

func (i ImgServetImpl) GetImgByID(ctx context.Context, request *api.IdRequest) (*api.Img, error) {
	return i.service.GetImgByID(ctx, request)
}

func (i ImgServetImpl) GetListImg(ctx context.Context, request *api.PagesRequest) (*api.ImgList, error) {
	return i.service.GetListImg(ctx, request)
}

func (i ImgServetImpl) MustEmbedUnimplementedImgServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewImgServetImpl(service *service.Service) ImgServetImpl {
	return ImgServetImpl{service}
}
