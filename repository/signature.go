package repository

import (
	"gorm.io/gorm"
	"test-signer/entity"
	"test-signer/util"
)

type SignatureRepository struct {
	db *gorm.DB
}

func (s SignatureRepository) Init(sr *util.SharedResources) {
	s.db = sr.Db
}

func (s SignatureRepository) Save(signature entity.Signature) error {
	return s.db.Save(signature).Error
}
