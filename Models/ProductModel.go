package Models

type Product struct {
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
	Stock       string `json:"stock"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Brand       string `json:"brand"`
}
