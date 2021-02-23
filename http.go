package really

import (
	"github.com/sirupsen/logrus"
)

func Get(url string) []byte {
	r, err := client.R().Get(url)
	if err != nil {
		logrus.Printf("发送请求失败: %+v", err)
		return []byte{}
	}

	return r.Body()
}

func PostJson(url string, v interface{}) ([]byte, error) {
	r, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(MustJsonMarshal(v)).
		Post(url)
	if err != nil {
		return []byte{}, err
	}

	return r.Body(), nil
}

func PostJsonString(url string, v string) []byte {
	r, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(v).
		Post(url)
	if err != nil {
		logrus.Printf("发送请求失败: %+v", err)
		return []byte{}
	}

	return r.Body()
}
