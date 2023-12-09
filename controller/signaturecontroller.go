package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-signer/entity"
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
	test := &entity.Test{}
	err := c.ShouldBindJSON(test)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{Error: "invalid_request_body"})
		return
	}
	err = s.testService.Save(*test)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Error: "failed_to_save_test"})
		return
	}
}

func (s SignatureController) verify(c *gin.Context) {

}
