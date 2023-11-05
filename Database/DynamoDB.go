package Database

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDB struct {
	table string
	conn  *dynamodb.DynamoDB
}

func NewDynamoDB(table string, conn *dynamodb.DynamoDB) *DynamoDB {
	return &DynamoDB{
		table: table,
		conn:  conn,
	}
}

func (dyn *DynamoDB) GetItem(key string, keyName string, castTo interface{}) error {
	fmt.Println("before get input")
	getInput := dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			keyName: {
				S: aws.String(key),
			},
		},
		TableName: aws.String(dyn.table),
	}
	fmt.Println("got input")
	result, getErr := dyn.conn.GetItem(&getInput)
	if getErr != nil {
		return getErr
	}
	fmt.Println("got item")
	if err := dynamodbattribute.UnmarshalMap(result.Item, &castTo); err != nil {
		return err
	}
	fmt.Println("unmarshaled map")
	return nil
}

func (dyn *DynamoDB) PutItem(item interface{}) error {
	fmt.Println("marshal item")
	attrVal, attrErr := dynamodbattribute.MarshalMap(item)
	if attrErr != nil {
		return attrErr
	}
	fmt.Println("marshaled item")
	putInput := dynamodb.PutItemInput{
		Item:      attrVal,
		TableName: aws.String(dyn.table),
	}
	fmt.Println("put input")
	_, putErr := dyn.conn.PutItem(&putInput)
	if putErr != nil {
		return putErr
	}
	fmt.Println("put complete")
	return nil
}
