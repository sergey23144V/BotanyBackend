package repository

import (
	"context"
	"fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/transect"
	"google.golang.org/genproto/protobuf/field_mask"
)

type TransectRepository interface {
	CreateTransect(ctx context.Context, in *transect.Transect) (*transect.Transect, error)
	GetTransectById(ctx context.Context, in *transect.Transect) (*transect.Transect, error)
	DeleteTransect(ctx context.Context, in *transect.Transect) error
	StrictUpdateTransect(ctx context.Context, in *transect.Transect) (*transect.Transect, error)
	UpdateTransect(ctx context.Context, in *transect.Transect, updateMask *field_mask.FieldMask) (*transect.Transect, error)
	GetListTransect(ctx context.Context, in *transect.Transect, request *api.PagesRequest) ([]*transect.Transect, error)
}

type TransectRepositoryImpl struct {
	db *gorm.DB
}

func NewTransectRepositoryImpl(db *gorm.DB) TransectRepository {
	return TransectRepositoryImpl{db}
}

func (t TransectRepositoryImpl) CreateTransect(ctx context.Context, in *transect.Transect) (*transect.Transect, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithBeforeCreate_); ok {
		if t.db, err = hook.BeforeCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (t TransectRepositoryImpl) GetTransectById(ctx context.Context, in *transect.Transect) (*transect.Transect, error) {
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
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithBeforeReadApplyQuery); ok {
		if t.db, err = hook.BeforeReadApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if t.db, err = gorm1.ApplyFieldSelection(ctx, t.db, nil, &transect.TransectORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithBeforeReadFind); ok {
		if t.db, err = hook.BeforeReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	ormResponse := transect.TransectORM{}
	if err = t.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(transect.TransectORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (t TransectRepositoryImpl) DeleteTransect(ctx context.Context, in *transect.Transect) error {
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
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithBeforeDelete_); ok {
		if t.db, err = hook.BeforeDelete_(ctx, t.db); err != nil {
			return err
		}
	}
	err = t.db.Where(&ormObj).Delete(&transect.TransectORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, t.db)
	}
	return err
}

func (t TransectRepositoryImpl) StrictUpdateTransect(ctx context.Context, in *transect.Transect) (*transect.Transect, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTransect")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &transect.TransectORM{}
	t.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithBeforeStrictUpdateCleanup); ok {
		if t.db, err = hook.BeforeStrictUpdateCleanup(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithBeforeStrictUpdateSave); ok {
		if t.db, err = hook.BeforeStrictUpdateSave(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithAfterStrictUpdateSave); ok {
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

func (t TransectRepositoryImpl) UpdateTransect(ctx context.Context, in *transect.Transect, updateMask *field_mask.FieldMask) (*transect.Transect, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj transect.Transect
	var err error
	if hook, ok := interface{}(&pbObj).(transect.TransectWithBeforePatchRead); ok {
		if t.db, err = hook.BeforePatchRead(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := transect.DefaultReadTransect(ctx, &transect.Transect{Id: in.GetId()}, t.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(transect.TransectWithBeforePatchApplyFieldMask); ok {
		if t.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	if _, err := transect.DefaultApplyFieldMaskTransect(ctx, &pbObj, in, updateMask, "", t.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(transect.TransectWithBeforePatchSave); ok {
		if t.db, err = hook.BeforePatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := transect.DefaultStrictUpdateTransect(ctx, &pbObj, t.db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(transect.TransectWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (t TransectRepositoryImpl) GetListTransect(ctx context.Context, in *transect.Transect, request *api.PagesRequest) ([]*transect.Transect, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithBeforeListApplyQuery); ok {
		if t.db, err = hook.BeforeListApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	t.db, err = gorm1.ApplyCollectionOperators(ctx, t.db, &transect.TransectORM{}, &transect.Transect{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithBeforeListFind); ok {
		if t.db, err = hook.BeforeListFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if request != nil && request.Page != 0 && request.Limit != 0 {
		offset := (request.Page - 1) * request.Limit
		t.db = t.db.Where(&ormObj).Offset(offset).Limit(request.Limit)
	} else {
		t.db = t.db.Where(&ormObj)
	}
	t.db = t.db.Order("id")
	ormResponse := []transect.TransectORM{}
	if err := t.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(transect.TransectORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, t.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*transect.Transect{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
