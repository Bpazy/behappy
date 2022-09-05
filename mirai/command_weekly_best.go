package mirai

import (
	"github.com/Bpazy/behappy/command"
)

type WeeklyBestCommand struct {
	command.DefaultCommander
}

func (s *WeeklyBestCommand) Keyword() string {
	return "最佳劳模"
}

func (s *WeeklyBestCommand) Run(event interface{}, arg string) (msgType command.MsgType, result string) {
	groupId := event.(*Event).Sender.Group.ID
	return command.TypeImage, GenerateWeeklyBestImage(groupId)
}

func (s *WeeklyBestCommand) Example() string {
	return "最佳劳模"
}

func (s *WeeklyBestCommand) Hidden() bool {
	return true
}

func init() {
	command.Commanders.Register(&WeeklyBestCommand{})
}
