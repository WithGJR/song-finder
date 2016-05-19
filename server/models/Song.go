package models

import (
	"github.com/jinzhu/gorm"
)

type Song struct {
	gorm.Model
	Name     string
	Country  string
	Lyricist string
	Composer string
	Length   string

	SingerId uint
	AlbumId  uint
}

func GetAllSong(songs *[]Song) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	db.Find(songs)
	return nil
}

func (this *Song) Create() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	db.Create(this)
	return nil
}
