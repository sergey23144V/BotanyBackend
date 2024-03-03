package repository

import (
	"context"
	"fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	"github.com/infobloxopen/protoc-gen-gorm/errors"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph"
	"google.golang.org/genproto/protobuf/field_mask"
)

type EcomorphRepository interface {
	CreateEcomorph(ctx context.Context, ecomorph *ecomorph.Ecomorph) (*ecomorph.Ecomorph, error)
	GetEcomorphById(ctx context.Context, in *ecomorph.Ecomorph) (*ecomorph.Ecomorph, error)
	DeleteEcomorph(ctx context.Context, in *ecomorph.Ecomorph) error
	StrictUpdateEcomorph(ctx context.Context, in *ecomorph.Ecomorph) (*ecomorph.Ecomorph, error)
	UpdateEcomorph(ctx context.Context, in *ecomorph.Ecomorph, updateMask *field_mask.FieldMask) (*ecomorph.Ecomorph, error)
	GetListEcomorph(ctx context.Context, in *ecomorph.Ecomorph, request *api.PagesRequest) ([]*ecomorph.Ecomorph, error)
}

type EcomorphRepositoryImpl struct {
	db *gorm.DB
}

func NewEcomorphRepositoryImpl(db *gorm.DB) EcomorphRepositoryImpl {
	return EcomorphRepositoryImpl{db}
}

func (e EcomorphRepositoryImpl) CreateEcomorph(ctx context.Context, in *ecomorph.Ecomorph) (*ecomorph.Ecomorph, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithBeforeCreate_); ok {
		if e.db, err = hook.BeforeCreate_(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (e EcomorphRepositoryImpl) GetEcomorphById(ctx context.Context, in *ecomorph.Ecomorph) (*ecomorph.Ecomorph, error) {
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
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithBeforeReadApplyQuery); ok {
		if e.db, err = hook.BeforeReadApplyQuery(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if e.db, err = gorm1.ApplyFieldSelection(ctx, e.db, nil, &ecomorph.EcomorphORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithBeforeReadFind); ok {
		if e.db, err = hook.BeforeReadFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	ormResponse := ecomorph.EcomorphORM{}
	if err = e.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ecomorph.EcomorphORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (e EcomorphRepositoryImpl) DeleteEcomorph(ctx context.Context, in *ecomorph.Ecomorph) error {
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
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithBeforeDelete_); ok {
		if e.db, err = hook.BeforeDelete_(ctx, e.db); err != nil {
			return err
		}
	}
	err = e.db.Where(&ormObj).Delete(&ecomorph.EcomorphORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, e.db)
	}
	return err
}

func (e EcomorphRepositoryImpl) StrictUpdateEcomorph(ctx context.Context, in *ecomorph.Ecomorph) (*ecomorph.Ecomorph, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateEcomorph")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ecomorph.EcomorphORM{}
	e.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithBeforeStrictUpdateCleanup); ok {
		if e.db, err = hook.BeforeStrictUpdateCleanup(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithBeforeStrictUpdateSave); ok {
		if e.db, err = hook.BeforeStrictUpdateSave(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithAfterStrictUpdateSave); ok {
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

func (e EcomorphRepositoryImpl) UpdateEcomorph(ctx context.Context, in *ecomorph.Ecomorph, updateMask *field_mask.FieldMask) (*ecomorph.Ecomorph, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj ecomorph.Ecomorph
	var err error
	if hook, ok := interface{}(&pbObj).(ecomorph.EcomorphWithBeforePatchRead); ok {
		if e.db, err = hook.BeforePatchRead(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := ecomorph.DefaultReadEcomorph(ctx, &ecomorph.Ecomorph{Id: in.GetId()}, e.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ecomorph.EcomorphWithBeforePatchApplyFieldMask); ok {
		if e.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	if _, err := ecomorph.DefaultApplyFieldMaskEcomorph(ctx, &pbObj, in, updateMask, "", e.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ecomorph.EcomorphWithBeforePatchSave); ok {
		if e.db, err = hook.BeforePatchSave(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ecomorph.DefaultStrictUpdateEcomorph(ctx, &pbObj, e.db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ecomorph.EcomorphWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (e EcomorphRepositoryImpl) GetListEcomorph(ctx context.Context, in *ecomorph.Ecomorph, request *api.PagesRequest) ([]*ecomorph.Ecomorph, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithBeforeListApplyQuery); ok {
		if e.db, err = hook.BeforeListApplyQuery(ctx, e.db); err != nil {
			return nil, err
		}
	}
	e.db, err = gorm1.ApplyCollectionOperators(ctx, e.db, &ecomorph.EcomorphORM{}, &ecomorph.Ecomorph{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithBeforeListFind); ok {
		if e.db, err = hook.BeforeListFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if request != nil && request.Page != 0 && request.Limit != 0 {
		offset := (request.Page - 1) * request.Limit
		e.db = e.db.Where(&ormObj).Offset(offset).Limit(request.Limit)
	} else {
		e.db = e.db.Where(&ormObj)
	}

	e.db = e.db.Order("id")
	ormResponse := []ecomorph.EcomorphORM{}
	if err := e.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph.EcomorphORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, e.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*ecomorph.Ecomorph{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
