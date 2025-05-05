// Copyright 2024 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/xiexianbin/gin-template/config"
	"github.com/xiexianbin/gin-template/internal/api"
	"github.com/xiexianbin/gin-template/pkg/env"
	"github.com/xiexianbin/gin-template/pkg/log"
	"github.com/xiexianbin/gin-template/pkg/orm"
	"github.com/xiexianbin/gin-template/pkg/redis"
)

//	@title			gin-template Swagger API
//	@version		1.0
//	@description	This is a sample server gin-template demo server.
//	@termsOfService	https://github.com/xiexianbin/gin-template

//	@contact.name	API Support
//	@contact.url	http://www.xiexianbin.cn
//	@contact.email	me@xiexianbin.cn

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		127.0.0.1:8888
// @BasePath	/api
func main() {
	// load env
	env.Load()

	// initialize all compontents
	log.Init()
	orm.Init(config.MSYQL_DSN)
	redis.Init(config.REDIS_DSN, config.REDIS_SENTINEL_DSN, config.REDIS_CLUTER_DSN)

	// new app
	app := api.NewAppEngine()
	app.Init()

	// start app
	if err := app.Engine.Run(":8888"); err != nil {
		log.Logger.Sugar().Fatalf("server startup failed: %s", err)
	}
}
