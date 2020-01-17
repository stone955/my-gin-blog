package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"log"
)

var db *gorm.DB

func Setup() {
	var err error
	if db, err = gorm.Open(setting.DatabaseCfg.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseCfg.User, setting.DatabaseCfg.Password, setting.DatabaseCfg.Host, setting.DatabaseCfg.Name)); err != nil {
		log.Fatalf("models.Setup err: %v\n", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseCfg.TablePrefix + defaultTableName
	}

	// 自动建表
	db.AutoMigrate(&Tag{}, &Article{}, &Auth{})
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
}

func CloseDB() {
	defer db.Close()
}
