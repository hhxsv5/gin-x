package helper

import "encoding/json"

func Json2String(d interface{}) (string, error) {
	j, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func String2Json(j string, v interface{}) error {
	return json.Unmarshal([]byte(j), v)
}
