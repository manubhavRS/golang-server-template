package utils

import "encoding/json"

func DecodeToJson(data interface{}) ([]byte, error) {
	jsonData, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		return jsonData, jsonErr
	}
	return jsonData, nil
}
