package repository

import (
	"context"
	"github.com/infobloxopen/protoc-gen-gorm/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, ecomorph *api.User) (*api.User, error)
	GetUserById(ctx context.Context, in *api.User) (*api.User, error)
	DeleteUser(ctx context.Context, in *api.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepositoryImpl {
	return UserRepositoryImpl{db}
}

func (e UserRepositoryImpl) CreateUser(ctx context.Context, in *api.User) (*api.User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = e.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}

	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (e UserRepositoryImpl) GetUserById(ctx context.Context, in *api.User) (*api.User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors.EmptyIdError
	}
	ormResponse := api.UserORM{}
	if err = e.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (e UserRepositoryImpl) GetListUser(ctx context.Context) ([]*api.User, error) {
	ormResponse := []*api.UserORM{}
	result := e.db.Find(&ormResponse)
	if result.Error != nil {
		return nil, result.Error
	}
	pbResponse := []*api.User{}
	for _, orm := range ormResponse {
		pb, err := orm.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &pb)
	}

	return pbResponse, nil
}

func (e UserRepositoryImpl) DeleteUser(ctx context.Context, in *api.User) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(api.UserORMWithBeforeDelete_); ok {
		if e.db, err = hook.BeforeDelete_(ctx, e.db); err != nil {
			return err
		}
	}
	//e.db = e.db.Unscoped()
	err = e.db.Where(&ormObj).Delete(&api.UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(api.UserORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, e.db)
	}
	return err
}
