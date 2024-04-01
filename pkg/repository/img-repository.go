package repository

import (
	"context"
	"github.com/infobloxopen/protoc-gen-gorm/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"gorm.io/gorm"
)

type ImgRepository interface {
	CreateImg(ctx context.Context, in *api.Img) (*api.Img, error)
	GetImgByID(ctx context.Context, in *api.Img) (*api.Img, error)
	GetListImg(ctx context.Context, ormObj *api.ImgORM, request *api.PagesRequest) ([]*api.Img, error)
	UpdateImg(ctx context.Context, in *api.Img, updateMask *field_mask.FieldMask) (*api.Img, error)
	DeleteImg(ctx context.Context, in *api.Img) error
}

type ImgRepositoryImpl struct {
	db *gorm.DB
}

func NewImgRepositoryImpl(db *gorm.DB) ImgRepositoryImpl {
	return ImgRepositoryImpl{db}
}

func (i ImgRepositoryImpl) CreateImg(ctx context.Context, in *api.Img) (*api.Img, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.ImgORMWithBeforeCreate_); ok {
		if i.db, err = hook.BeforeCreate_(ctx, i.db); err != nil {
			return nil, err
		}
	}
	if err = i.db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.ImgORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, i.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (i ImgRepositoryImpl) GetImgByID(ctx context.Context, in *api.Img) (*api.Img, error) {
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
	if hook, ok := interface{}(&ormObj).(api.ImgORMWithBeforeReadApplyQuery); ok {
		if i.db, err = hook.BeforeReadApplyQuery(ctx, i.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.ImgORMWithBeforeReadFind); ok {
		if i.db, err = hook.BeforeReadFind(ctx, i.db); err != nil {
			return nil, err
		}
	}
	ormResponse := api.ImgORM{}
	if err = i.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(api.ImgORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, i.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (i ImgRepositoryImpl) GetListImg(ctx context.Context, ormObj *api.ImgORM, request *api.PagesRequest) ([]*api.Img, error) {
	var err error
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.ImgORMWithBeforeListApplyQuery); ok {
		if i.db, err = hook.BeforeListApplyQuery(ctx, i.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.ImgORMWithBeforeListFind); ok {
		if i.db, err = hook.BeforeListFind(ctx, i.db); err != nil {
			return nil, err
		}
	}
	if request != nil && request.Page != 0 && request.Limit != 0 {
		offset := (request.Page - 1) * request.Limit
		i.db = i.db.Where(&ormObj).Offset(int(offset)).Limit(int(request.Limit))
	} else {
		i.db = i.db.Where(&ormObj)
	}
	i.db = i.db.Order("id")
	ormResponse := []api.ImgORM{}
	if err := i.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.ImgORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, i.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*api.Img{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

func (i ImgRepositoryImpl) UpdateImg(ctx context.Context, in *api.Img, updateMask *field_mask.FieldMask) (*api.Img, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.Img
	var err error
	if hook, ok := interface{}(&pbObj).(api.ImgWithBeforePatchRead); ok {
		if i.db, err = hook.BeforePatchRead(ctx, in, updateMask, i.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := api.DefaultReadImg(ctx, &api.Img{Id: in.GetId()}, i.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(api.ImgWithBeforePatchApplyFieldMask); ok {
		if i.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, i.db); err != nil {
			return nil, err
		}
	}
	if _, err := api.DefaultApplyFieldMaskImg(ctx, &pbObj, in, updateMask, "", i.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(api.ImgWithBeforePatchSave); ok {
		if i.db, err = hook.BeforePatchSave(ctx, in, updateMask, i.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := api.DefaultStrictUpdateImg(ctx, &pbObj, i.db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(api.ImgWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, i.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (i ImgRepositoryImpl) DeleteImg(ctx context.Context, in *api.Img) error {
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
	if hook, ok := interface{}(&ormObj).(api.ImgORMWithBeforeDelete_); ok {
		if i.db, err = hook.BeforeDelete_(ctx, i.db); err != nil {
			return err
		}
	}
	err = i.db.Where(&ormObj).Delete(&api.ImgORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(api.ImgORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, i.db)
	}
	return err
}
