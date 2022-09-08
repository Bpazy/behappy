package http

import (
	"github.com/Bpazy/behappy/util/bjson"
	"github.com/sirupsen/logrus"
)

func Get(url string) []byte {
	r, err := Client.R().Get(url)
	if err != nil {
		logrus.Printf("发送请求失败: %+v", err)
		return []byte{}
	}

	return r.Body()
}

func PostJson(url string, v interface{}) ([]byte, error) {
	r, err := Client.R().
		SetHeader("Accept", "application/json").
		SetBody(bjson.MustJsonMarshal(v)).
		Post(url)
	if err != nil {
		return []byte{}, err
	}

	return r.Body(), nil
}
