package common

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type V1BaseController struct {
	Context *gin.Context
	Orm     *gorm.DB
	// Logger
}

// MakeContext 设置http上下文
func (b *V1BaseController) MakeContext(c *gin.Context) *V1BaseController {
	b.Context = c
	return b
}

func (b *V1BaseController) Success(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func (b *V1BaseController) Error(c *gin.Context) {
	c.String(http.StatusBadRequest, "error")
}

type Pagination struct {
	Limit int    `json:"limit" form:"limit" uri:"limit"`
	Page  int    `json:"page" form:"page" uri:"page"`
	Sort  string `json:"sort" form:"sort" uri:"sort"`
}

func (b *V1BaseController) GeneratePaginationFromRequest(c *gin.Context) (pagination Pagination) {
	if err := c.ShouldBind(&pagination); err != nil {
		fmt.Printf("参数绑定错误:%s\n", err)
	}

	// 校验参数
	if pagination.Limit < 0 {
		pagination.Limit = 2
	}
	if pagination.Page < 1 {
		pagination.Page = 1
	}

	if len(pagination.Sort) == 0 {
		pagination.Sort = "created_at desc"
	}

	return
}

// PageOK 分页数据处理
// https://gorm.io/zh_CN/docs/scopes.html#pagination
func (b *V1BaseController) PageOK(result interface{}) {
}
