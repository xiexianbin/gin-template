package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	v1base "github.com/xiexianbin/gin-incubator/api/common"
	"github.com/xiexianbin/gin-incubator/api/v1/request"
)

type UserController struct {
	v1base.V1BaseController
}

func (u UserController) List(c *gin.Context) {
	u.Success(c)
}

func (u *UserController) Register(c *gin.Context) {
	var r request.RegisterRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		fmt.Println("register failed")
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	fmt.Println("save to db ...")
	c.JSON(http.StatusOK, "successful")
}
