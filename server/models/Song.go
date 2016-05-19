package models

import (
	"github.com/jinzhu/gorm"
)

type Song struct {
	gorm.Model
	Name    string
	Country string
	Lyricist string
	Composer string
    Length string

	SingerId uint
	AlbumId uint
}
