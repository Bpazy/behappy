package really

import "encoding/json"

func JsonUnmarshal(b []byte, v interface{}) {
	err := json.Unmarshal(b, v)
	if err != nil {
		panic(err)
	}
}

func JsonMarshal(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}
