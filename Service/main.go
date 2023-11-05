package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.Path
	fmt.Println(path)
	fmt.Println(request.Resource)
	var response events.APIGatewayProxyResponse
	var err error
	switch path {
	case "/add-product":
		fmt.Println("add product")
	case "/remove-product":
		fmt.Println("remove product")
	case "/empty-cart":
		fmt.Println("remove all")
	case "/update-product":
		fmt.Println("update product")
	case "/get-cart":
		fmt.Println("get cart data")
	default:
		fmt.Println("not valid")
	}
	return response, err
}

func main() {
	lambda.Start(handler)
}
