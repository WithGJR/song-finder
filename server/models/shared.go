package models

import (
	"github.com/jinzhu/gorm"
)

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "username:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}
