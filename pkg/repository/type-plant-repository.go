package repository

import (
	"context"
	"fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	type_plant "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/type-plant"
	"google.golang.org/genproto/protobuf/field_mask"
)

type TypePlantRepository interface {
	CreateTypePlant(ctx context.Context, in *type_plant.TypePlant) (*type_plant.TypePlant, error)
	GetTypePlantById(ctx context.Context, in *type_plant.TypePlant) (*type_plant.TypePlant, error)
	DeleteTypePlant(ctx context.Context, in *type_plant.TypePlant) error
	StrictUpdateTypePlant(ctx context.Context, in *type_plant.TypePlant) (*type_plant.TypePlant, error)
	UpdateTypePlant(ctx context.Context, in *type_plant.TypePlant, updateMask *field_mask.FieldMask) (*type_plant.TypePlant, error)
	GetListTypePlant(ctx context.Context, in *type_plant.TypePlant) ([]*type_plant.TypePlant, error)
}

type TypePlantRepositoryImpl struct {
	db *gorm.DB
}

func NewTypePlantRepositoryImpl(db *gorm.DB) TypePlantRepository {
	return TypePlantRepositoryImpl{db}
}

func (t TypePlantRepositoryImpl) CreateTypePlant(ctx context.Context, in *type_plant.TypePlant) (*type_plant.TypePlant, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithBeforeCreate_); ok {
		if t.db, err = hook.BeforeCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (t TypePlantRepositoryImpl) GetTypePlantById(ctx context.Context, in *type_plant.TypePlant) (*type_plant.TypePlant, error) {
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
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithBeforeReadApplyQuery); ok {
		if t.db, err = hook.BeforeReadApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if t.db, err = gorm1.ApplyFieldSelection(ctx, t.db, nil, &type_plant.TypePlantORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithBeforeReadFind); ok {
		if t.db, err = hook.BeforeReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	ormResponse := type_plant.TypePlantORM{}
	if err = t.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(type_plant.TypePlantORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (t TypePlantRepositoryImpl) DeleteTypePlant(ctx context.Context, in *type_plant.TypePlant) error {
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
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithBeforeDelete_); ok {
		if t.db, err = hook.BeforeDelete_(ctx, t.db); err != nil {
			return err
		}
	}
	err = t.db.Where(&ormObj).Delete(&type_plant.TypePlantORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, t.db)
	}
	return err
}

func (t TypePlantRepositoryImpl) StrictUpdateTypePlant(ctx context.Context, in *type_plant.TypePlant) (*type_plant.TypePlant, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTypePlant")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &type_plant.TypePlantORM{}
	t.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithBeforeStrictUpdateCleanup); ok {
		if t.db, err = hook.BeforeStrictUpdateCleanup(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Model(&ormObj).Association("EcomorphsEntity").Replace(ormObj.EcomorphsEntity).Error; err != nil {
		return nil, err
	}
	ormObj.EcomorphsEntity = nil
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithBeforeStrictUpdateSave); ok {
		if t.db, err = hook.BeforeStrictUpdateSave(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if err = t.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithAfterStrictUpdateSave); ok {
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

func (t TypePlantRepositoryImpl) UpdateTypePlant(ctx context.Context, in *type_plant.TypePlant, updateMask *field_mask.FieldMask) (*type_plant.TypePlant, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj type_plant.TypePlant
	var err error
	if hook, ok := interface{}(&pbObj).(type_plant.TypePlantWithBeforePatchRead); ok {
		if t.db, err = hook.BeforePatchRead(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := type_plant.DefaultReadTypePlant(ctx, &type_plant.TypePlant{Id: in.GetId()}, t.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(type_plant.TypePlantWithBeforePatchApplyFieldMask); ok {
		if t.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	if _, err := type_plant.DefaultApplyFieldMaskTypePlant(ctx, &pbObj, in, updateMask, "", t.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(type_plant.TypePlantWithBeforePatchSave); ok {
		if t.db, err = hook.BeforePatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := type_plant.DefaultStrictUpdateTypePlant(ctx, &pbObj, t.db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(type_plant.TypePlantWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, t.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (t TypePlantRepositoryImpl) GetListTypePlant(ctx context.Context, in *type_plant.TypePlant) ([]*type_plant.TypePlant, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithBeforeListApplyQuery); ok {
		if t.db, err = hook.BeforeListApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	t.db, err = gorm1.ApplyCollectionOperators(ctx, t.db, &type_plant.TypePlantORM{}, &type_plant.TypePlant{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithBeforeListFind); ok {
		if t.db, err = hook.BeforeListFind(ctx, t.db); err != nil {
			return nil, err
		}
	}
	t.db = t.db.Where(&ormObj)
	t.db = t.db.Order("id")
	ormResponse := []type_plant.TypePlantORM{}
	if err := t.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(type_plant.TypePlantORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, t.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*type_plant.TypePlant{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
