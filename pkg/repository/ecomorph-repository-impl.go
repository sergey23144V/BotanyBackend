package repository

import (
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"gorm.io/gorm"
)

type EcomorphRepositoryImpl struct {
	db gorm.DB
}

func (e EcomorphRepositoryImpl) Insert(input g_rpc.Ecomorph) (*g_rpc.Ecomorph, error) {
	if err := e.db.Create(input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (e EcomorphRepositoryImpl) GetById(inputId string) (*g_rpc.Ecomorph, error) {
	var result g_rpc.Ecomorph
	if err := e.db.Where("id = ?", inputId).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (e EcomorphRepositoryImpl) GetList() ([]*g_rpc.Ecomorph, error) {
	var result []*g_rpc.Ecomorph
	if err := e.db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func NewEcomorphsRepositoryImpl(db *gorm.DB) EcomorphRepositoryImpl {
	return EcomorphRepositoryImpl{db: *db}
}
