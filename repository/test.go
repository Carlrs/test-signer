package repository

import (
	"fmt"
	"github.com/google/uuid"
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

func (t TestRepository) FindBySignature(signature *entity.Signature, userUUID uuid.UUID) (*entity.Test, error) {
	test := entity.Test{}
	err := t.db.Model(&test).Where("uuid = ? AND user = ?", signature.Uuid, userUUID).Association("Signature").Find(&test)
	fmt.Printf("Test: %v\n", test)
	return &test, err
}
