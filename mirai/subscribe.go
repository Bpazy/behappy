package mirai

import (
	"fmt"
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/dao"
	"github.com/Bpazy/behappy/models"
	"strings"
)

type SubscribeCommand struct {
}

func (s *SubscribeCommand) Keyword() string {
	return "订阅"
}

func (s *SubscribeCommand) Run(event interface{}, arg string) string {
	split := strings.Split(arg, "/")
	if len(split) < 2 {
		return ""
	}
	steamID := split[0]
	alias := split[1]

	return s.saveOrUpdateSubscribe(event.(*Event), steamID, alias)
}

func (s *SubscribeCommand) saveOrUpdateSubscribe(event *Event, steamID string, alias string) string {
	savedSP := dao.GetSubPlayer(event.Sender.Group.ID, steamID)
	if savedSP == nil {
		// 不存在
		dao.SaveSubPlayer(&models.SubscribePlayer{
			GroupID:  event.Sender.Group.ID,
			PlayerID: steamID,
			Alias:    alias,
		})
		return "订阅成功"
	} else {
		// 存在则更新
		oldAlias := savedSP.Alias
		savedSP.Alias = alias
		dao.UpdateSubPlayer(savedSP)
		return fmt.Sprintf("%s 更新了订阅: [%s] 被变更为 [%s]", event.Sender.MemberName, oldAlias, savedSP.Alias)
	}
}

func (s *SubscribeCommand) Example() string {
	return "订阅/{steamID}/{别名}。比如：订阅/1231312/大傻子"
}

func init() {
	command.Commanders.Register(&SubscribeCommand{})
}
