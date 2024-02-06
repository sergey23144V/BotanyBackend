package user

import (
	"context"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	"github.com/infobloxopen/protoc-gen-gorm/errors"
	"github.com/jinzhu/gorm"
)

func ReadUserByEmailAndPassword(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &UserORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserORM{}
	if err = db.Where("email = ? AND password = ?", ormObj.Email, ormObj.Password).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func CheckingForDuplicateEmails(ctx context.Context, in *User, db *gorm.DB) (bool, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return false, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return false, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &UserORM{}, &User{}, nil, nil, nil, nil)
	if err != nil {
		return false, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return false, err
		}
	}
	db = db.Where("email = ?", ormObj.Email)
	db = db.Order("id")
	ormResponse := []UserORM{}
	tx := db.Find(&ormResponse)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected > 0 {
		return false, nil
	} else {
		return true, nil
	}

}
