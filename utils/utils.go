package utils

import (
	"encoding/json"
)

// ParseData ...
func ParseData[T any](data []byte) (*T, error) {
	var t T
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
