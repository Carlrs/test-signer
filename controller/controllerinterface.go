package controller

import (
	"github.com/gin-gonic/gin"
	"test-signer/util"
)

type Controller interface {
	Init(sr *util.SharedResources)
	Register(r *gin.Engine)
}
