package really

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

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
		c.JSON(200, nil)
	})
	log.Fatal(r.Run("0.0.0.0:10000"))
}

type FriendMessage struct {
	SessionKey   string         `json:"sessionKey"`
	Target       int            `json:"target"`
	MessageChain []PlainMessage `json:"messageChain"`
}

type PlainMessage struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

func NewPlainMessage(text string) PlainMessage {
	return PlainMessage{
		Type: "Plain",
		Text: text,
	}
}

func NewFriendMessage(target int, text string) *FriendMessage {
	return &FriendMessage{
		SessionKey: "",
		Target:     target,
		MessageChain: []PlainMessage{
			NewPlainMessage(text),
		},
	}
}
