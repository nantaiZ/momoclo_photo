package main

import (
	"api/domain"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func newDB() (db *gorm.DB, err error) {
	args := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_DATABASE"),
		"charset=utf8&loc=Asia%2FTokyo&parseTime=true",
	)
	println(args)
	db, err = gorm.Open("mysql", args)
	if err != nil {
		return
	}

	db.LogMode(true)

	return
}

func migrateAllTables() (err error) {
	db, err := newDB()
	if err != nil {
		return
	}

	db.AutoMigrate(
		&domain.Photo{},
		&domain.Collection{},
		&domain.Ask{},
		&domain.Sell{},
	)

	return
}
