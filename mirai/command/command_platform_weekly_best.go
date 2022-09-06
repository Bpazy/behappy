package command

import (
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/dao"
	"github.com/Bpazy/behappy/mirai"
)

type PlatformWeeklyBestCommand struct {
	command.DefaultCommander
}

func (s *PlatformWeeklyBestCommand) Keyword() string {
	return "全平台最佳劳模"
}

func (s *PlatformWeeklyBestCommand) Run(_ interface{}, _ string) (msgType command.MsgType, result string) {
	s.JustGo()
	return command.TypeNil, ""
}

func (s *PlatformWeeklyBestCommand) Example() string {
	return "全平台最佳劳模"
}

func (s *PlatformWeeklyBestCommand) Hidden() bool {
	return true
}

func (s *PlatformWeeklyBestCommand) JustGo() {
	groupIds := dao.ListAllGroupIds()
	for _, groupId := range groupIds {
		path := mirai.GenerateWeeklyBestImage(groupId)
		if path == "" {
			continue
		}
		mirai.NewMessageSender().SendGroupImageMessage(groupId, path)
	}
}

func init() {
	command.Commanders.Register(&PlatformWeeklyBestCommand{})
}
