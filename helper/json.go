package helper

import "encoding/json"

func Json2String(d interface{}) string {
	j, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func String2Json(j string, v interface{}) {
	err := json.Unmarshal([]byte(j), v)
	if err != nil {
		panic(err)
	}
}
