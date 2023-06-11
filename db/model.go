package db

import (
	"time"

	"github.com/xiexianbin/gin-incubator/api/common"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"varchar(20);comment:'username'" json:"name"`
	Email     string    `gorm:"comment:'email'" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:'create_at'" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:'update_at'" json:"updated_at"`
}

func (u *User) GetAllUsers(user *User, pagination *common.Pagination) (users []User, err error) {
	// 分页查询
	offset := (pagination.Page - 1) * pagination.Limit
	err = Pool().Where(user).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&users).Error
	return
}
