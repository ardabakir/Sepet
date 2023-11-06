package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"sepet/Database"
	"sepet/Repositories"
	"sepet/Service"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.Resource
	var response events.APIGatewayProxyResponse
	var err error
	conn, connErr := Service.CreateConnection()
	cartDb := Database.NewDynamoDB(os.Getenv("CART_TABLE"), conn)
	cartRepository := Repositories.NewCartRepository(cartDb)
	if connErr != nil {
		response.StatusCode = 500
		response.Body = "Couldn't establish connection with DynamoDB"
		return response, connErr
	}
	switch path {
	case "/add-product":
		requestBody, unmarshallErr := Service.UnmarshalRequestBody(&request.Body)
		if unmarshallErr != nil {
			response.StatusCode = 400
			response.Body = "Request is not in correct format."
			return response, unmarshallErr
		}
		err := Service.AddProductToCart(requestBody.CartId, requestBody.ProductInfo, cartRepository)
		if err != nil {
			response.StatusCode = 400
			response.Body = err.Error()
			return response, err
		}
		response.StatusCode = 200
		response.Body = "Successfully added product"
		return response, err
	case "/remove-product":
		requestBody, unmarshallErr := Service.UnmarshalRequestBody(&request.Body)
		if unmarshallErr != nil {
			response.StatusCode = 400
			response.Body = "Request is not in correct format."
			return response, unmarshallErr
		}
		err := Service.RemoveProductFromCart(requestBody.CartId, requestBody.ProductId, cartRepository)
		if err != nil {
			response.StatusCode = 400
			response.Body = err.Error()
			return response, err
		}
		response.StatusCode = 204
		response.Body = "Successfully removed product"
		return response, err
	case "/empty-cart":
		requestBody, unmarshallErr := Service.UnmarshalRequestBody(&request.Body)
		if unmarshallErr != nil {
			response.StatusCode = 400
			response.Body = "Request is not in correct format."
			return response, unmarshallErr
		}
		err = Service.EmptyCart(requestBody.CartId, cartRepository)
		if err != nil {
			response.StatusCode = 400
			response.Body = err.Error()
			return response, err
		}
		response.StatusCode = 200
		response.Body = "Successfully removed all products from the cart"
		return response, err
	case "/update-product":
		requestBody, unmarshallErr := Service.UnmarshalRequestBody(&request.Body)
		if unmarshallErr != nil {
			response.StatusCode = 400
			response.Body = "Request is not in correct format."
			return response, unmarshallErr
		}
		err = Service.UpdateProduct(requestBody.CartId, requestBody.ProductInfo, cartRepository)
		if err != nil {
			response.StatusCode = 400
			response.Body = err.Error()
			return response, err
		}
		response.StatusCode = 200
		response.Body = "Successfully updated product"
		return response, err
	case "/get-cart":
		cart, err := Service.GetCart(request.Headers["Authorization"], cartRepository)
		if err != nil {
			response.StatusCode = 400
			response.Body = err.Error()
			return response, err
		}
		cartString, marshalErr := json.Marshal(cart)
		if marshalErr != nil {
			response.StatusCode = 500
			response.Body = "Couldn't marshal cart to json"
			return response, marshalErr
		}
		response.StatusCode = 200
		response.Body = string(cartString)
		return response, err
	default:
		response.StatusCode = 400
		response.Body = "Not valid"
		return response, err
	}
}

func main() {
	lambda.Start(handler)
}
