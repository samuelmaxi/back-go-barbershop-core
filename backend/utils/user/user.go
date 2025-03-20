package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
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

func (u *User) Register(db *gorm.DB) (int, error) {
	db.AutoMigrate(&User{})
	hash, err := u.hashPassword()
	u.Password = hash
	if err != nil {
		panic(err)
	}

	result := db.Create(&u)
	if result.Error != nil {
		return 0, result.Error
	}

	fmt.Println("\n ")
	fmt.Println("ERRO DO RETORNO DA CRIACAO: ", result.Error)
	fmt.Println("\n ")
	fmt.Println("MENSAGEM DO RETORNO DA CRIACAO: ", result.RowsAffected)
	fmt.Println("\n ")

	fmt.Println("Usuario registrado com sucesso: ", u)
	return int(result.RowsAffected), nil
}
