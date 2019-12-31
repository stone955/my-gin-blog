package model

import "github.com/jinzhu/gorm"

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
