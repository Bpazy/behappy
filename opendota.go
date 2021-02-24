package really

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SubscribeFunc() {
	var sps []*SubscribePlayer
	if err := db.Distinct("player_id").Find(&sps).Error; err != nil {
		logrus.Infof("没有订阅的玩家")
		return
	}

	// 新比赛
	var newMatchPlayers []*MatchPlayer
	for _, sp := range sps {
		playerDetailRes, err := client.R().Get(fmt.Sprintf("https://api.opendota.com/api/players/%s/matches?limit=1", sp.PlayerId))
		if err != nil {
			logrus.Printf("从 opendota 获取玩家比赛列表失败: %+v", err)
			continue
		}

		var matchPlayers []*MatchPlayer
		MustJsonUnmarshal(playerDetailRes.Body(), &matchPlayers)

		for _, mp := range matchPlayers {
			mp.PlayerID = sp.PlayerId
			s := map[string]interface{}{
				"match_id":  mp.MatchID,
				"player_id": sp.PlayerId,
			}
			if err := db.Where(s).First(&MatchPlayer{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				newMatchPlayers = append(newMatchPlayers, mp)
				logrus.Printf("探测到新的比赛：%d", mp.MatchID)
				db.Create(mp)
			}
		}
	}

	// 逐个群通知
	for _, mp := range newMatchPlayers {
		// 待通知的订阅群组
		var allSub []*SubscribePlayer
		if err := db.Where("player_id = ?", mp.PlayerID).Find(&allSub).Error; err != nil {
			logrus.Info("没有订阅的玩家")
			return
		}
		for _, sp := range allSub {
			pretty := fmt.Sprintf("英雄: %s\n等级: %s\n\n击杀: %d, 死亡: %d, 助攻: %d", mp.HeroName(), mp.SkillString(), mp.Kills, mp.Deaths, mp.Assists)
			SendGroupMessage(sp.GroupId, fmt.Sprintf("「%s」有新「%s」的比赛了: \n\n%s", sp.Name(), mp.MatchResultString(), pretty))
		}
	}
}

func InitHeros() {
	logrus.Info("初始化英雄数据")
	b := Get(fmt.Sprintf("http://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001?key=%s&language=zh", config.SteamAPI.Key))

	var steamApiResult SteamApiResult
	MustJsonUnmarshal(b, &steamApiResult)

	heros := steamApiResult.Result.Heroes
	for _, hero := range heros {
		db.Create(&hero)
	}
}
