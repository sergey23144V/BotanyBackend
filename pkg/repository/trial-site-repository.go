package repository

import (
	"context"
	"fmt"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"gorm.io/gorm"
)

type TrialSiteRepository interface {
	GetListTrialSite(ctx context.Context, in *api.TrialSite, request *api.PagesRequest) ([]*api.TrialSite, error)
	GetTrialSiteById(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error)
	UpdateTrialSite(ctx context.Context, in *api.TrialSite, updateMask *field_mask.FieldMask) (*api.TrialSite, error)
	StrictUpdateTrialSite(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error)
	DeleteTrialSite(ctx context.Context, in *api.TrialSite) error
	CreateTrialSite(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error)
}

type TrialSiteRepositoryImpl struct {
	db *gorm.DB
}

func NewTrialSiteRepositoryImpl(db *gorm.DB) TrialSiteRepository {
	return TrialSiteRepositoryImpl{db}
}

func (t TrialSiteRepositoryImpl) CreateTrialSite(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithBeforeCreate_); ok {
		if t.db, err = hook.BeforeCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (t TrialSiteRepositoryImpl) GetTrialSiteById(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error) {
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
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithBeforeReadApplyQuery); ok {
		if t.db, err = hook.BeforeReadApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithBeforeReadFind); ok {
		if t.db, err = hook.BeforeReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	ormResponse := api.TrialSiteORM{}
	if err = t.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(api.TrialSiteORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (t TrialSiteRepositoryImpl) UpdateTrialSite(ctx context.Context, in *api.TrialSite, updateMask *field_mask.FieldMask) (*api.TrialSite, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.TrialSite
	var err error
	if hook, ok := interface{}(&pbObj).(api.TrialSiteWithBeforePatchRead); ok {
		if t.db, err = hook.BeforePatchRead(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := api.DefaultReadTrialSite(ctx, &api.TrialSite{Id: in.GetId()}, t.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(api.TrialSiteWithBeforePatchApplyFieldMask); ok {
		if t.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	if _, err := api.DefaultApplyFieldMaskTrialSite(ctx, &pbObj, in, updateMask, "", t.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(api.TrialSiteWithBeforePatchSave); ok {
		if t.db, err = hook.BeforePatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := api.DefaultStrictUpdateTrialSite(ctx, &pbObj, t.db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(api.TrialSiteWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (t TrialSiteRepositoryImpl) StrictUpdateTrialSite(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTrialSite")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &api.TrialSiteORM{}
	t.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithBeforeStrictUpdateCleanup); ok {
		if t.db, err = hook.BeforeStrictUpdateCleanup(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithBeforeStrictUpdateSave); ok {
		if t.db, err = hook.BeforeStrictUpdateSave(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Omit("EcomorphsEntity").Preload("Dominant").Preload("SubDominant").Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithAfterStrictUpdateSave); ok {
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

func (t TrialSiteRepositoryImpl) DeleteTrialSite(ctx context.Context, in *api.TrialSite) error {
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
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithBeforeDelete_); ok {
		if t.db, err = hook.BeforeDelete_(ctx, t.db); err != nil {
			return err
		}
	}
	err = t.db.Where(&ormObj).Delete(&api.TrialSiteORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, t.db)
	}
	return err
}

func (t TrialSiteRepositoryImpl) GetListTrialSite(ctx context.Context, in *api.TrialSite, request *api.PagesRequest) ([]*api.TrialSite, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithBeforeListApplyQuery); ok {
		if t.db, err = hook.BeforeListApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithBeforeListFind); ok {
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
	ormResponse := []api.TrialSiteORM{}
	if err := t.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TrialSiteORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, t.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*api.TrialSite{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
