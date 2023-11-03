package Models

type Product struct {
	ProductId   int     `json:"productId"`
	ProductName string  `json:"productName"`
	Stock       int     `json:"stock"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
}
