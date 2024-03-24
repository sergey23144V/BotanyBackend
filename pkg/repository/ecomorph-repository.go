package repository

import (
	"context"
	"fmt"
	"github.com/infobloxopen/protoc-gen-gorm/errors"
	"gorm.io/gorm"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
)

type EcomorphRepository interface {
	CreateEcomorph(ctx context.Context, ecomorph *api.Ecomorph) (*api.Ecomorph, error)
	GetEcomorphById(ctx context.Context, in *api.Ecomorph) (*api.Ecomorph, error)
	DeleteEcomorph(ctx context.Context, in *api.Ecomorph) error
	StrictUpdateEcomorph(ctx context.Context, in *api.Ecomorph) (*api.Ecomorph, error)
	UpdateEcomorph(ctx context.Context, in *api.Ecomorph, updateMask *field_mask.FieldMask) (*api.Ecomorph, error)
	GetListEcomorph(ctx context.Context, in *api.Ecomorph, request *api.PagesRequest) ([]*api.Ecomorph, error)
}

type EcomorphRepositoryImpl struct {
	db *gorm.DB
}

func NewEcomorphRepositoryImpl(db *gorm.DB) EcomorphRepositoryImpl {
	return EcomorphRepositoryImpl{db}
}

func (e EcomorphRepositoryImpl) CreateEcomorph(ctx context.Context, in *api.Ecomorph) (*api.Ecomorph, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithBeforeCreate_); ok {
		if e.db, err = hook.BeforeCreate_(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (e EcomorphRepositoryImpl) GetEcomorphById(ctx context.Context, in *api.Ecomorph) (*api.Ecomorph, error) {
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
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithBeforeReadApplyQuery); ok {
		if e.db, err = hook.BeforeReadApplyQuery(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithBeforeReadFind); ok {
		if e.db, err = hook.BeforeReadFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	ormResponse := api.EcomorphORM{}
	if err = e.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(api.EcomorphORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (e EcomorphRepositoryImpl) DeleteEcomorph(ctx context.Context, in *api.Ecomorph) error {
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
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithBeforeDelete_); ok {
		if e.db, err = hook.BeforeDelete_(ctx, e.db); err != nil {
			return err
		}
	}
	//e.db = e.db.Unscoped()
	err = e.db.Where(&ormObj).Delete(&api.EcomorphORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, e.db)
	}
	return err
}

func (e EcomorphRepositoryImpl) StrictUpdateEcomorph(ctx context.Context, in *api.Ecomorph) (*api.Ecomorph, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateEcomorph")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &api.EcomorphORM{}
	e.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithBeforeStrictUpdateCleanup); ok {
		if e.db, err = hook.BeforeStrictUpdateCleanup(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithBeforeStrictUpdateSave); ok {
		if e.db, err = hook.BeforeStrictUpdateSave(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

func (e EcomorphRepositoryImpl) UpdateEcomorph(ctx context.Context, in *api.Ecomorph, updateMask *field_mask.FieldMask) (*api.Ecomorph, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.Ecomorph
	var err error
	if hook, ok := interface{}(&pbObj).(api.EcomorphWithBeforePatchRead); ok {
		if e.db, err = hook.BeforePatchRead(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := api.DefaultReadEcomorph(ctx, &api.Ecomorph{Id: in.GetId()}, e.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(api.EcomorphWithBeforePatchApplyFieldMask); ok {
		if e.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	if _, err := api.DefaultApplyFieldMaskEcomorph(ctx, &pbObj, in, updateMask, "", e.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(api.EcomorphWithBeforePatchSave); ok {
		if e.db, err = hook.BeforePatchSave(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := api.DefaultStrictUpdateEcomorph(ctx, &pbObj, e.db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(api.EcomorphWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (e EcomorphRepositoryImpl) GetListEcomorph(ctx context.Context, in *api.Ecomorph, request *api.PagesRequest) ([]*api.Ecomorph, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithBeforeListApplyQuery); ok {
		if e.db, err = hook.BeforeListApplyQuery(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithBeforeListFind); ok {
		if e.db, err = hook.BeforeListFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if request != nil && request.Page != 0 && request.Limit != 0 {
		offset := (request.Page - 1) * request.Limit
		e.db = e.db.Where(&ormObj).Offset(int(offset)).Limit(int(request.Limit))
	} else {
		e.db = e.db.Where(&ormObj)
	}

	e.db = e.db.Order("id")
	ormResponse := []api.EcomorphORM{}
	if err := e.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, e.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*api.Ecomorph{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
