package Test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sepet/Models"
)

func GetRequest(requestKey string) (string, error) {
	request, err := ioutil.ReadFile("./Data/request.json")
	if err != nil {
		return "", err
	}

	var requestData map[string]interface{}
	err = json.Unmarshal(request, &requestData)

	if err != nil {
		return "", err
	}

	if val, ok := requestData[requestKey]; ok {
		body, err := json.Marshal(val)
		if err != nil {
			return "", err
		}
		return string(body), nil
	} else {
		return "", fmt.Errorf("NO_DATA")
	}
}

func GetResult(resultKey string) (interface{}, error) {
	content, err := ioutil.ReadFile("./Data/result.json")
	if err != nil {
		return nil, err
	}
	var data map[string]*Models.Cart
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	if val, ok := data[resultKey]; ok {
		return val, nil
	} else {
		return "", fmt.Errorf("NO_DATA")
	}
}
