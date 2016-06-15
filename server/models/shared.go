package models

import (
	"github.com/jinzhu/gorm"
	"song-finder/server/constants"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", constants.DBUsername+":"+constants.DBPassword+"@/"+constants.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}
