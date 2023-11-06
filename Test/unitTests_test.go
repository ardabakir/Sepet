package Test

import (
	"github.com/stretchr/testify/assert"
	"sepet/Models"
	"sepet/Repositories"
	"sepet/Service"
	"testing"
)

func TestCreateConnection(t *testing.T) {
	conn, err := Service.CreateConnection()
	assert.Nil(t, err, "CreateConnection function failed")
	assert.NotNil(t, conn, "Returned connection shouldn't be nil")
}

func TestAddProductToCart(t *testing.T) {
	conn, err := Service.CreateConnection()
	assert.Nil(t, err)

	cartDb := NewDynamoDBMock("Carts", conn)
	cartRepo := Repositories.NewCartRepository(cartDb)

	requestBody, rbErr := GetRequest("addProductRequest")
	assert.Nil(t, rbErr)

	request, unmarshalErr := Service.UnmarshalRequestBody(&requestBody)
	assert.Nil(t, unmarshalErr)

	responseErr := Service.AddProductToCart(request.CartId, request.ProductInfo, cartRepo)
	assert.Nil(t, responseErr)
}

func TestRemoveProductFromCart(t *testing.T) {
	conn, err := Service.CreateConnection()
	assert.Nil(t, err)

	cartDb := NewDynamoDBMock("Carts", conn)
	cartRepo := Repositories.NewCartRepository(cartDb)

	requestBody, rbErr := GetRequest("addProductRequest")
	assert.Nil(t, rbErr)

	request, unmarshalErr := Service.UnmarshalRequestBody(&requestBody)
	assert.Nil(t, unmarshalErr)

	responseErr := Service.RemoveProductFromCart(request.CartId, request.ProductId, cartRepo)
	assert.Nil(t, responseErr)
}

func TestGetCart(t *testing.T) {
	conn, err := Service.CreateConnection()
	assert.Nil(t, err)

	cartDb := NewDynamoDBMock("Carts", conn)
	cartRepo := Repositories.NewCartRepository(cartDb)

	authorizationHeader := "1"
	response, responseErr := Service.GetCart(authorizationHeader, cartRepo)
	assert.Nil(t, responseErr)

	result, resultErr := GetResult("getCartResult")
	assert.Nil(t, resultErr)

	assert.Equal(t, result, response)
}

func TestEmptyCart(t *testing.T) {
	conn, err := Service.CreateConnection()
	assert.Nil(t, err)

	cartDb := NewDynamoDBMock("Carts", conn)
	cartRepo := Repositories.NewCartRepository(cartDb)

	requestBody, rbErr := GetRequest("emptyCartRequest")
	assert.Nil(t, rbErr)

	request, unmarshalErr := Service.UnmarshalRequestBody(&requestBody)
	assert.Nil(t, unmarshalErr)

	responseErr := Service.EmptyCart(request.CartId, cartRepo)
	assert.Nil(t, responseErr)
}

func TestUpdateProduct(t *testing.T) {
	conn, err := Service.CreateConnection()
	assert.Nil(t, err)

	cartDb := NewDynamoDBMock("Carts", conn)
	cartRepo := Repositories.NewCartRepository(cartDb)

	request := Models.CartRequest{
		CartId: "1",
		ProductInfo: Models.CartItem{
			ProductId: "29",
			Amount:    "8",
		},
	}

	responseErr := Service.UpdateProduct(request.CartId, request.ProductInfo, cartRepo)
	assert.Nil(t, responseErr)
}
