package Database

import (
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
	getInput := dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			keyName: {
				S: aws.String(key),
			},
		},
		TableName: aws.String(dyn.table),
	}
	result, getErr := dyn.conn.GetItem(&getInput)
	if getErr != nil {
		return getErr
	}
	if err := dynamodbattribute.UnmarshalMap(result.Item, &castTo); err != nil {
		return err
	}
	return nil
}

func (dyn *DynamoDB) PutItem(item interface{}) error {
	attrVal, attrErr := dynamodbattribute.MarshalMap(item)
	if attrErr != nil {
		return attrErr
	}
	putInput := dynamodb.PutItemInput{
		Item:      attrVal,
		TableName: aws.String(dyn.table),
	}
	_, putErr := dyn.conn.PutItem(&putInput)
	if putErr != nil {
		return putErr
	}
	return nil
}
