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

package config

import (
	"github.com/xiexianbin/gin-template/pkg/env"
)

var (
	APP_ENV  = env.GetStr("APP_ENV", "dev")
	APP_PORT = env.GetInt("APP_PORT", 8888)

	// log level, support debug, info, warn, error, dpanic, panic, fatal (default is debug)
	LOG_LEVEL = env.GetStr("LOG_LEVEL", "debug")

	// MYSQL connection dsn
	// gorm connection ref https://gorm.io/docs/connecting_to_the_database.html#MySQL
	// Data source name
	//   ref https://en.wikipedia.org/wiki/Data_source_name
	//   format `mysql://user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local`
	MSYQL_DSN = env.GetStr("MSYQL_DSN", "db.sqlite")

	// Redis connection dsn
	// ref https://pkg.go.dev/github.com/go-redis/redis/v8#ParseURL
	// format redis://<user>:<password>@<host>:<port>/<db_number>
	REDIS_DSN = env.GetStr("REDIS_DSN", "")

	// Redis Sentinel Cluster connection dsn, like REDIS_CLUTER_DSN
	// ref https://pkg.go.dev/github.com/redis/go-redis/v9#FailoverOptions
	REDIS_SENTINEL_DSN = env.GetStr("REDIS_SENTINEL_DSN", "")

	// Redis Cluster connection dns
	// ref https://pkg.go.dev/github.com/redis/go-redis/v9#ParseClusterURL
	// format:
	// redis://<user>:<password>@<host>:<port>?addr=<host2>:<port2>&addr=<host3>:<port3>
	// or
	// rediss://<user>:<password>@<host>:<port>?addr=<host2>:<port2>&addr=<host3>:<port3>
	REDIS_CLUTER_DSN = env.GetStr("REDIS_CLUTER_DSN", "")
)
