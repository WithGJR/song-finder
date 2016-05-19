package models

import (
	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	Name          string
	Company       string
	PublishedDate string
	Language      string
	Reward        string
	Photo         string
}
