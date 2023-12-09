package request

import "test-signer/entity"

type TestRequest struct {
	UserJWT string      `json:"userJWT"`
	Test    entity.Test `json:"test"`
}
