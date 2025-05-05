// Copyright 2024 xiexianbin<me@xiexianbin.cn>
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

package api

import (
	"sync"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/xiexianbin/gin-template/docs"
	"github.com/xiexianbin/gin-template/pkg/ginmiddleware"
)

var (
	once      sync.Once
	appEngine *AppEngine
)

type AppEngine struct {
	Engine *gin.Engine
}

// NewAppEngine returns an Engine instance
func NewAppEngine() *AppEngine {
	once.Do(func() {
		appEngine.Engine = gin.New()
	})
	return appEngine
}

func (e *AppEngine) Init() {
	// Set gin mode
	gin.SetMode(gin.ReleaseMode)

	// Use custom middlewares
	appEngine.Engine.Use(
		ginmiddleware.RateLimiter(),
		ginmiddleware.RequestID(),
		ginmiddleware.Logger(),
		ginmiddleware.Gzip(),
		ginmiddleware.CORS(),
		ginmiddleware.RateLimiter(),
		gin.Recovery(),
	)

	// initialize API routes

	// Static resource dir
	e.Engine.Static("/static", "static")

	// swagger
	e.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
