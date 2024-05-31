package repository

import (
	"context"
	"fmt"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TrialSiteRepository interface {
	GetListTrialSite(ctx context.Context, in *api.TrialSite, request *api.TrialSiteListRequest) ([]*api.TrialSite, error)
	GetTrialSiteById(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error)
	UpdateTrialSite(ctx context.Context, in *api.TrialSite, updateMask *field_mask.FieldMask) (*api.TrialSite, error)
	StrictUpdateTrialSite(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error)
	DeleteTrialSite(ctx context.Context, in *api.TrialSite) error
	CreateTrialSite(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error)
	CreatePlant(context.Context, *api.Plant) (*api.Plant, error)
	// Получение сущности по id
	GetPlant(context.Context, *api.Plant) (*api.Plant, error)
	// Обновление сущности по id
	UpdatePlant(ctx context.Context, in *api.Plant, updateMask *field_mask.FieldMask) (*api.Plant, error)
	// Удаление сущности по заголовку
	DeletePlant(context.Context, *api.Plant) error
	// Получение списка всех сущностей
	GetAllPlant(context.Context, *api.Plant, *api.PagesRequest) ([]*api.Plant, error)
	GetAllPlantORM(ctx context.Context, ormObj *api.PlantORM, request *api.PagesRequest) ([]*api.Plant, error)
}

type TrialSiteRepositoryImpl struct {
	db *gorm.DB
}

func NewTrialSiteRepositoryImpl(db *gorm.DB) TrialSiteRepository {
	return TrialSiteRepositoryImpl{db}
}

func (t TrialSiteRepositoryImpl) CreatePlant(ctx context.Context, in *api.Plant) (*api.Plant, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	if err = t.db.Omit().Preload("TypePlantId").Create(&ormObj).Error; err != nil {
		return nil, err
	}

	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (t TrialSiteRepositoryImpl) GetPlant(ctx context.Context, in *api.Plant) (*api.Plant, error) {
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
	ormResponse := api.PlantORM{}
	if err = t.db.Where(&ormObj).Preload("TypePlant").First(&ormResponse).Error; err != nil {
		return nil, err
	}

	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (t TrialSiteRepositoryImpl) StrictUpdatePlant(ctx context.Context, in *api.Plant) (*api.Plant, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdatePlant")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &api.PlantORM{}
	t.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)

	if err = t.db.Omit("TypePlant").Omit("TrialSiteId").Preload("TypePlant").Save(&ormObj).Error; err != nil {
		return nil, err
	}

	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

func (t TrialSiteRepositoryImpl) UpdatePlant(ctx context.Context, in *api.Plant, updateMask *field_mask.FieldMask) (*api.Plant, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.Plant
	var err error
	pbReadRes, err := api.DefaultReadPlant(ctx, &api.Plant{Id: in.GetId()}, t.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if _, err := api.DefaultApplyFieldMaskPlant(ctx, &pbObj, in, updateMask, "", t.db); err != nil {
		return nil, err
	}
	pbResponse, err := t.StrictUpdatePlant(ctx, &pbObj)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}

func (t TrialSiteRepositoryImpl) DeletePlant(ctx context.Context, in *api.Plant) error {
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
	err = t.db.Where(&ormObj).Delete(&api.PlantORM{}).Error
	if err != nil {
		return err
	}
	return err
}

func (t TrialSiteRepositoryImpl) GetAllPlant(ctx context.Context, in *api.Plant, request *api.PagesRequest) ([]*api.Plant, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if request != nil && request.Page != 0 && request.Limit != 0 {
		offset := (request.Page - 1) * request.Limit
		t.db = t.db.Where(&ormObj).Offset(int(offset)).Limit(int(request.Limit))
	} else {
		t.db = t.db.Where(&ormObj)
	}
	t.db = t.db.Order("id")
	ormResponse := []api.PlantORM{}
	if err := t.db.Preload("TypePlant").Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*api.Plant{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

func (t TrialSiteRepositoryImpl) GetAllPlantORM(ctx context.Context, ormObj *api.PlantORM, request *api.PagesRequest) ([]*api.Plant, error) {
	var err error

	if hook, ok := interface{}(&ormObj).(api.PlantORMWithBeforeListApplyQuery); ok {
		if t.db, err = hook.BeforeListApplyQuery(ctx, t.db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(api.PlantORMWithBeforeListFind); ok {
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
	ormResponse := []api.PlantORM{}
	if err := t.db.Preload("TypePlant").Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(api.PlantORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, t.db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*api.Plant{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

func (t TrialSiteRepositoryImpl) CreateTrialSite(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error) {
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

func (t TrialSiteRepositoryImpl) GetTrialSiteById(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error) {
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
	ormResponse := api.TrialSiteORM{}
	if err = t.db.Where(&ormObj).Preload("Dominant").Preload("SubDominant").
		Preload("Plant").Preload("Plant.TypePlant").Preload("Img").First(&ormResponse).Error; err != nil {
		return nil, err
	}

	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (t TrialSiteRepositoryImpl) UpdateTrialSite(ctx context.Context, in *api.TrialSite, updateMask *field_mask.FieldMask) (*api.TrialSite, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.TrialSite
	var err error

	pbReadRes, err := api.DefaultReadTrialSite(ctx, &api.TrialSite{Id: in.GetId()}, t.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes

	if _, err := api.DefaultApplyFieldMaskTrialSite(ctx, &pbObj, in, updateMask, "", t.db); err != nil {
		return nil, err
	}

	pbResponse, err := t.StrictUpdateTrialSite(ctx, &pbObj)
	if err != nil {
		return nil, err
	}

	return pbResponse, nil
}

func (t TrialSiteRepositoryImpl) StrictUpdateTrialSite(ctx context.Context, in *api.TrialSite) (*api.TrialSite, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTrialSite")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &api.TrialSiteORM{}
	t.db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)

	if err = t.db.Omit("TransectId").Preload("Dominant").Preload("SubDominant").Save(&ormObj).Error; err != nil {
		return nil, err
	}

	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

func (t TrialSiteRepositoryImpl) DeleteTrialSite(ctx context.Context, in *api.TrialSite) error {
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
	err = t.db.Where(&ormObj).Delete(&api.TrialSiteORM{}).Error
	if err != nil {
		return err
	}
	return err
}

func (t TrialSiteRepositoryImpl) GetListTrialSite(ctx context.Context, in *api.TrialSite, request *api.TrialSiteListRequest) ([]*api.TrialSite, error) {
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
	ormResponse := []api.TrialSiteORM{}
	if err := t.db.Preload("Plant").Preload("Img").Find(&ormResponse).Error; err != nil {
		return nil, err
	}

	pbResponse := []*api.TrialSite{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

func (s TrialSiteRepositoryImpl) GetWhereList(filter *api.FilterTrialSite) []clause.Expression {
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

	if filter.SearchTitle != "" {

		var interfaceIds []interface{}

		interfaceIds = append(interfaceIds, filter.SearchTitle)

		conditions = append(conditions, clause.Expr{
			SQL:  "title ~ ?",
			Vars: interfaceIds,
		})
	}

	return conditions
}
