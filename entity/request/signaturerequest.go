package request

import "test-signer/entity"

type SignatureRequest struct {
	UserJWT   string           `json:"userJWT"`
	Signature entity.Signature `json:"signature"`
}
