package base

import "encoding/json"

func MustMarshal(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func MustUnMarshal(data []byte, v interface{}) {
	_ = json.Unmarshal(data, v)
	return
}

func StructToMap(v interface{}) map[string]interface{} {
	b, _ := json.Marshal(v)
	result := make(map[string]interface{})

	MustUnMarshal(b, &result)
	return result
}
