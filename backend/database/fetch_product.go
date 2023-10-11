package database

import "backend/model"

type FetchProductResult struct {
	Count int64            `json:"count"`
	Data  []*model.Product `json:"data"`
}

func (d *Database) FetchProduct(limit, offset int) (*FetchProductResult, error) {
	var products []*model.Product
	var count int64
	var db = d.db

	if tx := db.Model(&products).Count(&count); tx.Error != nil {
		return nil, tx.Error
	}

	result := db.
		Limit(limit).
		Offset(offset).
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &FetchProductResult{
		Data:  products,
		Count: count,
	}, nil
}
