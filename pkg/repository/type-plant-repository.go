package repository

import (
	"context"
	er "errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
)

type TypePlantRepository interface {
	CreateTypePlant(ctx context.Context, in *api.TypePlant) (*api.TypePlant, error)
	GetTypePlantById(ctx context.Context, in *api.TypePlant) (*api.TypePlant, error)
	DeleteTypePlant(ctx context.Context, in *api.TypePlant, userRole api.RoleType) error
	StrictUpdateTypePlant(ctx context.Context, in *api.TypePlant) (*api.TypePlant, error)
	UpdateTypePlant(ctx context.Context, in *api.TypePlant, updateMask *field_mask.FieldMask, userRole api.RoleType) (*api.TypePlant, error)
	GetListTypePlant(ctx context.Context, in *api.TypePlant, request *api.TypePlantListRequest) ([]*api.TypePlant, error)
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
	if err = t.db.Create(&ormObj).Error; err != nil {
		return nil, err
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

	ormResponse := api.TypePlantORM{}
	if err = t.db.Where(&ormObj).Preload("EcomorphsEntity").Preload("Img").First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (t TypePlantRepositoryImpl) DeleteTypePlant(ctx context.Context, in *api.TypePlant, userRole api.RoleType) error {
	if in == nil {
		return errors.NilArgumentError
	}
	pbReadRes, err := api.DefaultReadEcomorphsEntity(ctx, &api.EcomorphsEntity{Id: in.GetId()}, t.db)
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
	err = t.db.Where(&ormObj).Delete(&api.TypePlantORM{}).Error
	if err != nil {
		return err
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
	if err = t.db.Model(&ormObj).Association("EcomorphsEntity").Replace(ormObj.EcomorphsEntity); err != nil {
		return nil, err
	}
	ormObj.EcomorphsEntity = nil
	if err = t.db.Save(&ormObj).Error; err != nil {
		return nil, err
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

func (t TypePlantRepositoryImpl) UpdateTypePlant(ctx context.Context, in *api.TypePlant, updateMask *field_mask.FieldMask, userRole api.RoleType) (*api.TypePlant, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.TypePlant
	var err error
	pbReadRes, err := t.GetTypePlantById(ctx, &api.TypePlant{Id: in.GetId()})
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if pbReadRes.UserId == nil && userRole != api.RoleType_SuperUser {
		return nil, er.New("has no rights")
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
	pbResponse, err := t.StrictUpdateTypePlant(ctx, &pbObj)
	if err != nil {
		return nil, err
	}

	return pbResponse, nil
}

func (t TypePlantRepositoryImpl) GetListTypePlant(ctx context.Context, in *api.TypePlant, request *api.TypePlantListRequest) ([]*api.TypePlant, error) {
	expression := t.GetWhereList(request.Filter)
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	if request.Page != nil && request.Page.Page != 0 && request.Page.Limit != 0 {
		offset := (request.Page.Page - 1) * request.Page.Limit
		t.db = t.db.Where(&ormObj).Offset(int(offset)).Limit(int(request.Page.Limit))
	} else {
		t.db = t.db.Where(&ormObj)
	}
	t.db = t.db.Clauses(expression...).Order("id").Or("user_id IS NULL")
	ormResponse := []api.TypePlantORM{}
	if err := t.db.Preload("Img").Preload("EcomorphsEntity").Find(&ormResponse).Error; err != nil {
		return nil, err
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

func (s TypePlantRepositoryImpl) GetWhereList(filter *api.FilterTypePlant) []clause.Expression {
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

	if filter.EcomorphsEntity != nil {

		var interfaceIds []interface{}
		for _, ecomorphsEntity := range filter.EcomorphsEntity {
			interfaceIds = append(interfaceIds, ecomorphsEntity.Id.ResourceId)
		}

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

	return conditions
}
