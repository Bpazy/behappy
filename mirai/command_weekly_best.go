package mirai

import (
	"github.com/Bpazy/behappy/command"
	"github.com/Bpazy/behappy/images"
	"github.com/sirupsen/logrus"
	"time"
)

type WeeklyBestCommand struct {
	command.DefaultCommander
}

func (s *WeeklyBestCommand) Keyword() string {
	return "最佳劳模"
}

func (s *WeeklyBestCommand) Run(event interface{}, arg string) (msgType command.MsgType, result string) {
	year, week := time.Now().ISOWeek()
	path, err := images.Test("南帅", year, week, 23)
	if err != nil {
		panic(err)
	}
	logrus.Infof("生成的图像路径: %s", path)
	return command.TYPE_IMAGE, path
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
