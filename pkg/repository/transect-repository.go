package repository

import (
	"context"
	"fmt"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"gorm.io/gorm"
)

type TransectRepository interface {
	CreateTransect(ctx context.Context, in *api.Transect) (*api.Transect, error)
	GetTransectById(ctx context.Context, in *api.Transect) (*api.Transect, error)
	DeleteTransect(ctx context.Context, in *api.Transect) error
	StrictUpdateTransect(ctx context.Context, in *api.Transect) (*api.Transect, error)
	UpdateTransect(ctx context.Context, in *api.Transect, updateMask *field_mask.FieldMask) (*api.Transect, error)
	GetListTransect(ctx context.Context, in *api.Transect, request *api.PagesRequest) ([]*api.Transect, error)
}

type TransectRepositoryImpl struct {
	db *gorm.DB
}

func NewTransectRepositoryImpl(db *gorm.DB) TransectRepository {
	return TransectRepositoryImpl{db}
}

func (t TransectRepositoryImpl) CreateTransect(ctx context.Context, in *api.Transect) (*api.Transect, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithBeforeCreate_); ok {
		if t.db, err = hook.BeforeCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (t TransectRepositoryImpl) GetTransectById(ctx context.Context, in *api.Transect) (*api.Transect, error) {
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
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithBeforeReadApplyQuery); ok {
		if t.db, err = hook.BeforeReadApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithBeforeReadFind); ok {
		if t.db, err = hook.BeforeReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	ormResponse := api.TransectORM{}
	if err = t.db.Where(&ormObj).Preload("TrialSite").Preload("Dominant").Preload("Dominant").
		Preload("TrialSite.Dominant").Preload("TrialSite.SubDominant").Preload("TrialSite.Plant").
		Preload("TrialSite.Plant.TypePlant").
		Preload("Img").First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(api.TransectORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (t TransectRepositoryImpl) DeleteTransect(ctx context.Context, in *api.Transect) error {
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
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithBeforeDelete_); ok {
		if t.db, err = hook.BeforeDelete_(ctx, t.db); err != nil {
			return err
		}
	}
	err = t.db.Where(&ormObj).Delete(&api.TransectORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, t.db)
	}
	return err
}

func (t TransectRepositoryImpl) StrictUpdateTransect(ctx context.Context, in *api.Transect) (*api.Transect, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTransect")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &api.TransectORM{}
	t.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithBeforeStrictUpdateCleanup); ok {
		if t.db, err = hook.BeforeStrictUpdateCleanup(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithBeforeStrictUpdateSave); ok {
		if t.db, err = hook.BeforeStrictUpdateSave(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

func (t TransectRepositoryImpl) UpdateTransect(ctx context.Context, in *api.Transect, updateMask *field_mask.FieldMask) (*api.Transect, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.Transect
	var err error
	if hook, ok := interface{}(&pbObj).(api.TransectWithBeforePatchRead); ok {
		if t.db, err = hook.BeforePatchRead(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := t.GetTransectById(ctx, &api.Transect{Id: in.GetId()})
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(api.TransectWithBeforePatchApplyFieldMask); ok {
		if t.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	if _, err := api.DefaultApplyFieldMaskTransect(ctx, &pbObj, in, updateMask, "", t.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(api.TransectWithBeforePatchSave); ok {
		if t.db, err = hook.BeforePatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := t.StrictUpdateTransect(ctx, &pbObj)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(api.TransectWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (t TransectRepositoryImpl) GetListTransect(ctx context.Context, in *api.Transect, request *api.PagesRequest) ([]*api.Transect, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithBeforeListApplyQuery); ok {
		if t.db, err = hook.BeforeListApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithBeforeListFind); ok {
		if t.db, err = hook.BeforeListFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if request != nil && request.Page != 0 && request.Limit != 0 {
		offset := (request.Page - 1) * request.Limit
		t.db = t.db.Where(&ormObj).Offset(int(offset)).Limit(int(request.Limit))
	} else {
		t.db = t.db.Where(&ormObj)
	}
	t.db = t.db.Order("id")
	ormResponse := []api.TransectORM{}
	if err := t.db.Preload("Img").Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TransectORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, t.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*api.Transect{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
