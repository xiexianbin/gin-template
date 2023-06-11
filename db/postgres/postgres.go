package postgres

import (
	"gorm.io/gorm"
)

type PostgreSQL struct{}

func (m *PostgreSQL) New() (*gorm.DB, error) {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// return db, err
	return nil, nil
}

func (m *PostgreSQL) Ping() {

}

func (m *PostgreSQL) Close() {

}

func (m *PostgreSQL) Stats() {

}
