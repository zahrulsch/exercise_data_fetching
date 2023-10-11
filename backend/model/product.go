package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID                 int
	Url                string
	ImageUrl           string
	ImageUrlLarge      string
	CatId              int
	GaKey              string
	CountReview        int
	DiscountPercentage int
	Preorder           bool
	Name               string
	Price              string
	PriceInt           int
	Original_price     string
	Rating             float64
}
