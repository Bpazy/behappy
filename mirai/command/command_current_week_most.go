package command

import (
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/mirai"
	"github.com/Bpazy/behappy/util/btime"
)

type CurrentWeekMost struct {
	command.DefaultCommander
}

func (s *CurrentWeekMost) Keyword() string {
	return "本周劳模"
}

func (s *CurrentWeekMost) Run(event interface{}, _ string) (msgType command.MsgType, result string) {
	groupId := event.(*mirai.Event).Sender.Group.ID
	start, end := btime.GetCurrentWeekRange()
	return command.TypeImage, mirai.GenerateWeeklyBestImage(groupId, start, end)
}

func (s *CurrentWeekMost) Example() string {
	return "本周劳模"
}

func init() {
	command.Commanders.Register(&CurrentWeekMost{})
}
