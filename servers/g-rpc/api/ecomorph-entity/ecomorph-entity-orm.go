package ecomorph_entity

import (
	"context"
	"github.com/infobloxopen/protoc-gen-gorm/errors"
	"github.com/jinzhu/gorm"
)

func CreateEcomorphsEntity(ctx context.Context, in *EcomorphsEntity, db *gorm.DB) (*EcomorphsEntity, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(EcomorphsEntityORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit("ecomorphs").Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(EcomorphsEntityORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}
