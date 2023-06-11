package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xiexianbin/gin-incubator/api"
	"github.com/xiexianbin/gin-incubator/api/v1/controller"
)

func initDefaultRouter() {
	defaultController := controller.DefaultController{}
	router := api.GetRouterGroup("/")
	{
		router.GET("/ping", defaultController.Ping)
	}
}

func initAdminRouter() {
	adminRouter := api.GetRouterGroup("/v1/admin")
	{
		adminRouter.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
		adminRouter.GET("/users", controller.UserController{}.List)
	}
}

func initUserRouter() {
	uc := controller.UserController{}
	userRouter := api.GetRouterGroup("/v1/user")
	{
		userRouter.POST("/register", uc.Register)
	}
}

func InitV1Router() {
	initDefaultRouter()
	initAdminRouter()
	initUserRouter()
}
