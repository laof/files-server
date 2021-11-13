package libs

import "encoding/json"

func JsonData(v interface{}) []byte {
	res, err := json.Marshal(v)

	if err != nil {
		return []byte("")
	}

	return res
}
