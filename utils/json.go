package utils

import (
	"encoding/json"
)

func MustJson(b interface{}) string {
	body, _ := json.Marshal(b)
	return string(body)
}
