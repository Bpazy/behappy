package mirai

import (
	"fmt"
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/dao"
)

type AllSubscribeCommand struct {
}

func (s *AllSubscribeCommand) Keyword() string {
	return "查询订阅"
}

func (s *AllSubscribeCommand) Run(event interface{}, _ string) (result string) {
	subscribePlayers := dao.ListSubPlayersByGroupId(event.(*Event).Sender.Group.ID)
	for _, player := range subscribePlayers {
		result = result + fmt.Sprintf("* %s (%s)", player.PlayerID, player.Alias)
	}
	return
}

func (s *AllSubscribeCommand) Example() string {
	return "查询订阅"
}

func init() {
	command.Commanders.Register(&AllSubscribeCommand{})
}
