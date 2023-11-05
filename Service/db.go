package Service

import (
	"errors"
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

//TODO: change product info parameter to productId
func AddProductToCart(cartId string, productInfo Models.CartItem) error {
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
		TableName: aws.String(os.Getenv("CART_TABLE")),
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
	for _, item := range items {
		if item.ProductId == productInfo.ProductId {
			err = errors.New("product is already in cart")
			return err
		}
	}
	cart.CartItems = append(cart.CartItems, productInfo)
	attrVal, attrErr := dynamodbattribute.MarshalMap(cart)
	if attrErr != nil {
		return attrErr
	}
	putInput := dynamodb.PutItemInput{
		Item:      attrVal,
		TableName: aws.String(os.Getenv("CART_TABLE")),
	}
	_, putErr := conn.PutItem(&putInput)
	if putErr != nil {
		return putErr
	}

	return nil
}

func RemoveProductFromCart(cartId string, productId string) error {
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
		TableName: aws.String(os.Getenv("CART_TABLE")),
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
		TableName: aws.String(os.Getenv("CART_TABLE")),
	}
	_, putErr := conn.PutItem(&putInput)
	if putErr != nil {
		return putErr
	}

	return err
}

func EmptyCart(cartId string) error {
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
		TableName: aws.String(os.Getenv("CART_TABLE")),
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
		TableName: aws.String(os.Getenv("CART_TABLE")),
	}
	_, putErr := conn.PutItem(&putInput)
	if putErr != nil {
		return putErr
	}
	return nil
}

func UpdateProduct(cartId string, productInfo Models.CartItem) error {
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
		TableName: aws.String(os.Getenv("CART_TABLE")),
	}
	result, getErr := conn.GetItem(&getInput)
	if getErr != nil {
		return getErr
	}
	var cart *Models.Cart
	if err := dynamodbattribute.UnmarshalMap(result.Item, &cart); err != nil {
		return err
	}

	for index, item := range cart.CartItems {
		if item.ProductId == productInfo.ProductId {
			cart.CartItems[index].Amount = productInfo.Amount
			attrVal, attrErr := dynamodbattribute.MarshalMap(cart)
			if attrErr != nil {
				return attrErr
			}
			putInput := dynamodb.PutItemInput{
				Item:      attrVal,
				TableName: aws.String(os.Getenv("CART_TABLE")),
			}
			_, putErr := conn.PutItem(&putInput)
			if putErr != nil {
				return putErr
			}
			return nil
		}
	}
	err = errors.New("product is not in the cart")
	return err
}

func GetCart(userId int) error {
	_, err := CreateConnection()
	if err != nil {
		return err
	}
	return err
}
