package command

import (
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/mirai"
)

type WeeklyBestCommand struct {
	command.DefaultCommander
}

func (s *WeeklyBestCommand) Keyword() string {
	return "最佳劳模"
}

func (s *WeeklyBestCommand) Run(event interface{}, _ string) (msgType command.MsgType, result string) {
	groupId := event.(*mirai.Event).Sender.Group.ID
	return command.TypeImage, mirai.GenerateWeeklyBestImage(groupId)
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
