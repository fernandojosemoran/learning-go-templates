package helpers

import "encoding/json"

func ConvertToStruct[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}
