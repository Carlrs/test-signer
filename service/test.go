package service

import (
	"github.com/google/uuid"
	"test-signer/entity"
	"test-signer/repository"
	"test-signer/util"
)

type TestService struct {
	repository *repository.TestRepository
}

func (t TestService) Init(sr *util.SharedResources) {
	t.repository.Init(sr)
}

func (t TestService) Save(test entity.Test) (*entity.Test, error) {
	return t.repository.Save(test)
}

func (t TestService) FindBySignature(signature *entity.Signature, userUUID uuid.UUID) (*entity.Test, error) {
	return t.repository.FindBySignature(signature, userUUID)
}
