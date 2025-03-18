package user

import (
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (u *User) Register(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.Create(u)
}
