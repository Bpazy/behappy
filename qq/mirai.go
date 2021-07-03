package qq

import (
	"encoding/json"
	"fmt"
	"github.com/Bpazy/behappy/bjson"
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/dao"
	"github.com/Bpazy/behappy/http"
	"github.com/Bpazy/behappy/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/errgo.v2/fmt/errors"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var subscribeR = regexp.MustCompile("订阅\\s+(\\d+)\\s+别名\\s+(.+)")

// ServeMirai 开启 Mirai事件上报监听器
func ServeMirai() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	logrus.SetLevel(logrus.DebugLevel)
	r.POST("/post", func(c *gin.Context) {
		all, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		logrus.Debugf("接受到来自 Mirai 的上报: %s", string(all))

		e := NewEvent(all)
		logrus.Debugf("接受到来自 Mirai 的事件：%s", e)

		if e.IsGroupMessage() {
			submatch := subscribeR.FindStringSubmatch(e.MessageChain.PlainText())
			if len(submatch) != 3 {
				return
			}

			savedSP := dao.GetSubPlayer(e.Sender.Group.ID, submatch[1])
			if savedSP == nil {
				// 不存在
				dao.SaveSubPlayer(&models.SubscribePlayer{
					GroupID:  e.Sender.Group.ID,
					PlayerID: submatch[1],
					Alias:    submatch[2],
				})
				SendGroupMessage(e.Sender.Group.ID, "订阅成功")
			} else {
				// 存在则更新
				oldAlias := savedSP.Alias
				savedSP.Alias = submatch[2]
				dao.UpdateSubPlayer(savedSP)
				msg := fmt.Sprintf("%s 更新了订阅: [%s] 被变更为 [%s]", e.Sender.MemberName, oldAlias, savedSP.Alias)
				SendGroupMessage(e.Sender.Group.ID, msg)
			}

		}

		c.JSON(200, nil)
	})

	logrus.Infof("开始监听 %s, 等待 Mirai 事件上报", config.Addr)
	logrus.Fatal(r.Run(config.Addr))
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

func NewSendMessage(session string, target int, text string) *SendMessage {
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
	session, err := Auth()
	if err != nil {
		logrus.Printf("发送群消息获取 session 失败: %+v", err)
	}
	defer Release(session)

	response := SendMessageResponse{}
	b, err := http.PostJson("http://localhost:8080/sendGroupMessage", NewSendMessage(session, target, text))
	if err != nil {
		logrus.Printf("发送群消息失败: %+v", err)
		return
	}
	bjson.MustJsonUnmarshal(b, &response)
	if response.Code != 0 {
		logrus.Printf("发送消息失败: %+v", response)
	}
}

func UploadGroupImage() {
	session, err := Auth()
	if err != nil {
		logrus.Printf("发送群消息获取 session 失败: %+v", err)
	}
	defer Release(session)

	file, err := os.OpenFile("C:/Users/hanzi/Desktop/20210305234729.jpg", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	post, err := http.Client.R().
		SetFileReader("img", "test", file).
		SetFormData(map[string]string{
			"sessionKey": session,
			"type":       "group",
		}).
		Post("http://localhost:8080/uploadImage")
	if err != nil {
		panic(err)
	}

	body := post.String()
	logrus.Printf(body)
}

func Auth() (string, error) {
	authKey := "1234567890"
	rb, err := http.PostJson("http://localhost:8080/auth", map[string]string{
		"authKey": authKey,
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
		return "", errors.Newf("auth 失败，请检查 authKey: %s\n", authKey)
	}
	logrus.Info("Mirai auth 成功")

	type VerifyResult struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	response := VerifyResult{}
	b, err := http.PostJson("http://localhost:8080/verify", map[string]interface{}{
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

func Release(session string) {
	rb, err := http.PostJson("http://localhost:8080/release", map[string]string{
		"sessionKey": session,
		"qq":         config.GetConfig().Mirai.BotQQ,
	})
	if err != nil {
		logrus.Printf("释放 session 失败: %+v", err)
		return
	}
	logrus.Printf("释放 session 成功: %s", string(rb))
}
