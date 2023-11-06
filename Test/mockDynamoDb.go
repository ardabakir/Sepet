package Test

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"io/ioutil"
)

type mockDynamoDb struct {
	table string
	conn  *dynamodb.DynamoDB
}

func NewDynamoDBMock(table string, conn *dynamodb.DynamoDB) *mockDynamoDb {
	return &mockDynamoDb{
		table: table,
		conn:  conn,
	}
}

func (db mockDynamoDb) GetItem(key string, keyName string, castTo interface{}) error {
	content, err := ioutil.ReadFile("./Data/Cart.json")
	if err != nil {
		return err
	}
	var data map[string]map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return err
	}
	if val, ok := data[db.table]; ok {
		jsonStr, err := json.Marshal(val["cart1"])
		if err != nil {
			return err
		}
		if err := json.Unmarshal(jsonStr, &castTo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("NO_DATA")
	}
	return nil
}

func (db mockDynamoDb) PutItem(item interface{}) error {
	return nil
}
