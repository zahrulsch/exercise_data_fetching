package database

import "backend/model"

func (db *Database) FetchSingleProduct(id int64) (*model.Product, error) {
	var result model.Product

	if tx := db.db.Where("id = ?", id).First(&result); tx.Error != nil {
		return nil, tx.Error
	}

	return &result, nil
}
