package infra

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}