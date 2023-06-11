package api

import (
	"html/template"

	"github.com/gin-gonic/gin"

	"github.com/xiexianbin/gin-incubator/common/utils"
	"github.com/xiexianbin/gin-incubator/pkg/middleware"
)

var r *gin.Engine

func InitRouter() *gin.Engine {
	// set gin run mode
	// gin.SetMode(gin.ReleaseMode)

	r = gin.Default()

	// set middleware

	// 加载自定义模板函数（必须在load模板前）
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": utils.UnixToTime,
	})

	// 加载模板
	r.LoadHTMLGlob("templates/**/**")
	//r.LoadHTMLFiles("templates/index.html", "templates/about.html")

	// 加载静态文件
	r.Static("/static", "./static")

	middleware.InitMiddleware(r)

	// 路由 + 中间件示例
	// r.GET("/", initMiddleware, initMiddleware, func(c *gin.Context) {
	// 	c.String(http.StatusOK, "home page...")
	// })
	// 注册主路由
	// router.DefaultRoutersInit(r)
	// router.AdminRoutersInit(r)

	return r
}

func GetRouterGroup(group string) *gin.RouterGroup {
	return r.Group(group)
}
