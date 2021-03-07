package behappy

import (
	"fmt"
	"github.com/Bpazy/behappy/bjson"
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/dao"
	"github.com/Bpazy/behappy/http"
	"github.com/Bpazy/behappy/models"
	"github.com/Bpazy/behappy/opendota"
	"github.com/Bpazy/behappy/qq"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const multiWinMsgTemplate = `嫩吊带，「%s」%s排竟然赢了

比赛ID：%d
比赛等级：%s

%s`

const multiFailMsgTemplate = `惨的，「%s」%s排送分啊

比赛ID：%d
比赛等级：%s

%s`

const singleMatchDescTemplate = `英雄: %s
比赛ID：%d
比赛等级: %s

击杀: %d, 死亡: %d, 助攻: %d`

func SubscribeFunc() {
	playerIDs := dao.ListAllPlayerIDs()
	if len(playerIDs) == 0 {
		return
	}

	// 新比赛
	var newMatchPlayers []*models.MatchPlayer
	for _, pid := range playerIDs {
		matchPlayers := opendota.GetMatchPlayers(pid, 1)
		for _, mp := range matchPlayers {
			mp.PlayerID = pid
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

	groupID2MatchPlayers := map[int][]*models.MatchPlayer{}
	for _, mp := range newMatchPlayers {
		// 待通知的订阅群组
		allSub := dao.ListSubPlayers(mp.PlayerID)
		// 逐个群通知
		for _, sp := range allSub {
			matchPlayers, ok := groupID2MatchPlayers[sp.GroupID]
			if !ok {
				groupID2MatchPlayers[sp.GroupID] = []*models.MatchPlayer{mp}
			} else {
				groupID2MatchPlayers[sp.GroupID] = append(matchPlayers, mp)
			}
		}
	}

	for groupID, matchPlayers := range groupID2MatchPlayers {
		matchID2MatchPlayers := map[int64][]*models.MatchPlayer{}
		for _, matchPlayer := range matchPlayers {
			matchPlayers, ok := matchID2MatchPlayers[matchPlayer.MatchID]
			if !ok {
				matchID2MatchPlayers[matchPlayer.MatchID] = []*models.MatchPlayer{matchPlayer}
			} else {
				matchID2MatchPlayers[matchPlayer.MatchID] = append(matchPlayers, matchPlayer)
			}
		}

		playerID2Name := dao.GetSubPlayerMapByGroupId(groupID)
		for _, matchPlayers := range matchID2MatchPlayers {
			message := ""
			if len(matchPlayers) == 1 {
				mp := matchPlayers[0]

				sp := dao.GetSubPlayer(groupID, mp.PlayerID)

				pretty := fmt.Sprintf(singleMatchDescTemplate, dao.GetHeroName(mp.HeroID), mp.MatchID, mp.SkillString(), mp.Kills, mp.Deaths, mp.Assists)

				if mp.IsWin() {
					message = fmt.Sprintf("嫩吊带，%s竟然赢了 \n\n%s", sp.Name(), pretty)
				} else {
					message = fmt.Sprintf("惨的，%s又输了 \n\n%s", sp.Name(), pretty)
				}
			} else {
				mp := matchPlayers[0]

				pretty := ""
				for _, mp := range matchPlayers {
					sp := dao.GetSubPlayer(groupID, mp.PlayerID)
					pretty += fmt.Sprintf("%s玩%s杀了 %d 个, 死了 %d 次, 助攻 %d 个\n", sp.Name(), dao.GetHeroName(mp.HeroID), mp.Kills, mp.Deaths, mp.Assists)
				}
				pretty = pretty[:len(pretty)-1]

				if mp.IsWin() {
					message = fmt.Sprintf(multiWinMsgTemplate, hanziJoin(matchPlayers, playerID2Name), num2Hanzi(len(matchPlayers)), mp.MatchID, mp.SkillString(), pretty)
				} else {
					message = fmt.Sprintf(multiFailMsgTemplate, hanziJoin(matchPlayers, playerID2Name), num2Hanzi(len(matchPlayers)), mp.MatchID, mp.SkillString(), pretty)
				}
			}
			qq.SendGroupMessage(groupID, message)
		}
	}
}

func hanziJoin(matchPlayers []*models.MatchPlayer, playerID2Name map[string]string) string {
	var names []string
	for _, mp := range matchPlayers[:len(matchPlayers)-1] {
		names = append(names, playerID2Name[mp.PlayerID])
	}
	jm := strings.Join(names, ", ")

	lastMp := matchPlayers[len(matchPlayers)-1:][0]
	lastName := playerID2Name[lastMp.PlayerID]
	joinKey := ""
	if len(matchPlayers) == 2 {
		joinKey = "和"
	} else {
		joinKey = "还有"
	}
	jm += joinKey + lastName
	return jm
}

func num2Hanzi(i int) string {
	switch i {
	case 1:
		return "单"
	case 2:
		return "双"
	case 3:
		return "三"
	case 4:
		return "四"
	case 5:
		return "五"
	}
	return strconv.Itoa(i)
}

func InitHeros() {
	logrus.Info("初始化英雄数据")
	b := http.Get(fmt.Sprintf("http://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001?key=%s&language=zh", config.GetConfig().SteamAPI.Key))

	var steamApiResult models.SteamApiResult
	bjson.MustJsonUnmarshal(b, &steamApiResult)

	heros := steamApiResult.Result.Heroes
	dao.AddHeros(heros)
}

func TellJoke() {
	joke := dao.RandJoke()
	if joke == nil {
		return
	}

	GIDs := dao.ListGroupIDs()
	for _, ID := range GIDs {
		switch joke.Type {
		case models.JokeTypeText:
			qq.SendGroupMessage(ID, "每日刀梗：\n\n"+joke.Content)
		case models.JokeTypeImg:

		default:
			logrus.Errorf("Joke type %s 不存在", joke.Type)
		}
	}
}
