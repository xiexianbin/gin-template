package sqlite

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type SQLite struct{}

func (m *SQLite) New() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	return db, err
}

func (m *SQLite) Ping() {

}

func (m *SQLite) Close() {

}

func (m *SQLite) Stats() {

}
