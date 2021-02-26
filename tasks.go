package really

import (
	"fmt"
	"github.com/Bpazy/really/bjson"
	"github.com/Bpazy/really/config"
	"github.com/Bpazy/really/dao"
	"github.com/Bpazy/really/http"
	"github.com/Bpazy/really/models"
	"github.com/Bpazy/really/opendota"
	"github.com/sirupsen/logrus"
)

func SubscribeFunc() {
	sps := dao.GetAllSubPlayers()
	if len(sps) == 0 {
		return
	}

	// 新比赛
	var newMatchPlayers []*models.MatchPlayer
	for _, sp := range sps {
		matchPlayers := opendota.GetMatchPlayers(sp.PlayerID, 1)
		for _, mp := range matchPlayers {
			mp.PlayerID = sp.PlayerID
			m := dao.GetMatchPlayer(mp.MatchID, mp.PlayerID)
			if m != nil {
				// 比赛已存在
				continue
			}

			logrus.Printf("探测到新的比赛：%d", mp.MatchID)
			newMatchPlayers = append(newMatchPlayers, &mp)
			dao.SaveMatchPlayer(&mp)
		}
	}

	for _, mp := range newMatchPlayers {
		// 待通知的订阅群组
		allSub := dao.GetSubPlayers(mp.PlayerID)
		// 逐个群通知
		for _, sp := range allSub {
			pretty := fmt.Sprintf("英雄: %s\n等级: %s\n\n击杀: %d, 死亡: %d, 助攻: %d", dao.GetHeroName(mp.HeroID), mp.SkillString(), mp.Kills, mp.Deaths, mp.Assists)
			message := ""
			if mp.IsWin() {
				message = fmt.Sprintf("嫩吊带，「%s」竟然「赢」了 \n\n%s", sp.Name(), pretty)
			} else {
				message = fmt.Sprintf("「%s」又「输」了 \n\n%s", sp.Name(), pretty)
			}
			SendGroupMessage(sp.GroupID, message)
		}
	}
}

func InitHeros() {
	logrus.Info("初始化英雄数据")
	b := http.Get(fmt.Sprintf("http://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001?key=%s&language=zh", config.GetConfig().SteamAPI.Key))

	var steamApiResult models.SteamApiResult
	bjson.MustJsonUnmarshal(b, &steamApiResult)

	heros := steamApiResult.Result.Heroes
	dao.AddHeros(heros)
}
