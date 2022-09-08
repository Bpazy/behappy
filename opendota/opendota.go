package opendota

import (
	"encoding/json"
	"fmt"
	"github.com/Bpazy/behappy/dto"
	"github.com/Bpazy/behappy/http"
	"github.com/sirupsen/logrus"
)

func GetMatchPlayers(playerID string) (result []dto.MatchPlayerDto) {
	r, err := http.Client.R().Get(fmt.Sprintf("https://api.opendota.com/api/players/%s/recentMatches", playerID))
	if err != nil {
		logrus.Printf("从 opendota 获取玩家比赛列表失败: %+v", err)
		return result
	}

	body := r.Body()
	err = json.Unmarshal(body, &result)
	if err != nil {
		logrus.Printf("从 opendota 获取玩家比赛列表失败, 源字符串 %s, 错误: %+v", body, err)
		return result
	}
	if len(result) == 0 {
		return result
	}
	return result[0:1]
}
