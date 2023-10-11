package main

import (
	"encoding/json"
	"log"
	"os"

	"backend/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	for _, product := range products {
		err := db.Create(product).Error

		if err != nil {
			if err == gorm.ErrDuplicatedKey {
				log.Println("data already exist")
			}
		}

	}

}

func databaseSetup() (*gorm.DB, error) {
	database := sqlite.Open("database.db")
	db, err := gorm.Open(database, &gorm.Config{
		Logger: logger.Discard,
	})

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
