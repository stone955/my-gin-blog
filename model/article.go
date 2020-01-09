package model

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Article struct {
	gorm.Model
	TagId uint
	Tag   Tag

	Title   string
	Desc    string
	Content string
	State   int

	CreatedBy string
	UpdatedBy string
	DeletedBy string
}

func (a *Article) BeforeSave(scope *gorm.Scope) {
	log.Println("Call BeforeSave")
}

func (a *Article) BeforeCreate(scope *gorm.Scope) {
	log.Println("Call BeforeCreate")
}

func (a *Article) BeforeUpdate(scope *gorm.Scope) {
	log.Println("Call BeforeUpdate")
}

func (a *Article) BeforeDelete(scope *gorm.Scope) {
	log.Println("Call BeforeDelete")
}

func (a *Article) AfterSave(scope *gorm.Scope) {
	log.Println("Call AfterSave")
}

func (a *Article) AfterCreate(scope *gorm.Scope) {
	log.Println("Call AfterCreate")
}

func (a *Article) AfterUpdate(scope *gorm.Scope) {
	log.Println("Call AfterUpdate")
}

func (a *Article) AfterDelete(scope *gorm.Scope) {
	log.Println("Call AfterDelete")
}

func (a *Article) AfterFind(scope *gorm.Scope) {
	log.Println("Call AfterFind")
}
