package Service

import (
	"encoding/json"
	"sepet/Models"
)

func UnmarshalRequestBody(requestBody *string) (*Models.CartRequest, error) {
	jsonMap := &Models.CartRequest{}
	if err := json.Unmarshal([]byte(*requestBody), jsonMap); err != nil {
		return nil, err
	}
	return jsonMap, nil
}
