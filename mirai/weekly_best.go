package mirai

import (
	"fmt"
	"github.com/Bpazy/behappy/dao"
	"github.com/Bpazy/behappy/images"
	"github.com/sirupsen/logrus"
	"time"
)

func GenerateWeeklyBestImage(groupId int) string {
	var playerIds []string
	for _, player := range dao.ListSubPlayersByGroupId(groupId) {
		playerIds = append(playerIds, player.PlayerID)
	}

	max := &dao.PlayerMatchCount{}
	for _, mc := range dao.GetMatchesCount(playerIds) {
		if max == nil || (mc.Count != 0 && max.Count < mc.Count) {
			max = &mc
		}
	}
	if max.Count == 0 {
		return ""
	}

	lastWeek := time.Now().Add(-24 * 7 * time.Hour)
	year, week := lastWeek.ISOWeek()
	player := dao.GetSubPlayer(groupId, max.PlayerID)
	path, err := images.HonorTemplate(player.Name(), year, week, int(lastWeek.Month()), max.Count)
	if err != nil {
		panic(fmt.Errorf("生成图像失败: %+v", err))
	}
	logrus.Infof("生成的图像路径: %s", path)
	return path
}
