package service

import (
	"context"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type UserService interface {
	GetMe(ctx context.Context) (*api.UserResponse, error)
	GetUserById(ctx context.Context, request *api.IdRequest) (*api.User, error)
	CheckUserRol(ctx context.Context, request *api.IdRequest) (*api.RoleType, error)
	DeleteUser(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error)
}

type UserServiceImpl struct {
	repository *repository.Repository
}

func NewUserServiceImpl(repository *repository.Repository) UserService {
	return UserServiceImpl{repository: repository}
}

func (u UserServiceImpl) GetMe(ctx context.Context) (*api.UserResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	user, err := u.GetUserById(ctx, &api.IdRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return &api.UserResponse{User: user}, err
}

func (u UserServiceImpl) GetUserById(ctx context.Context, request *api.IdRequest) (*api.User, error) {
	return u.repository.GetUserById(ctx, &api.User{Id: request.Id})
}

func (u UserServiceImpl) CheckUserRol(ctx context.Context, userId *api.IdRequest) (*api.RoleType, error) {
	user, err := u.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &user.Role, nil

}

func (u UserServiceImpl) DeleteUser(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	result := &api.BoolResponse{Result: true}
	err := u.repository.UserRepository.DeleteUser(ctx, &api.User{Id: request.Id})
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}
