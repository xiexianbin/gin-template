// Copyright 2022 xiexianbin<me@xiexianbin.cn>
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
	"github.com/xiexianbin/gin-incubator/api"
	v1 "github.com/xiexianbin/gin-incubator/api/v1"
)

//	@title			gin-incubator API
//	@version		1.0
//	@description	gin demo framework.
//	@termsOfService	https://www.xiexianbin.cn

//	@contact.name	API Support
//	@contact.url	https://www.xiexianbin.cn
//	@contact.email	me@xiexianbin.cn

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

//	@securitydefinitions.oauth2.application	OAuth2Application
//	@tokenUrl								https://example.com/oauth/token
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

//	@securitydefinitions.oauth2.implicit	OAuth2Implicit
//	@authorizationUrl						https://example.com/oauth/authorize
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

//	@securitydefinitions.oauth2.password	OAuth2Password
//	@tokenUrl								https://example.com/oauth/token
//	@scope.read								Grants read access
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode	OAuth2AccessCode
// @tokenUrl								https://example.com/oauth/token
// @authorizationUrl						https://example.com/oauth/authorize
// @scope.admin							Grants read and write access to administrative information
func main() {
	// 默认路由引擎
	r := api.InitRouter()

	v1.InitV1Router()

	// 启动 HTTP 服务
	//r.Run("0.0.0.0:8000")
	//r.Run(":8000")
	_ = r.Run()
}
