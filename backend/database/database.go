package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() (*Database, error) {
	driver := sqlite.Open("database.db")
	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (db *Database) GetDB() *gorm.DB {
	return db.db
}
