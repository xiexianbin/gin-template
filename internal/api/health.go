package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gintemplate "github.com/xiexianbin/gin-template"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "UP",
		"version": gintemplate.GetVersion().Version,
	})
}
