package service

import (
	"test-signer/entity"
	"test-signer/repository"
	"test-signer/util"
)

type SignatureService struct {
	repository *repository.SignatureRepository
}

func (s SignatureService) Init(sr *util.SharedResources) {
	s.repository.Init(sr)
}

func (s SignatureService) Sign(test entity.Test) (entity.Signature, error) {

	return entity.Signature{}, nil
}
