package really

import "log"

func Get(url string) []byte {
	r, err := client.R().Get(url)
	if err != nil {
		log.Printf("发送请求失败: %+v\n", err)
		return []byte{}
	}

	return r.Body()
}

func PostJson(url string, v interface{}) []byte {
	r, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(JsonMarshal(v)).
		Post(url)
	if err != nil {
		log.Printf("发送请求失败: %+v\n", err)
		return []byte{}
	}

	return r.Body()
}

func PostJsonString(url string, v string) []byte {
	r, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(v).
		Post(url)
	if err != nil {
		log.Printf("发送请求失败: %+v\n", err)
		return []byte{}
	}

	return r.Body()
}
