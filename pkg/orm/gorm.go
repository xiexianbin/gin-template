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

// Package orm implements the app GORM object by gorm.io/gorm.
package orm

import (
	"strings"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gorml "gorm.io/gorm/logger"

	"github.com/xiexianbin/gin-template/pkg/log"
	"github.com/xiexianbin/gin-template/pkg/util"
)

var (
	once sync.Once
	DB   *gorm.DB
)

func Init(dsn string) {
	if dsn == "" {
		dsn = "db.sqlite"
	}
	once.Do(func() {
		var err error

		gormLogger := &GormLogger{
			ZapLogger: log.Logger,
			Config: gorml.Config{
				SlowThreshold:             200 * time.Microsecond,
				IgnoreRecordNotFoundError: true,
				LogLevel:                  gorml.Silent,
			},
		}

		if strings.HasPrefix(dsn, "mysql://") {
			mysqlDSN := strings.Replace(dsn, "mysql://", "", 1)
			DB, err = gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{
				Logger: gormLogger,
			})
			util.Mustf(err, "init MYSQL failed")
		} else if strings.HasSuffix(dsn, "sqlite") || strings.HasSuffix(dsn, "sqlite3") {
			DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: gormLogger})
			util.Mustf(err, "init SQLITE failed")
		}

		// Connection Pool
		// ref https://gorm.io/docs/connecting_to_the_database.html#Connection-Pool
		sqlDB, err := DB.DB()
		util.Mustf(err, "get sql.DB failed")

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
}

// AutoMigrate run auto migration for given models
func AutoMigrate(dst ...interface{}) error {
	// Auto Migrate
	return DB.AutoMigrate(dst)
}

// Pool return *gorm.DB from Connection Pool
func Pool() *gorm.DB {
	return DB
}
