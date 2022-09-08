package mirai

import (
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/util/berrors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"strings"
)

// ServeMirai 开启 Mirai事件上报监听器
func ServeMirai() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	logrus.SetLevel(logrus.DebugLevel)
	r.Any("/post", receiveMessage())

	logrus.Infof("开始监听 %s, 等待 Mirai 事件上报", config.Addr)
	logrus.Fatal(r.Run(config.Addr))
}

func receiveMessage() func(c *gin.Context) {
	return func(c *gin.Context) {
		all := berrors.Unwrap(io.ReadAll(c.Request.Body))
		logrus.Debugf("接受到来自 Mirai 的上报: %s", string(all))
		e := NewEvent(all)
		logrus.Debugf("接受到来自 Mirai 的事件：%s", e)
		handleMessage(e)
		// 返回空指令给 mirai-http-api
		c.JSON(200, map[string]string{
			"command": "",
		})
	}
}

func handleMessage(e *Event) {
	if e.IsGroupMessage() {
		dispatchGroupMessage(e)
	}
}

func dispatchGroupMessage(e *Event) {
	if !e.MessageChain.IsAtMe() {
		return
	}
	content := strings.TrimSpace(e.MessageChain.PlainText())
	anyMatch := false
	for _, cmd := range command.Commanders.RegisteredCommanders {
		prefix := cmd.Keyword()
		if !strings.HasPrefix(content, prefix) {
			continue
		}
		anyMatch = true
		mt, ret := cmd.Run(e, strings.TrimLeft(content, prefix+"/"))
		if mt == command.TypeText && ret != "" {
			NewMessageSender().SendGroupMessage(e.Sender.Group.ID, ret)
		}
		if mt == command.TypeImage {
			if ret != "" {
				NewMessageSender().SendGroupImageMessage(e.Sender.Group.ID, ret)
			} else {
				NewMessageSender().SendGroupMessage(e.Sender.Group.ID, "惨的，这段时间没人刀，DeadGame 坐实了")
			}
		}
	}
	if !anyMatch {
		NewMessageSender().SendGroupMessage(e.Sender.Group.ID, command.Commanders.GetHelpMessage())
	}
}
