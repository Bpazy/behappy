package command

import (
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/mirai"
	"github.com/Bpazy/behappy/util/btime"
)

type LastWeekMost struct {
	command.DefaultCommander
}

func (s *LastWeekMost) Keyword() string {
	return "上周劳模"
}

func (s *LastWeekMost) Run(event interface{}, _ string) (msgType command.MsgType, result string) {
	groupId := event.(*mirai.Event).Sender.Group.ID
	start, end := btime.GetLastWeekRange()
	return command.TypeImage, mirai.GenerateWeeklyBestImage(groupId, start, end)
}

func (s *LastWeekMost) Example() string {
	return "上周劳模"
}

func init() {
	command.Commanders.Register(&LastWeekMost{})
}
