package implementation

import (
	"context"
	"github.com/sergey23144V/BotanyBackend/pkg/service"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type UserServetImpl struct {
	service *service.Service
}

func NewUserServetImpl(service *service.Service) api.UserServiceServer {
	return UserServetImpl{service}
}

func (u UserServetImpl) GetMe(ctx context.Context, request *api.EmptyRequest) (*api.UserResponse, error) {
	return u.service.GetMe(ctx)
}

func (u UserServetImpl) MustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}
