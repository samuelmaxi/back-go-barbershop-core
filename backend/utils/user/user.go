package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (u *User) hashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes), err
}

func (u *User) Register(db *gorm.DB) {
	db.AutoMigrate(&User{})
	hash, err := u.hashPassword()
	u.Password = hash
	if err != nil {
		panic(err)
	}

	db.Create(u)
	fmt.Println("Usuario registrado com sucesso: ", u)
}
