package mirai

import (
	"fmt"
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/dao"
)

type AllSubscribeCommand struct {
	command.DefaultCommander
}

func (s *AllSubscribeCommand) Keyword() string {
	return "查询订阅"
}

func (s *AllSubscribeCommand) Run(event interface{}, arg string) (msgType command.MsgType, result string) {
	subscribePlayers := dao.ListSubPlayersByGroupId(event.(*Event).Sender.Group.ID)
	if len(subscribePlayers) == 0 {
		return command.TypeText, "还没有订阅"
	}
	for _, player := range subscribePlayers {
		result = result + fmt.Sprintf("* %s (%s)", player.PlayerID, player.Alias)
	}
	return command.TypeText, result
}

func (s *AllSubscribeCommand) Example() string {
	return "查询订阅"
}

func init() {
	command.Commanders.Register(&AllSubscribeCommand{})
}
