package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Connect() error {
	var err error
	Client, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	Client.AutoMigrate(&Node{}, &Tag{})

	return nil
}