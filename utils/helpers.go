package utils

import "encoding/json"

func ConvertResultToDto(data interface{}, obj interface{}) error {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonStr, &obj)
	if err != nil {
		return err
	}

	return nil
}
