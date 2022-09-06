package mirai

import (
	"encoding/json"
	"fmt"
	"github.com/Bpazy/behappy/bjson"
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/http"
	"github.com/sirupsen/logrus"
	"gopkg.in/errgo.v2/fmt/errors"
	"strconv"
	"strings"
)

// Event Mirai 上报的事件
type Event struct {
	Type         string        `json:"type"`
	MessageChain MessageChains `json:"messageChain"`
	QQ           int           `json:"qq"`
	Sender       struct {
		ID         int    `json:"id"`
		MemberName string `json:"memberName"`
		Permission string `json:"permission"`
		Nickname   string `json:"nickname"`
		Remark     string `json:"remark"`
		Group      struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Permission string `json:"permission"`
		} `json:"group"`
	} `json:"sender"`
}

func (e Event) IsMessage() bool {
	return strings.Contains(e.Type, "Message")
}

func (e Event) IsGroupMessage() bool {
	return e.Type == "GroupMessage"
}

func (e Event) IsFriendMessage() bool {
	return e.Type == "FriendMessage"
}

func (e Event) isEvent() bool {
	return strings.Contains(e.Type, "Event")
}

func (e Event) String() string {
	if e.IsGroupMessage() {
		sender := e.Sender
		group := sender.Group
		return fmt.Sprintf("群「%s(%d)」的「%s(%d)」的消息: %s", group.Name, group.ID, sender.MemberName, sender.ID, e.MessageChain.PlainText())
	}

	if e.IsFriendMessage() {
		return fmt.Sprintf("%s(%d)的消息: %s", e.Sender.Nickname, e.Sender.ID, e.MessageChain.PlainText())
	}

	if e.isEvent() {
		return fmt.Sprintf("%s Event from %d", e.Type, e.QQ)
	}

	return e.MessageChain.PlainText()
}

func NewEvent(b []byte) *Event {
	var e Event
	if err := json.Unmarshal(b, &e); err != nil {
		panic(errors.Newf("反序列化JSON错误: %+v", err))
	}
	return &e
}

type MessageChain struct {
	Type    string      `json:"type"`
	ID      int         `json:"id,omitempty"`
	Time    int         `json:"time,omitempty"`
	Target  int         `json:"target,omitempty"`
	Display string      `json:"display,omitempty"`
	Text    string      `json:"text,omitempty"`
	ImageID string      `json:"imageId,omitempty"`
	URL     string      `json:"url,omitempty"`
	Path    interface{} `json:"path,omitempty"`
	FaceID  int         `json:"faceId,omitempty"`
	Name    string      `json:"name,omitempty"`
}

// MessageChains Mirai 上报事件中的 Message
type MessageChains []MessageChain

func (cs *MessageChains) IsAtMe() bool {
	for _, c := range *cs {
		if c.Type != "At" {
			continue
		}
		t := strconv.Itoa(c.Target)
		logrus.Infof("At target: %s, Mirai.BotQQ: %s", t, config.GetConfig().Mirai.BotQQ)
		return t == config.GetConfig().Mirai.BotQQ
	}
	return false
}

type SendMessage struct {
	SessionKey   string        `json:"sessionKey"`
	Target       int           `json:"target"`
	MessageChain MessageChains `json:"messageChain"`
}

func (cs *MessageChains) PlainText() string {
	var ret string
	for _, c := range *cs {
		if c.Type == "Plain" {
			ret += c.Text
		} else {
			switch c.Type {
			case "Face":
				ret += "「表情」"
			case "Image":
				ret += "「图片」"
			case "AtAll":
				ret += "@全体成员"
			}
		}
	}
	return ret
}

type MessageSender struct {
	VerifyKey    string
	MiraiBaseUrl string
	Session      string
}

type SendMessageResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	MessageID int    `json:"messageId"`
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

const sendGroupMessageErrTemplate = "发送群消息失败: %+v"

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
		logrus.Printf(sendGroupMessageErrTemplate, err)
		return
	}
	bjson.MustJsonUnmarshal(b, &response)
	if response.Code != 0 {
		logrus.Printf(sendGroupMessageErrTemplate, response)
	}
}

func (ms *MessageSender) SendGroupImageMessage(target int, path string) {
	session, err := ms.Auth()
	if err != nil {
		logrus.Printf("发送群消息获取 session 失败: %+v", err)
		return
	}
	ms.Session = session
	defer ms.Release()
	b, err := http.PostJson(ms.GetMiraiSendGroupMessageUrl(), NewSendImageMessage(ms.Session, target, path))
	if err != nil {
		logrus.Printf(sendGroupMessageErrTemplate, err)
		return
	}
	response := SendMessageResponse{}
	bjson.MustJsonUnmarshal(b, &response)
	if response.Code != 0 {
		logrus.Printf("发送群消息失败: %+v", response)
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

func (ms *MessageSender) Release() {
	if ms.Session == "" {
		logrus.Printf("session 不存在")
		return
	}
	rb, err := http.PostJson("http://localhost:8080/release", map[string]string{
		"sessionKey": ms.Session,
		"qq":         config.GetConfig().Mirai.BotQQ,
	})
	if err != nil {
		logrus.Printf("释放 session 失败: %+v", err)
		return
	}
	logrus.Printf("释放 session 成功: %s", string(rb))
}

func NewSendMessage(session string, target int, text string) *SendMessage {
	return &SendMessage{
		SessionKey:   session,
		Target:       target,
		MessageChain: MessageChains{MessageChain{Type: "Plain", Text: text}},
	}
}

func NewSendImageMessage(session string, target int, path string) *SendMessage {
	return &SendMessage{
		SessionKey:   session,
		Target:       target,
		MessageChain: MessageChains{MessageChain{Type: "Image", Path: path}},
	}
}
