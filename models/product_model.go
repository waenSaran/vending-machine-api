package models

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	ImageURL string `json:"image_url"`
}

type AddProductResponse struct {
	Status    string `json:"status"`
	Data      string `json:"data"`
	CreatedBy string `json:"created_by"`
}
