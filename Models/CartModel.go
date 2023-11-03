package Models

type Cart struct {
	CartId    int        `json:"cartId"`
	UserId    int        `json:"userId"`
	CartItems []CartItem `json:"cartItems"`
}

type CartItem struct {
	CartId    int `json:"cartId"`
	ProductId int `json:"productId"`
	Amount    int `json:"amount"`
}
