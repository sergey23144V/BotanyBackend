package repository

import (
	"context"
	"fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	trial_site "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/trial-site"
	"google.golang.org/genproto/protobuf/field_mask"
)

type TrialSiteRepository interface {
	GetListTrialSite(ctx context.Context, in *trial_site.TrialSite) ([]*trial_site.TrialSite, error)
	GetTrialSiteById(ctx context.Context, in *trial_site.TrialSite) (*trial_site.TrialSite, error)
	UpdateTrialSite(ctx context.Context, in *trial_site.TrialSite, updateMask *field_mask.FieldMask) (*trial_site.TrialSite, error)
	StrictUpdateTrialSite(ctx context.Context, in *trial_site.TrialSite) (*trial_site.TrialSite, error)
	DeleteTrialSite(ctx context.Context, in *trial_site.TrialSite) error
	CreateTrialSite(ctx context.Context, in *trial_site.TrialSite) (*trial_site.TrialSite, error)
}

type TrialSiteRepositoryImpl struct {
	db *gorm.DB
}

func NewTrialSiteRepositoryImpl(db *gorm.DB) TrialSiteRepository {
	return TrialSiteRepositoryImpl{db}
}

func (t TrialSiteRepositoryImpl) CreateTrialSite(ctx context.Context, in *trial_site.TrialSite) (*trial_site.TrialSite, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithBeforeCreate_); ok {
		if t.db, err = hook.BeforeCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (t TrialSiteRepositoryImpl) GetTrialSiteById(ctx context.Context, in *trial_site.TrialSite) (*trial_site.TrialSite, error) {
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
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithBeforeReadApplyQuery); ok {
		if t.db, err = hook.BeforeReadApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if t.db, err = gorm1.ApplyFieldSelection(ctx, t.db, nil, &trial_site.TrialSiteORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithBeforeReadFind); ok {
		if t.db, err = hook.BeforeReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	ormResponse := trial_site.TrialSiteORM{}
	if err = t.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(trial_site.TrialSiteORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (t TrialSiteRepositoryImpl) UpdateTrialSite(ctx context.Context, in *trial_site.TrialSite, updateMask *field_mask.FieldMask) (*trial_site.TrialSite, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj trial_site.TrialSite
	var err error
	if hook, ok := interface{}(&pbObj).(trial_site.TrialSiteWithBeforePatchRead); ok {
		if t.db, err = hook.BeforePatchRead(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := trial_site.DefaultReadTrialSite(ctx, &trial_site.TrialSite{Id: in.GetId()}, t.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(trial_site.TrialSiteWithBeforePatchApplyFieldMask); ok {
		if t.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	if _, err := trial_site.DefaultApplyFieldMaskTrialSite(ctx, &pbObj, in, updateMask, "", t.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(trial_site.TrialSiteWithBeforePatchSave); ok {
		if t.db, err = hook.BeforePatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := trial_site.DefaultStrictUpdateTrialSite(ctx, &pbObj, t.db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(trial_site.TrialSiteWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (t TrialSiteRepositoryImpl) StrictUpdateTrialSite(ctx context.Context, in *trial_site.TrialSite) (*trial_site.TrialSite, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTrialSite")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &trial_site.TrialSiteORM{}
	t.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithBeforeStrictUpdateCleanup); ok {
		if t.db, err = hook.BeforeStrictUpdateCleanup(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Model(&ormObj).Association("TypePlant").Replace(ormObj.TypePlant).Error; err != nil {
		return nil, err
	}
	ormObj.TypePlant = nil
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithBeforeStrictUpdateSave); ok {
		if t.db, err = hook.BeforeStrictUpdateSave(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithAfterStrictUpdateSave); ok {
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

func (t TrialSiteRepositoryImpl) DeleteTrialSite(ctx context.Context, in *trial_site.TrialSite) error {
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
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithBeforeDelete_); ok {
		if t.db, err = hook.BeforeDelete_(ctx, t.db); err != nil {
			return err
		}
	}
	err = t.db.Where(&ormObj).Delete(&trial_site.TrialSiteORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, t.db)
	}
	return err
}

func (t TrialSiteRepositoryImpl) GetListTrialSite(ctx context.Context, in *trial_site.TrialSite) ([]*trial_site.TrialSite, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithBeforeListApplyQuery); ok {
		if t.db, err = hook.BeforeListApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	t.db, err = gorm1.ApplyCollectionOperators(ctx, t.db, &trial_site.TrialSiteORM{}, &trial_site.TrialSite{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithBeforeListFind); ok {
		if t.db, err = hook.BeforeListFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	t.db = t.db.Where(&ormObj)
	t.db = t.db.Order("id")
	ormResponse := []trial_site.TrialSiteORM{}
	if err := t.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(trial_site.TrialSiteORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, t.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*trial_site.TrialSite{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
