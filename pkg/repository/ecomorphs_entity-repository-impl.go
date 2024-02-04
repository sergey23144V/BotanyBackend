package repository

import (
	g_rpc "github.com/sergey23144V/BotanyBackend/servers/g-rpc"
	"gorm.io/gorm"
)

type EcomorphsEntityRepositoryImpl struct {
	db gorm.DB
}

func (e EcomorphsEntityRepositoryImpl) Insert(input g_rpc.EcomorphsEntity) (*g_rpc.EcomorphsEntity, error) {
	if err := e.db.Create(input).Preload("Ecomorphs").Find(input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (e EcomorphsEntityRepositoryImpl) GetById(inputId string) (*g_rpc.EcomorphsEntity, error) {
	var result g_rpc.EcomorphsEntity
	if err := e.db.Where("id = ?", inputId).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (e EcomorphsEntityRepositoryImpl) GetList() ([]*g_rpc.EcomorphsEntity, error) {
	var result []*g_rpc.EcomorphsEntity
	if err := e.db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func NewEcomorphsEntityRepositoryImpl(db *gorm.DB) EcomorphsEntityRepository {
	return EcomorphsEntityRepositoryImpl{db: *db}
}
