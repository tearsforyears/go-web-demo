package util

import (
	"bytes"
	"encoding/json"
)

func ToJSONString(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		panic(err)
	}
	return out.String()
}

func ToMap(jsonStr string) map[string]interface{} {
	var res = map[string]interface{}{}
	json.Unmarshal([]byte(jsonStr), &res)
	return res
}
