package repository

import (
	context "context"
	"fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	"github.com/jinzhu/gorm"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	ecomorph_entity "github.com/sergey23144V/BotanyBackend/servers/g-rpc/api/ecomorph-entity"
	"google.golang.org/genproto/protobuf/field_mask"
)

type EcomorphsEntityRepository interface {
	CreateEcomorphsEntity(ctx context.Context, ecomorph *ecomorph_entity.EcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error)
	GetEcomorphsEntityById(ctx context.Context, in *ecomorph_entity.EcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error)
	DeleteEcomorphsEntity(ctx context.Context, in *ecomorph_entity.EcomorphsEntity) error
	StrictUpdateEcomorphsEntity(ctx context.Context, in *ecomorph_entity.EcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error)
	UpdateEcomorphsEntity(ctx context.Context, in *ecomorph_entity.EcomorphsEntity, updateMask *field_mask.FieldMask) (*ecomorph_entity.EcomorphsEntity, error)
	GetListEcomorphsEntity(ctx context.Context, in *ecomorph_entity.EcomorphsEntity) ([]*ecomorph_entity.EcomorphsEntity, error)
}

type EcomorphsEntityRepositoryImpl struct {
	db *gorm.DB
}

func NewEcomorphsEntityRepositoryImpl(db *gorm.DB) EcomorphsEntityRepository {
	return EcomorphsEntityRepositoryImpl{db}
}

func (e EcomorphsEntityRepositoryImpl) CreateEcomorphsEntity(ctx context.Context, in *ecomorph_entity.EcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithBeforeCreate_); ok {
		if e.db, err = hook.BeforeCreate_(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (e EcomorphsEntityRepositoryImpl) GetEcomorphsEntityById(ctx context.Context, in *ecomorph_entity.EcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
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
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithBeforeReadApplyQuery); ok {
		if e.db, err = hook.BeforeReadApplyQuery(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if e.db, err = gorm1.ApplyFieldSelection(ctx, e.db, nil, &ecomorph_entity.EcomorphsEntityORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithBeforeReadFind); ok {
		if e.db, err = hook.BeforeReadFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	ormResponse := ecomorph_entity.EcomorphsEntityORM{}
	if err = e.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ecomorph_entity.EcomorphsEntityORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (e EcomorphsEntityRepositoryImpl) DeleteEcomorphsEntity(ctx context.Context, in *ecomorph_entity.EcomorphsEntity) error {
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
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithBeforeDelete_); ok {
		if e.db, err = hook.BeforeDelete_(ctx, e.db); err != nil {
			return err
		}
	}
	tx := e.db.Where(&ormObj).Delete(&ecomorph_entity.EcomorphsEntityORM{})
	if tx.Error != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, e.db)
	}
	if tx.RowsAffected == 0 {
		return errors.NotDelete
	}
	return err
}

func (e EcomorphsEntityRepositoryImpl) StrictUpdateEcomorphsEntity(ctx context.Context, in *ecomorph_entity.EcomorphsEntity) (*ecomorph_entity.EcomorphsEntity, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateEcomorphsEntity")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ecomorph_entity.EcomorphsEntityORM{}
	e.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithBeforeStrictUpdateCleanup); ok {
		if e.db, err = hook.BeforeStrictUpdateCleanup(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Model(&ormObj).Association("Ecomorphs").Replace(ormObj.Ecomorphs).Error; err != nil {
		return nil, err
	}
	ormObj.Ecomorphs = nil
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithBeforeStrictUpdateSave); ok {
		if e.db, err = hook.BeforeStrictUpdateSave(ctx, e.db); err != nil {
			return nil, err
		}
	}
	if err = e.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithAfterStrictUpdateSave); ok {
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

func (e EcomorphsEntityRepositoryImpl) UpdateEcomorphsEntity(ctx context.Context, in *ecomorph_entity.EcomorphsEntity, updateMask *field_mask.FieldMask) (*ecomorph_entity.EcomorphsEntity, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj ecomorph_entity.EcomorphsEntity
	var err error
	if hook, ok := interface{}(&pbObj).(ecomorph_entity.EcomorphsEntityWithBeforePatchRead); ok {
		if e.db, err = hook.BeforePatchRead(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := ecomorph_entity.DefaultReadEcomorphsEntity(ctx, &ecomorph_entity.EcomorphsEntity{Id: in.GetId()}, e.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ecomorph_entity.EcomorphsEntityWithBeforePatchApplyFieldMask); ok {
		if e.db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	if _, err := ecomorph_entity.DefaultApplyFieldMaskEcomorphsEntity(ctx, &pbObj, in, updateMask, "", e.db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ecomorph_entity.EcomorphsEntityWithBeforePatchSave); ok {
		if e.db, err = hook.BeforePatchSave(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ecomorph_entity.DefaultStrictUpdateEcomorphsEntity(ctx, &pbObj, e.db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ecomorph_entity.EcomorphsEntityWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, e.db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

func (e EcomorphsEntityRepositoryImpl) GetListEcomorphsEntity(ctx context.Context, in *ecomorph_entity.EcomorphsEntity) ([]*ecomorph_entity.EcomorphsEntity, error) {

	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithBeforeListApplyQuery); ok {
		if e.db, err = hook.BeforeListApplyQuery(ctx, e.db); err != nil {
			return nil, err
		}
	}
	e.db, err = gorm1.ApplyCollectionOperators(ctx, e.db, &ecomorph_entity.EcomorphsEntityORM{}, &ecomorph_entity.EcomorphsEntity{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithBeforeListFind); ok {
		if e.db, err = hook.BeforeListFind(ctx, e.db); err != nil {
			return nil, err
		}
	}
	e.db = e.db.Where(&ormObj)
	e.db = e.db.Order("id")
	ormResponse := []ecomorph_entity.EcomorphsEntityORM{}
	if err := e.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ecomorph_entity.EcomorphsEntityORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, e.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*ecomorph_entity.EcomorphsEntity{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}
