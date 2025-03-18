package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataBase struct {
	UserNameDB string
	PasswordDB string
	HostDB     string
	PortDB     string
}

func (d *DataBase) Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
