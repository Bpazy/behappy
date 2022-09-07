package command

import (
	"fmt"
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/dao"
	"github.com/Bpazy/behappy/mirai"
	"strings"
)

type SubscribeCommand struct {
	command.DefaultCommander
}

func (s *SubscribeCommand) Keyword() string {
	return "订阅"
}

func (s *SubscribeCommand) Run(event interface{}, arg string) (command.MsgType, string) {
	split := strings.Split(arg, "/")
	if len(split) < 2 {
		return command.TypeText, ""
	}
	steamID := split[0]
	alias := split[1]

	return command.TypeText, s.saveOrUpdateSubscribe(event.(*mirai.Event), steamID, alias)
}

func (s *SubscribeCommand) saveOrUpdateSubscribe(event *mirai.Event, steamID string, alias string) string {
	subscription := dao.GetSubscription(event.Sender.Group.ID, steamID)
	if subscription == nil {
		// 不存在
		dao.SaveSubPlayer(event.Sender.Group.ID, steamID, alias)
		return "订阅成功"
	} else {
		// 存在则更新
		oldAlias := subscription.Alias
		subscription.Alias = alias
		dao.UpdateSubPlayer(subscription)
		return fmt.Sprintf("%s 更新了订阅: [%s] 被变更为 [%s]", event.Sender.MemberName, oldAlias, subscription.Alias)
	}
}

func (s *SubscribeCommand) Example() string {
	return "订阅/{steamID}/{别名}。比如：订阅/1231312/大傻子"
}

func init() {
	command.Commanders.Register(&SubscribeCommand{})
}
