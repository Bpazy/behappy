package really

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func SubscribeFunc() {
	var sps []*SubscribePlayer
	if err := db.Distinct("player_id").Find(&sps).Error; err != nil {
		log.Println("没有订阅的玩家")
		return
	}

	// 新比赛
	var newMatchPlayers []*MatchPlayer
	for _, sp := range sps {
		playerDetailRes, err := client.R().Get(fmt.Sprintf("https://api.opendota.com/api/players/%s/matches?limit=1", sp.PlayerId))
		if err != nil {
			log.Printf("从 opendota 获取玩家比赛列表失败: %+v\n", err)
			continue
		}

		var matchPlayers []*MatchPlayer
		JsonUnmarshal(playerDetailRes.Body(), &matchPlayers)

		for _, mp := range matchPlayers {
			mp.PlayerID = sp.PlayerId
			s := map[string]interface{}{
				"match_id":  mp.MatchID,
				"player_id": sp.PlayerId,
			}
			if err := db.Where(s).First(&MatchPlayer{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				newMatchPlayers = append(newMatchPlayers, mp)
				log.Printf("探测到新的比赛：%d\n", mp.MatchID)
				db.Create(mp)
			}
		}
	}

	if err := db.Find(&sps).Error; err != nil {
		log.Println("没有订阅的玩家")
		return
	}
	// 逐个群通知
	for _, mp := range newMatchPlayers {
		for _, sp := range sps {
			pretty := fmt.Sprintf("英雄: %s\n等级: %s\n\n击杀: %d, 死亡: %d, 助攻: %d", mp.HeroName(), mp.SkillString(), mp.Kills, mp.Deaths, mp.Assists)
			SendGroupMessage(sp.GroupId, fmt.Sprintf("「%s」有新「%s」的比赛了: \n\n%s", sp.Name(), mp.MatchResultString(), pretty))
		}
	}
}

func InitHeros() {
	b := Get(fmt.Sprintf("http://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001?key=%s&language=zh", config.SteamAPI.Key))

	var steamApiResult SteamApiResult
	JsonUnmarshal(b, &steamApiResult)

	heros := steamApiResult.Result.Heroes
	for _, hero := range heros {
		db.Create(&hero)
	}
}
