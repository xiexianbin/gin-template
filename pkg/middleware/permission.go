package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xiexianbin/gin-incubator/common/utils"
	auth "github.com/xiexianbin/gin-incubator/pkg/auth"
)

// 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("JWT_PAYLOAD")
		v := data.(auth.MapClaims)
		e, err := auth.Casbin("dbtype", "conn")
		utils.HasError(err, "", 500)
		//检查权限
		res, err := e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		// logger.Infof(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		utils.HasError(err, "", 500)

		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  fmt.Sprintf("对不起，您没有 <%v-%v> 访问权限，请联系管理员", c.Request.URL.Path, c.Request.Method),
			})
			c.Abort()
			return
		}
	}
}
