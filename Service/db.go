package Service

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"sepet/Models"
	"sepet/Repositories"
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

func AddProductToCart(cartId string, productInfo Models.CartItem, cartRepo *Repositories.CartRepository) error {

	cart, getErr := cartRepo.GetItem(cartId, "cartId")
	if getErr != nil {
		return getErr
	}

	items := cart.CartItems
	for _, item := range items {
		if item.ProductId == productInfo.ProductId {
			err := errors.New("product is already in cart")
			return err
		}
	}
	cart.CartItems = append(cart.CartItems, productInfo)
	putErr := cartRepo.PutItem(*cart)
	if putErr != nil {
		return putErr
	}

	return nil
}

func RemoveProductFromCart(cartId string, productId string, cartRepo *Repositories.CartRepository) error {

	cart, getErr := cartRepo.GetItem(cartId, "cartId")
	if getErr != nil {
		return getErr
	}
	items := cart.CartItems
	for index, item := range items {
		if item.ProductId == productId {
			items[index] = items[len(items)-1]
			cart.CartItems = items[:len(items)-1]
			break
		}
	}
	putErr := cartRepo.PutItem(*cart)
	if putErr != nil {
		return putErr
	}

	return nil
}

func EmptyCart(cartId string, cartRepo *Repositories.CartRepository) error {
	cart, getErr := cartRepo.GetItem(cartId, "cartId")
	if getErr != nil {
		return getErr
	}
	cart.CartItems = nil
	putErr := cartRepo.PutItem(*cart)
	if putErr != nil {
		return putErr
	}
	return nil
}

func UpdateProduct(cartId string, productInfo Models.CartItem, cartRepo *Repositories.CartRepository) error {
	cart, getErr := cartRepo.GetItem(cartId, "cartId")
	if getErr != nil {
		return getErr
	}

	for index, item := range cart.CartItems {
		if item.ProductId == productInfo.ProductId {
			cart.CartItems[index].Amount = productInfo.Amount
			putErr := cartRepo.PutItem(*cart)
			if putErr != nil {
				return putErr
			}
			return nil
		}
	}
	err := errors.New("product is not in the cart")
	return err
}

func GetCart(userId string, cartRepo *Repositories.CartRepository) (*Models.Cart, error) {
	cart, getErr := cartRepo.GetItem(userId, "userId")
	if getErr != nil {
		return nil, getErr
	}
	return cart, nil
}
