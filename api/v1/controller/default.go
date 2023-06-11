package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1base "github.com/xiexianbin/gin-incubator/api/common"
)

type DefaultController struct {
	v1base.V1BaseController
}

func (d *DefaultController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
