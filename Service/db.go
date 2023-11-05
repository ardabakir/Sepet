package Service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"sepet/Models"
)

func CreateConnection() (*dynamodb.DynamoDB, error) {
	region := "eu-central-1"
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

func AddProductToCart(cartId int, productInfo Models.CartItem) error {
	conn, err := CreateConnection()
	if err != nil {
		return err
	}
	getInput := dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"cartId": {
				N: aws.String(string(cartId)),
			},
		},
		TableName: aws.String(os.Getenv("cart-table")),
	}
	result, getErr := conn.GetItem(&getInput)
	if getErr != nil {
		return getErr
	}
	var cart *Models.Cart
	if err := dynamodbattribute.UnmarshalMap(result.Item, &cart); err != nil {
		return err
	}
	cart.CartItems = append(cart.CartItems, productInfo)
	attrVal, attrErr := dynamodbattribute.MarshalMap(cart)
	if attrErr != nil {
		return attrErr
	}
	putInput := dynamodb.PutItemInput{
		Item:      attrVal,
		TableName: aws.String(os.Getenv("cart-table")),
	}
	_, putErr := conn.PutItem(&putInput)
	if putErr != nil {
		return putErr
	}

	return nil
}

func RemoveProductFromCart(cartId int, productId int) error {
	conn, err := CreateConnection()
	if err != nil {
		return err
	}
	getInput := dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"cartId": {
				N: aws.String(string(cartId)),
			},
		},
		TableName: aws.String(os.Getenv("cart-table")),
	}
	result, getErr := conn.GetItem(&getInput)
	if getErr != nil {
		return getErr
	}
	var cart *Models.Cart
	if err := dynamodbattribute.UnmarshalMap(result.Item, &cart); err != nil {
		return err
	}
	items := cart.CartItems
	for index, item := range items {
		if item.ProductId == productId {
			items[index] = items[len(items)-1]
			cart.CartItems = items[:len(items)-1]
			break
		}
	}
	attrVal, attrErr := dynamodbattribute.MarshalMap(cart)
	if attrErr != nil {
		return attrErr
	}
	putInput := dynamodb.PutItemInput{
		Item:      attrVal,
		TableName: aws.String(os.Getenv("cart-table")),
	}
	_, putErr := conn.PutItem(&putInput)
	if putErr != nil {
		return putErr
	}

	return err
}

func EmptyCart(cartId int) error {
	conn, err := CreateConnection()
	if err != nil {
		return err
	}
	getInput := dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"cartId": {
				N: aws.String(string(cartId)),
			},
		},
		TableName: aws.String(os.Getenv("cart-table")),
	}
	result, getErr := conn.GetItem(&getInput)
	if getErr != nil {
		return getErr
	}
	var cart *Models.Cart
	if err := dynamodbattribute.UnmarshalMap(result.Item, &cart); err != nil {
		return err
	}
	cart.CartItems = nil
	attrVal, attrErr := dynamodbattribute.MarshalMap(cart)
	if attrErr != nil {
		return attrErr
	}
	putInput := dynamodb.PutItemInput{
		Item:      attrVal,
		TableName: aws.String(os.Getenv("cart-table")),
	}
	_, putErr := conn.PutItem(&putInput)
	if putErr != nil {
		return putErr
	}
	return nil
}
