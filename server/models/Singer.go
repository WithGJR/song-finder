package models

import (
	"github.com/jinzhu/gorm"
)

type Singer struct {
	gorm.Model
	Name    string
	Country string
}
