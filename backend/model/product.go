package model

type Product struct {
	ID                 int     `json:"id"`
	Url                string  `json:"url"`
	ImageUrl           string  `json:"imageUrl"`
	ImageUrlLarge      string  `json:"imageUrlLarge"`
	CatId              int     `json:"catId"`
	GaKey              string  `json:"gaKey"`
	CountReview        int     `json:"countReview"`
	DiscountPercentage int     `json:"discountPercentage"`
	Preorder           bool    `json:"preorder"`
	Name               string  `json:"name"`
	Price              string  `json:"price"`
	PriceInt           int     `json:"priceInt"`
	OriginalPrice      string  `json:"original_price"`
	Rating             float64 `json:"rating"`
}
