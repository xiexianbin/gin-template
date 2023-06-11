// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
package db

import (
	"sync"

	"gorm.io/gorm"

	"github.com/xiexianbin/gin-incubator/db/mysql"
)

var db *gorm.DB
var dbOnce sync.Once

// DB https://gorm.io/zh_CN/docs/generic_interface.html
type DB interface {
	New() (*gorm.DB, error)
	Ping()
	Close()
	Stats()
}

func Pool() *gorm.DB {
	//数据库连接单例
	dbOnce.Do(func() {
		sql := mysql.MySQL{}
		db, _ = sql.New()
	})

	return db
}
