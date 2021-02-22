package really

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/errgo.v2/fmt/errors"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var subscribeR = regexp.MustCompile("订阅\\s+(\\d+)\\s+别名\\s+(.+)")

// serveMirai 开启 Mirai事件上报监听器
func serveMirai() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/post", func(c *gin.Context) {
		all, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		log.Printf("接受到来自 Mirai 的上报: %s\n", string(all))

		e := NewEvent(all)
		log.Printf("接受到来自 Mirai 的事件：%s\n", e)

		if e.IsGroupMessage() {
			submatch := subscribeR.FindStringSubmatch(e.MessageChain.PlainText())
			if len(submatch) != 3 {
				return
			}
			sp := SubscribePlayer{
				GroupId:  e.Sender.Group.ID,
				PlayerId: submatch[1],
				Alias:    submatch[2],
			}
			db.Create(&sp)
			SendGroupMessage(e.Sender.Group.ID, "订阅成功")
		}

		c.JSON(200, nil)
	})

	Auth()

	log.Fatal(r.Run("0.0.0.0:10000"))
}

func (cs MessageChains) PlainText() string {
	var ret string
	for _, c := range cs {
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
			case "At":
				ret += strconv.Itoa(c.Target)
			}
		}
	}
	return ret
}

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
	return e.Type == "GroupMessage"
}

func (e Event) isEvent() bool {
	return strings.Contains(e.Type, "Event")
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

type SendMessage struct {
	SessionKey   string        `json:"sessionKey"`
	Target       int           `json:"target"`
	MessageChain MessageChains `json:"messageChain"`
}

func NewSendMessage(target int, text string) *SendMessage {
	return &SendMessage{
		SessionKey:   session,
		Target:       target,
		MessageChain: MessageChains{MessageChain{Type: "Plain", Text: text}},
	}
}

type SendMessageResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	MessageID int    `json:"messageId"`
}

func SendGroupMessage(target int, text string) {
	response := SendMessageResponse{}
	b, err := PostJson("http://localhost:8080/sendGroupMessage", NewSendMessage(target, text))
	if err != nil {
		log.Printf("发送群消息失败: %+v\n", err)
		return
	}
	JsonUnmarshal(b, &response)
	if response.Code != 0 {
		log.Printf("发送消息失败: %+v", response)
	}
}

var session string

func Auth() {
	rb, err := PostJson("http://localhost:8080/auth", map[string]string{
		"authKey": "1234567890",
	})
	if err != nil {
		log.Fatalf("Mirai Auth 失败: %+v\n", err)
	}
	type AuthResult struct {
		Code    int    `json:"code"`
		Session string `json:"session"`
	}
	var authResult AuthResult
	JsonUnmarshal(rb, &authResult)
	if authResult.Code != 0 {
		log.Fatalf("auth 失败，请检查 authKey\n")
		return
	}
	log.Println("Mirai auth 成功")

	type VerifyResult struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	response := VerifyResult{}
	b, err := PostJson("http://localhost:8080/verify", map[string]interface{}{
		"sessionKey": authResult.Session,
		"qq":         config.Mirai.BotQQ,
	})
	if err != nil {
		log.Fatalf("verify session 失败，请检查配置文件中的 BotQQ: %+v\n", err)
	}
	JsonUnmarshal(b, &response)
	if response.Code != 0 {
		log.Fatalf("verify session 失败，请检查配置文件中的 BotQQ: %+v\n", response)
	}

	log.Println("Mirai session verify 成功")
	session = authResult.Session
}
