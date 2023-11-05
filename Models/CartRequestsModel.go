package Models

type CartRequest struct {
	CartId      int      `json:"cartId"`
	ProductId   int      `json:"productId"`
	ProductInfo CartItem `json:"productInfo"`
}
