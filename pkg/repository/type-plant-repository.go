package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"

	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
)

type TypePlantRepository interface {
	CreateTypePlant(ctx context.Context, in *api.TypePlant) (*api.TypePlant, error)
	GetTypePlantById(ctx context.Context, in *api.TypePlant) (*api.TypePlant, error)
	DeleteTypePlant(ctx context.Context, in *api.TypePlant) error
	StrictUpdateTypePlant(ctx context.Context, in *api.TypePlant) (*api.TypePlant, error)
	UpdateTypePlant(ctx context.Context, in *api.TypePlant, updateMask *field_mask.FieldMask) (*api.TypePlant, error)
	GetListTypePlant(ctx context.Context, in *api.TypePlant, request *api.PagesRequest) ([]*api.TypePlant, error)
}

type TypePlantRepositoryImpl struct {
	db *gorm.DB
}

func NewTypePlantRepositoryImpl(db *gorm.DB) TypePlantRepository {
	return TypePlantRepositoryImpl{db}
}

func (t TypePlantRepositoryImpl) CreateTypePlant(ctx context.Context, in *api.TypePlant) (*api.TypePlant, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithBeforeCreate_); ok {
		if t.db, err = hook.BeforeCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (t TypePlantRepositoryImpl) GetTypePlantById(ctx context.Context, in *api.TypePlant) (*api.TypePlant, error) {
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
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithBeforeReadApplyQuery); ok {
		if t.db, err = hook.BeforeReadApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithBeforeReadFind); ok {
		if t.db, err = hook.BeforeReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	ormResponse := api.TypePlantORM{}
	if err = t.db.Where(&ormObj).Preload("EcomorphsEntity").Preload("Img").First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(api.TypePlantORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (t TypePlantRepositoryImpl) DeleteTypePlant(ctx context.Context, in *api.TypePlant) error {
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
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithBeforeDelete_); ok {
		if t.db, err = hook.BeforeDelete_(ctx, t.db); err != nil {
			return err
		}
	}
	err = t.db.Where(&ormObj).Delete(&api.TypePlantORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, t.db)
	}
	return err
}

func (t TypePlantRepositoryImpl) StrictUpdateTypePlant(ctx context.Context, in *api.TypePlant) (*api.TypePlant, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTypePlant")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &api.TypePlantORM{}
	t.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithBeforeStrictUpdateCleanup); ok {
		if t.db, err = hook.BeforeStrictUpdateCleanup(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Model(&ormObj).Association("EcomorphsEntity").Replace(ormObj.EcomorphsEntity); err != nil {
		return nil, err
	}
	ormObj.EcomorphsEntity = nil
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithBeforeStrictUpdateSave); ok {
		if t.db, err = hook.BeforeStrictUpdateSave(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, t.db); err != nil {
			return nil, err
		}
	}
	ormResponse := api.TypePlantORM{}
	if err = t.db.Where(&ormObj).Preload("EcomorphsEntity").First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

func (t TypePlantRepositoryImpl) UpdateTypePlant(ctx context.Context, in *api.TypePlant, updateMask *field_mask.FieldMask) (*api.TypePlant, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.TypePlant
	var err error
	if hook, ok := interface{}(&pbObj).(api.TypePlantWithBeforePatchRead); ok {
		if t.db, err = hook.BeforePatchRead(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := t.GetTypePlantById(ctx, &api.TypePlant{Id: in.GetId()})
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(api.TypePlantWithBeforePatchApplyFieldMask); ok {
		if t.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	if in.EcomorphsEntity != nil {
		for _, input := range pbObj.EcomorphsEntity {
			in.EcomorphsEntity = append(in.EcomorphsEntity, input)
		}

		updateMask.Paths = append(updateMask.Paths, "EcomorphsEntity")
	}
	if _, err := api.DefaultApplyFieldMaskTypePlant(ctx, &pbObj, in, updateMask, "", t.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(api.TypePlantWithBeforePatchSave); ok {
		if t.db, err = hook.BeforePatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := t.StrictUpdateTypePlant(ctx, &pbObj)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(api.TypePlantWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (t TypePlantRepositoryImpl) GetListTypePlant(ctx context.Context, in *api.TypePlant, request *api.PagesRequest) ([]*api.TypePlant, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithBeforeListApplyQuery); ok {
		if t.db, err = hook.BeforeListApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithBeforeListFind); ok {
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
	ormResponse := []api.TypePlantORM{}
	if err := t.db.Preload("Img").Preload("EcomorphsEntity").Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.TypePlantORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, t.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*api.TypePlant{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
