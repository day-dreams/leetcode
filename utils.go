package leetcode

import (
	"encoding/json"
)

func Json(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}
