package bjson

import (
	"encoding/json"
	"fmt"
)

func MustJsonMarshal(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Errorf("JSON 序列化错误, 源对象: %+v, 错误: %w", v, err))
	}
	return string(b)
}

func MustJsonUnmarshal(b []byte, v interface{}) {
	err := json.Unmarshal(b, v)
	if err != nil {
		panic(fmt.Errorf("JSON 反序列化错误, 源字符串 %s, 错误: %w", string(b), err))
	}
}
