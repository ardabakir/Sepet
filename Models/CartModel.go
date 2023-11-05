package Models

type Cart struct {
	CartId    string     `json:"cartId"`
	UserId    string     `json:"userId"`
	CartItems []CartItem `json:"cartItems"`
}

type CartItem struct {
	ProductId string `json:"productId"`
	Amount    string `json:"amount"`
}
