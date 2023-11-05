package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"sepet/Service"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.Resource
	var response events.APIGatewayProxyResponse
	var err error
	switch path {
	case "/add-product":
		fmt.Println("add product")
		response.StatusCode = 200
		response.Body = "Add product"
		requestBody, unmarshallErr := Service.UnmarshalRequestBody(&request.Body)
		if unmarshallErr != nil {
			response.StatusCode = 400
			response.Body = "Request is not in correct format."
			return response, unmarshallErr
		}
		err := Service.AddProductToCart(requestBody.CartId, requestBody.ProductInfo)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		return response, err
	case "/remove-product":
		fmt.Println("remove product")
		response.StatusCode = 200
		response.Body = "Remove product"
		requestBody, unmarshallErr := Service.UnmarshalRequestBody(&request.Body)
		if unmarshallErr != nil {
			response.StatusCode = 400
			response.Body = "Request is not in correct format."
			return response, unmarshallErr
		}
		err := Service.RemoveProductFromCart(requestBody.CartId, requestBody.ProductId)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		return response, err
	case "/empty-cart":
		fmt.Println("remove all")
		response.StatusCode = 200
		response.Body = "Remove All"
		requestBody, unmarshallErr := Service.UnmarshalRequestBody(&request.Body)
		if unmarshallErr != nil {
			response.StatusCode = 400
			response.Body = "Request is not in correct format."
			return response, unmarshallErr
		}
		err = Service.EmptyCart(requestBody.CartId)
		if err != nil {
			response.StatusCode = 400
			return response, err
		}
		return response, err
	case "/update-product":
		fmt.Println("update product")
		response.StatusCode = 200
		response.Body = "Update product"
		requestBody, unmarshallErr := Service.UnmarshalRequestBody(&request.Body)
		if unmarshallErr != nil {
			response.StatusCode = 400
			response.Body = "Request is not in correct format."
			return response, unmarshallErr
		}
		err = Service.UpdateProduct(requestBody.CartId, requestBody.ProductInfo)
		return response, err
	case "/get-cart":
		fmt.Println("get cart data")
		response.Body = "Get Cart Data"
		cart, err := Service.GetCart(request.Headers["Authorization"])
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
		fmt.Println("not valid")
		response.StatusCode = 400
		response.Body = "Not valid"
		return response, err
	}
}

func main() {
	lambda.Start(handler)
}
