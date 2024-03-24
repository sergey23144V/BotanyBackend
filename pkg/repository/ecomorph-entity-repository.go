package repository

import (
	context "context"
	"fmt"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"gorm.io/gorm"
)

type EcomorphsEntityRepository interface {
	CreateEcomorphsEntity(ctx context.Context, ecomorph *api.EcomorphsEntity) (*api.EcomorphsEntity, error)
	GetEcomorphsEntityById(ctx context.Context, in *api.EcomorphsEntity) (*api.EcomorphsEntity, error)
	DeleteEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity) error
	StrictUpdateEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity) (*api.EcomorphsEntity, error)
	UpdateEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity, updateMask *field_mask.FieldMask) (*api.EcomorphsEntity, error)
	GetListEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity, request *api.PagesRequest) ([]*api.EcomorphsEntity, error)
}

type EcomorphsEntityRepositoryImpl struct {
	db *gorm.DB
}

func NewEcomorphsEntityRepositoryImpl(db *gorm.DB) EcomorphsEntityRepository {
	return EcomorphsEntityRepositoryImpl{db}
}

func (e EcomorphsEntityRepositoryImpl) CreateEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity) (*api.EcomorphsEntity, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithBeforeCreate_); ok {
		if e.db, err = hook.BeforeCreate_(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Create(&ormObj).Preload("Ecomorphs").Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (e EcomorphsEntityRepositoryImpl) GetEcomorphsEntityById(ctx context.Context, in *api.EcomorphsEntity) (*api.EcomorphsEntity, error) {
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
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithBeforeReadApplyQuery); ok {
		if e.db, err = hook.BeforeReadApplyQuery(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithBeforeReadFind); ok {
		if e.db, err = hook.BeforeReadFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	ormResponse := api.EcomorphsEntityORM{}
	if err = e.db.Where(&ormObj).Preload("Ecomorphs").First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(api.EcomorphsEntityORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (e EcomorphsEntityRepositoryImpl) DeleteEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity) error {
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
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithBeforeDelete_); ok {
		if e.db, err = hook.BeforeDelete_(ctx, e.db); err != nil {
			return err
		}
	}
	tx := e.db.Delete(&ormObj)
	if tx.Error != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, e.db)
	}
	if tx.RowsAffected == 0 {
		return errors.NotDelete
	}
	return err
}

func (e EcomorphsEntityRepositoryImpl) StrictUpdateEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity) (*api.EcomorphsEntity, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateEcomorphsEntity")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &api.EcomorphsEntityORM{}
	e.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithBeforeStrictUpdateCleanup); ok {
		if e.db, err = hook.BeforeStrictUpdateCleanup(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Model(&ormObj).Association("Ecomorphs").Replace(ormObj.Ecomorphs); err != nil {
		return nil, err
	}
	ormObj.Ecomorphs = nil
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithBeforeStrictUpdateSave); ok {
		if e.db, err = hook.BeforeStrictUpdateSave(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithAfterStrictUpdateSave); ok {
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

func (e EcomorphsEntityRepositoryImpl) UpdateEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity, updateMask *field_mask.FieldMask) (*api.EcomorphsEntity, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.EcomorphsEntity
	var err error
	if hook, ok := interface{}(&pbObj).(api.EcomorphsEntityWithBeforePatchRead); ok {
		if e.db, err = hook.BeforePatchRead(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := api.DefaultReadEcomorphsEntity(ctx, &api.EcomorphsEntity{Id: in.GetId()}, e.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(api.EcomorphsEntityWithBeforePatchApplyFieldMask); ok {
		if e.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	if _, err := api.DefaultApplyFieldMaskEcomorphsEntity(ctx, &pbObj, in, updateMask, "", e.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(api.EcomorphsEntityWithBeforePatchSave); ok {
		if e.db, err = hook.BeforePatchSave(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := api.DefaultStrictUpdateEcomorphsEntity(ctx, &pbObj, e.db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(api.EcomorphsEntityWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (e EcomorphsEntityRepositoryImpl) GetListEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity, request *api.PagesRequest) ([]*api.EcomorphsEntity, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithBeforeListApplyQuery); ok {
		if e.db, err = hook.BeforeListApplyQuery(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithBeforeListFind); ok {
		if e.db, err = hook.BeforeListFind(ctx, e.db); err != nil {
			return nil, err
		}
	}

	e.db = e.db.Where(&ormObj)
	if request != nil && request.Page != 0 && request.Limit != 0 {
		offset := (request.Page - 1) * request.Limit
		e.db = e.db.Where(&ormObj).Offset(int(offset)).Limit(int(request.Limit))
	} else {
		e.db = e.db.Where(&ormObj)
	}
	e.db = e.db.Order("id")
	ormResponse := []api.EcomorphsEntityORM{}
	if err := e.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.EcomorphsEntityORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, e.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*api.EcomorphsEntity{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
