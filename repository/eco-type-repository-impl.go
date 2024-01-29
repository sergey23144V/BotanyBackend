package repository

import (
	"BotanyBackend/servers/g-rpc/Type"
	"gorm.io/gorm"
)

type EcoTypeRepositoryImpl struct {
	db gorm.DB
}

func NewEcoTypeRepositoryImpl(db *gorm.DB) EcoTypeRepository {
	return EcoTypeRepositoryImpl{db: *db}
}

func (e EcoTypeRepositoryImpl) Insert(input Type.EcoEntity) (*Type.EcoEntity, error) {
	tx := e.db.Create(input)

	if tx.Error != nil {
		return &input, nil
	}
	return nil, tx.Error
}

func (e EcoTypeRepositoryImpl) GetById(inputId string) (Type.EcoEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (e EcoTypeRepositoryImpl) GetList() ([]Type.EcoEntity, error) {
	//TODO implement me
	panic("implement me")
}
