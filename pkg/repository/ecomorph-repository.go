package repository

import (
	"context"
	er "errors"
	"fmt"
	"github.com/infobloxopen/protoc-gen-gorm/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
)

type EcomorphRepository interface {
	CreateEcomorph(ctx context.Context, ecomorph *api.Ecomorph) (*api.Ecomorph, error)
	GetEcomorphById(ctx context.Context, in *api.Ecomorph) (*api.Ecomorph, error)
	DeleteEcomorph(ctx context.Context, in *api.Ecomorph, userRole api.RoleType) error
	StrictUpdateEcomorph(ctx context.Context, in *api.Ecomorph) (*api.Ecomorph, error)
	UpdateEcomorph(ctx context.Context, in *api.Ecomorph, updateMask *field_mask.FieldMask, userRole api.RoleType) (*api.Ecomorph, error)
	GetListEcomorph(ctx context.Context, in *api.Ecomorph, request *api.EcomorphListRequest) ([]*api.Ecomorph, error)
	GetWhereList(filter []*api.Ecomorph) []clause.Expression
	GetListEcomorphById(ctx context.Context, in []clause.Expression) ([]*api.Ecomorph, error)
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
	if err = e.db.Create(&ormObj).Error; err != nil {
		return nil, err
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
	ormResponse := api.EcomorphORM{}
	if err = e.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (e EcomorphRepositoryImpl) GetListEcomorphById(ctx context.Context, in []clause.Expression) ([]*api.Ecomorph, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormResponse := []*api.EcomorphORM{}
	result := e.db.Clauses(in...).Find(&ormResponse)
	if result.Error != nil {
		return nil, result.Error
	}
	pbResponse := []*api.Ecomorph{}
	for _, orm := range ormResponse {
		pb, err := orm.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &pb)
	}

	return pbResponse, nil
}

func (e EcomorphRepositoryImpl) DeleteEcomorph(ctx context.Context, in *api.Ecomorph, userRole api.RoleType) error {
	if in == nil {
		return errors.NilArgumentError
	}
	pbReadRes, err := api.DefaultReadEcomorph(ctx, &api.Ecomorph{Id: in.GetId()}, e.db)
	if err != nil {
		return err
	}

	if pbReadRes.UserId == nil && userRole != api.RoleType_SuperUser {
		return er.New("has no rights")
	}

	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.EmptyIdError
	}
	err = e.db.Where(&ormObj).Delete(&api.EcomorphORM{}).Error
	if err != nil {
		return err
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

	if err = e.db.Save(&ormObj).Error; err != nil {
		return nil, err
	}

	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

func (e EcomorphRepositoryImpl) UpdateEcomorph(ctx context.Context, in *api.Ecomorph, updateMask *field_mask.FieldMask, userRole api.RoleType) (*api.Ecomorph, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.Ecomorph
	var err error

	pbReadRes, err := api.DefaultReadEcomorph(ctx, &api.Ecomorph{Id: in.GetId()}, e.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes

	if pbReadRes.UserId == nil && userRole != api.RoleType_SuperUser {
		return nil, er.New("has no rights")
	}
	if _, err := api.DefaultApplyFieldMaskEcomorph(ctx, &pbObj, in, updateMask, "", e.db); err != nil {
		return nil, err
	}

	pbResponse, err := api.DefaultStrictUpdateEcomorph(ctx, &pbObj, e.db)
	if err != nil {
		return nil, err
	}

	return pbResponse, nil
}

func (e EcomorphRepositoryImpl) GetListEcomorph(ctx context.Context, in *api.Ecomorph, request *api.EcomorphListRequest) ([]*api.Ecomorph, error) {
	expression := e.GetWhereListFromEcomorphListRequest(request.Filter)
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	e.db = e.db.Where(&ormObj)
	if request.Page != nil && request.Page.Page != 0 && request.Page.Limit != 0 {
		offset := (request.Page.Page - 1) * request.Page.Limit
		e.db = e.db.Where(&ormObj).Offset(int(offset)).Limit(int(request.Page.Limit))
	} else {
		e.db = e.db.Where(&ormObj)
	}

	e.db = e.db.Order("id").Clauses(expression...).Or("user_id IS NULL")
	ormResponse := []api.EcomorphORM{}
	if err := e.db.Find(&ormResponse).Error; err != nil {
		return nil, err
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

func (s EcomorphRepositoryImpl) GetWhereList(filter []*api.Ecomorph) []clause.Expression {
	// Объявление переменных
	var conditions []clause.Expression

	// Проверка наличия фильтра по Ids
	if filter != nil {
		// Преобразование Ids в массив интерфейсов
		var interfaceIds []interface{}
		for _, ecomorph := range filter {
			interfaceIds = append(interfaceIds, ecomorph.Id.ResourceId)
		}

		// Добавление фильтра по Ids к условиям
		conditions = append(conditions, clause.IN{
			Column: clause.Column{Name: "id"},
			Values: interfaceIds,
		})
	}

	return conditions
}

func (s EcomorphRepositoryImpl) GetWhereListFromEcomorphListRequest(filter *api.FilterEcomorph) []clause.Expression {
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
			Column: clause.Column{Name: "\"ecomorphs\".\"id\""},
			Values: interfaceIds,
		})
	}

	if filter.SearchTitle != nil {

		var interfaceIds []interface{}

		interfaceIds = append(interfaceIds, filter.SearchTitle)

		conditions = append(conditions, clause.Expr{
			SQL:  "title ~ ?",
			Vars: interfaceIds,
		})
	}

	return conditions
}
