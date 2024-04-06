package repository

import (
	context "context"
	"fmt"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EcomorphsEntityRepository interface {
	CreateEcomorphsEntity(ctx context.Context, ecomorph *api.EcomorphsEntity) (*api.EcomorphsEntity, error)
	GetEcomorphsEntityById(ctx context.Context, in *api.EcomorphsEntity) (*api.EcomorphsEntity, error)
	DeleteEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity) error
	StrictUpdateEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity) (*api.EcomorphsEntity, error)
	UpdateEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity, updateMask *field_mask.FieldMask) (*api.EcomorphsEntity, error)
	GetListEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity, request *api.EcomorphsEntityListRequest) ([]*api.EcomorphsEntity, error)
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

	tx := e.db.Delete(&ormObj)
	if tx.Error != nil {
		return err
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

	if err = e.db.Model(&ormObj).Association("Ecomorphs").Replace(ormObj.Ecomorphs); err != nil {
		return nil, err
	}
	ormObj.Ecomorphs = nil

	if err = e.db.Save(&ormObj).Error; err != nil {
		return nil, err
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

	pbReadRes, err := api.DefaultReadEcomorphsEntity(ctx, &api.EcomorphsEntity{Id: in.GetId()}, e.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes

	if _, err := api.DefaultApplyFieldMaskEcomorphsEntity(ctx, &pbObj, in, updateMask, "", e.db); err != nil {
		return nil, err
	}

	pbResponse, err := api.DefaultStrictUpdateEcomorphsEntity(ctx, &pbObj, e.db)
	if err != nil {
		return nil, err
	}

	return pbResponse, nil
}

func (e EcomorphsEntityRepositoryImpl) GetListEcomorphsEntity(ctx context.Context, in *api.EcomorphsEntity, request *api.EcomorphsEntityListRequest) ([]*api.EcomorphsEntity, error) {
	expression := e.GetWhereList(request.Filter)
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	e.db = e.db.Where(&ormObj)
	if request != nil && request.Page.Page != 0 && request.Page.Limit != 0 {
		offset := (request.Page.Page - 1) * request.Page.Limit
		e.db = e.db.Where(&ormObj).Offset(int(offset)).Limit(int(request.Page.Limit))
	} else {
		e.db = e.db.Where(&ormObj)
	}
	e.db = e.db.Order("id")
	ormResponse := []api.EcomorphsEntityORM{}
	if err := e.db.Clauses(expression...).Find(&ormResponse).Error; err != nil {
		return nil, err
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

func (s EcomorphsEntityRepositoryImpl) GetWhereList(filter *api.FilterEcomorphsEntity) []clause.Expression {
	// Объявление переменных
	var conditions []clause.Expression
	if filter == nil {
		return nil
	}
	// Проверка наличия фильтра по Ids
	if filter.Id != nil {
		// Преобразование Ids в массив интерфейсов
		var interfaceIds []interface{}
		for _, id := range filter.Id {
			interfaceIds = append(interfaceIds, id.ResourceId)
		}

		// Добавление фильтра по Ids к условиям
		conditions = append(conditions, clause.IN{
			Column: clause.Column{Name: "id"},
			Values: interfaceIds,
		})
	}

	if filter.SearchTitle != nil {

		var interfaceIds []interface{}

		interfaceIds = append(interfaceIds, filter.SearchTitle)

		conditions = append(conditions, clause.IN{
			Column: clause.Column{Name: "title"},
			Values: interfaceIds,
		})
	}

	if filter.Title != nil {

		var interfaceIds []interface{}
		var columns []clause.OrderByColumn

		interfaceIds = append(interfaceIds, filter.SearchTitle)

		columns = append(columns, clause.OrderByColumn{
			Column: clause.Column{Name: "title"},
			Desc:   *filter.Title == api.Direction_DESCENDING,
		})
		conditions = append(conditions, clause.OrderBy{
			Columns:    columns,
			Expression: nil,
		})
	}

	return conditions
}
