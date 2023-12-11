package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test-signer/entity"
	"test-signer/repository"
	"test-signer/util"
	"time"
)

type SignatureService struct {
	repository *repository.SignatureRepository
	db         gorm.DB
}

func (s SignatureService) Init(sr *util.SharedResources) {
	s.repository.Init(sr)
}

func (s SignatureService) Sign(signable entity.Signable, UserUUID uuid.UUID) (entity.Signature, error) {
	signature := entity.Signature{
		Uuid:      uuid.New(),
		User:      UserUUID,
		Timestamp: time.Now(),
	}
	signable.Sign(signature)
	s.db.Save(&signable)
	return entity.Signature{}, nil
}
