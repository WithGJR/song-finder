package models

import (
	"time"
)

type BaseModel struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	childModel interface{}
}

func (this *BaseModel) Init(childModel interface{}) {
	this.childModel = childModel
}

func First(modelInstances interface{}) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.First(modelInstances)
	db.Close()
	return nil
}

func FindAll(modelInstances interface{}) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Find(modelInstances)
	db.Close()
	return nil
}

func (this *BaseModel) Create() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Create(this.childModel)
	db.Close()
	return nil
}

func (this *BaseModel) Update() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Save(this.childModel)
	db.Close()
	return nil
}

func (this *BaseModel) Delete() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	if err := db.First(this.childModel).Error; err != nil {
		return err
	}
	db.Delete(this.childModel)
	db.Close()
	return nil
}
