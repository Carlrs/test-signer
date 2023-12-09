package repository

import (
	"gorm.io/gorm"
	"test-signer/entity"
	"test-signer/util"
)

type TestRepository struct {
	db *gorm.DB
}

func (t TestRepository) Init(sr *util.SharedResources) {
	t.db = sr.Db
}

func (t TestRepository) Save(test entity.Test) (*entity.Test, error) {
	err := t.db.Save(test).Error
	return &test, err
}
