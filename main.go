package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"sepet/Models"
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
		jsonMap := &Models.CartRequest{}
		if err := json.Unmarshal([]byte(request.Body), jsonMap); err != nil {
			response.StatusCode = 400
			return response, err
		}
		err := Service.AddProductToCart(jsonMap.CartId, jsonMap.ProductInfo)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		return response, err
	case "/remove-product":
		fmt.Println("remove product")
		response.StatusCode = 200
		response.Body = "Remove product"
		jsonMap := &Models.CartRequest{}
		if err := json.Unmarshal([]byte(request.Body), jsonMap); err != nil {
			response.StatusCode = 400
			return response, err
		}
		err := Service.RemoveProductFromCart(jsonMap.CartId, jsonMap.ProductId)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		return response, err
	case "/empty-cart":
		fmt.Println("remove all")
		response.StatusCode = 200
		response.Body = "Remove All"
		return response, err
	case "/update-product":
		fmt.Println("update product")
		response.StatusCode = 200
		response.Body = "Update product"
		return response, err
	case "/get-cart":
		fmt.Println("get cart data")
		response.StatusCode = 200
		response.Body = "Get Cart Data"
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
