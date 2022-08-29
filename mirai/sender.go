package mirai

import (
	"github.com/Bpazy/behappy/bjson"
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/http"
	"github.com/sirupsen/logrus"
	"gopkg.in/errgo.v2/fmt/errors"
)

type MessageSender struct {
	VerifyKey    string
	MiraiBaseUrl string
	Session      string
}

func NewMessageSender() *MessageSender {
	return &MessageSender{
		VerifyKey:    "1234567890",
		MiraiBaseUrl: "http://localhost:8080",
	}
}

func (ms *MessageSender) GetMiraiVerifyUrl() string {
	return ms.MiraiBaseUrl + "/verify"
}

func (ms *MessageSender) GetMiraiSendGroupMessageUrl() string {
	return ms.MiraiBaseUrl + "/sendGroupMessage"
}

func (ms *MessageSender) GetMiraiBindUrl() string {
	return ms.MiraiBaseUrl + "/bind"
}

func (ms *MessageSender) SendGroupMessage(target int, text string) {
	session, err := ms.Auth()
	if err != nil {
		logrus.Printf("发送群消息获取 session 失败: %+v", err)
		return
	}
	ms.Session = session
	defer ms.Release()

	response := SendMessageResponse{}
	b, err := http.PostJson(ms.GetMiraiSendGroupMessageUrl(), NewSendMessage(ms.Session, target, text))
	if err != nil {
		logrus.Printf("发送群消息失败: %+v", err)
		return
	}
	bjson.MustJsonUnmarshal(b, &response)
	if response.Code != 0 {
		logrus.Printf("发送消息失败: %+v", response)
	}
}

func (ms *MessageSender) Auth() (string, error) {
	rb, err := http.PostJson(ms.GetMiraiVerifyUrl(), map[string]string{
		"verifyKey": ms.VerifyKey,
	})
	if err != nil {
		return "", err
	}
	type AuthResult struct {
		Code    int    `json:"code"`
		Session string `json:"session"`
	}
	var authResult AuthResult
	bjson.MustJsonUnmarshal(rb, &authResult)
	if authResult.Code != 0 {
		return "", errors.Newf("verify 失败，请检查 verifyKey: %s\n", ms.VerifyKey)
	}
	logrus.Info("Mirai auth 成功")

	type VerifyResult struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	response := VerifyResult{}
	b, err := http.PostJson(ms.GetMiraiBindUrl(), map[string]interface{}{
		"sessionKey": authResult.Session,
		"qq":         config.GetConfig().Mirai.BotQQ,
	})
	if err != nil {
		return "", err
	}
	bjson.MustJsonUnmarshal(b, &response)
	if response.Code != 0 {
		return "", errors.Newf("verify session 失败，请检查配置文件中的 BotQQ: %+v\n", response)
	}

	logrus.Info("Mirai session verify 成功")
	return authResult.Session, nil
}
