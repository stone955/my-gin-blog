package model

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

func GetTags(page int, size int, query interface{}) []Tag {
	var tags []Tag
	db.Where(query).Offset(page).Limit(size).Find(&tags)
	return tags
}

func AddTag(name string, state int, createdBy string) {
	db.Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	})
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
