package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-signer/entity"
	"test-signer/entity/request"
	"test-signer/service"
	"test-signer/util"
)

type SignatureController struct {
	testService      *service.TestService
	signatureService *service.SignatureService
}

func (s SignatureController) Register(r *gin.Engine) {
	r.POST("/signature", s.sign)
	r.POST("/signature/verify", s.verify)
}

func (s SignatureController) Init(sr *util.SharedResources) {
	s.signatureService.Init(sr)
	s.testService.Init(sr)
}

func (s SignatureController) sign(c *gin.Context) {
	testRequest := &request.TestRequest{}
	err := c.ShouldBindJSON(testRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{Error: "invalid_request_body"})
		return
	}
	test := &testRequest.Test
	test, err = s.testService.Save(*test)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Error: "failed_to_save_test"})
		return
	}
	signature, err := s.signatureService.Sign(*test)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Error: "failed_to_sign_test"})
		return
	}
	c.JSON(http.StatusOK, entity.Response{Data: signature})
}

func (s SignatureController) verify(c *gin.Context) {
	c.ShouldBindJSON()
}
