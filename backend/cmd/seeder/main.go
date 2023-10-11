package main

import (
	"encoding/json"
	"log"
	"os"

	"backend/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := databaseSetup()
	if err != nil {
		log.Fatal("Failed to setup database", err)
	}

	products, err := loadDataFromJson()
	if err != nil {
		log.Fatal("Failed to load json resource", err)
	}

	if tx := db.Create(products); tx.Error != nil {
		log.Fatal("Seeding process failed", tx.Error)
	}
}

func databaseSetup() (*gorm.DB, error) {
	database := sqlite.Open("database.db")
	db, err := gorm.Open(database, &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Product{})
	return db, err
}

func loadDataFromJson() (products []*model.Product, err error) {
	jsonPath := "./products.json"
	file, err := os.ReadFile(jsonPath)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &products)
	if err != nil {
		return
	}

	return
}
