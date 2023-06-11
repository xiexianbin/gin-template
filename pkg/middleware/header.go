// Copyright 2023 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")

		// 放行所有 OPTIONS 方法，参考 https://www.xiexianbin.cn/nginx/conf/nginx-proxy_pass/index.html#cors-%E8%B7%A8%E5%9F%9F%E8%AF%B7%E6%B1%82%E4%BB%A3%E7%90%86
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func NoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		c.Next()
	}
}
