package opendota

import (
	"encoding/json"
	"fmt"
	"github.com/Bpazy/behappy/http"
	"github.com/Bpazy/behappy/models"
	"github.com/sirupsen/logrus"
)

func GetMatchPlayers(playerID string) []models.MatchPlayer {
	var mps []models.MatchPlayer
	r, err := http.Client.R().Get(fmt.Sprintf("https://api.opendota.com/api/players/%s/recentMatches", playerID))
	if err != nil {
		logrus.Printf("从 opendota 获取玩家比赛列表失败: %+v", err)
		return mps
	}

	body := r.Body()
	err = json.Unmarshal(body, &mps)
	if err != nil {
		logrus.Printf("从 opendota 获取玩家比赛列表失败, 源字符串 %s, 错误: %+v", body, err)
		return mps
	}
	if len(mps) == 0 {
		return mps
	}
	return mps[0:1]
}
