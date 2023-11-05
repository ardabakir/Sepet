package Models

type CartRequest struct {
	CartId      string   `json:"cartId"`
	ProductId   string   `json:"productId"`
	ProductInfo CartItem `json:"productInfo"`
}
