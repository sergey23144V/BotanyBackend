package repository

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/pkg/errors"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/genproto/protobuf/field_mask"
	"gorm.io/gorm"
)

type AnalysisRepository interface {
	CreatAnalysis(context.Context, *api.Analysis) (*api.Analysis, error)
	RepeatedAnalysis(context.Context, *api.Analysis, *field_mask.FieldMask) (*api.Analysis, error)
	GetAnalysis(context.Context, *api.Analysis) (*api.Analysis, error)
	GetListAnalysis(context.Context, *api.Analysis, *api.PagesRequest) ([]*api.Analysis, error)
	DeleteAnalysis(context.Context, *api.Analysis) error
}

type AnalysisRepositoryImpl struct {
	db *gorm.DB
}

func NewAnalysisRepositoryImpl(db *gorm.DB) AnalysisRepository {
	return AnalysisRepositoryImpl{db}
}

func (a AnalysisRepositoryImpl) CreatAnalysis(ctx context.Context, in *api.Analysis) (*api.Analysis, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = a.db.Omit("Transect").Preload("Transect").Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func (a AnalysisRepositoryImpl) RepeatedAnalysis(ctx context.Context, in *api.Analysis, updateMask *field_mask.FieldMask) (*api.Analysis, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj api.Analysis
	var err error

	pbReadRes, err := api.DefaultReadAnalysis(ctx, &api.Analysis{Id: in.GetId()}, a.db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes

	if _, err := api.DefaultApplyFieldMaskAnalysis(ctx, &pbObj, in, updateMask, "", a.db); err != nil {
		return nil, err
	}
	pbResponse, err := api.DefaultStrictUpdateAnalysis(ctx, &pbObj, a.db)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}

func (a AnalysisRepositoryImpl) GetAnalysis(ctx context.Context, in *api.Analysis) (*api.Analysis, error) {
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

	ormResponse := api.AnalysisORM{}
	if err = a.db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}

	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

func (a AnalysisRepositoryImpl) GetListAnalysis(ctx context.Context, in *api.Analysis, request *api.PagesRequest) ([]*api.Analysis, error) {
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	if request != nil && request.Page != 0 && request.Limit != 0 {
		offset := (request.Page - 1) * request.Limit
		a.db = a.db.Where(&ormObj).Offset(int(offset)).Limit(int(request.Limit))
	} else {
		a.db = a.db.Where(&ormObj)
	}
	a.db = a.db.Order("id")
	ormResponse := []api.AnalysisORM{}
	if err := a.db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}

	pbResponse := []*api.Analysis{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

func (a AnalysisRepositoryImpl) DeleteAnalysis(ctx context.Context, in *api.Analysis) error {
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

	err = a.db.Where(&ormObj).Delete(&api.AnalysisORM{}).Error
	if err != nil {
		return err
	}

	return err
}
