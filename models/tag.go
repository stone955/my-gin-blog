package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Tag struct {
	gorm.Model
	Name      string
	CreatedBy string
	UpdatedBy string
	DeletedBy string
	State     int
}

func GetTags(page int, size int, query interface{}) ([]Tag, error) {
	var tags []Tag
	if err := db.Where(query).Offset(page).Limit(size).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func AddTag(name string, state int, createdBy string) (Tag, error) {
	tag := Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	}
	if err := db.Create(&tag).Error; err != nil {
		return Tag{}, err
	}
	return tag, nil
}

func (t *Tag) BeforeSave(scope *gorm.Scope) {
	log.Println("Call BeforeSave")
}

func (t *Tag) BeforeCreate(scope *gorm.Scope) {
	log.Println("Call BeforeCreate")
}

func (t *Tag) BeforeUpdate(scope *gorm.Scope) {
	log.Println("Call BeforeUpdate")
}

func (t *Tag) BeforeDelete(scope *gorm.Scope) {
	log.Println("Call BeforeDelete")
}

func (t *Tag) AfterSave(scope *gorm.Scope) {
	log.Println("Call AfterSave")
}

func (t *Tag) AfterCreate(scope *gorm.Scope) {
	log.Println("Call AfterCreate")
}

func (t *Tag) AfterUpdate(scope *gorm.Scope) {
	log.Println("Call AfterUpdate")
}

func (t *Tag) AfterDelete(scope *gorm.Scope) {
	log.Println("Call AfterDelete")
}

func (t *Tag) AfterFind(scope *gorm.Scope) {
	log.Println("Call AfterFind")
}
