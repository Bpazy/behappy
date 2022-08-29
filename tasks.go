package behappy

import (
	"fmt"
	"github.com/Bpazy/behappy/bjson"
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/dao"
	"github.com/Bpazy/behappy/http"
	"github.com/Bpazy/behappy/mirai"
	"github.com/Bpazy/behappy/models"
	"github.com/Bpazy/behappy/opendota"
	"github.com/Bpazy/behappy/templates"
	"github.com/Bpazy/behappy/util"
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

func InitHeros() {
	logrus.Info("初始化英雄数据")
	b := http.Get(fmt.Sprintf("http://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001?key=%s&language=zh", config.GetConfig().SteamAPI.Key))

	var steamApiResult models.SteamApiResult
	bjson.MustJsonUnmarshal(b, &steamApiResult)

	heros := steamApiResult.Result.Heroes
	dao.AddHeros(heros)
}

func SubscribeFunc() {
	playerIDs := dao.ListAllPlayerIDs()
	logrus.Printf("待检测 playerIDs: %v", playerIDs)
	if len(playerIDs) == 0 {
		return
	}

	for groupID, subNewMatchPlayers := range getNewMatchPlayersByGroupId(detectAndSaveNewMatches(playerIDs)) {
		playerID2Name := dao.GetSubPlayerMapByGroupId(groupID)
		for _, matchPlayers := range getNewMatchPlayersByMatchId(subNewMatchPlayers) {
			mirai.NewMessageSender().SendGroupMessage(groupID, buildMessage(matchPlayers, groupID, playerID2Name))
		}
	}
}

func buildMessage(matchPlayers []*models.MatchPlayer, groupID int, playerID2Name map[string]string) string {
	if len(matchPlayers) == 1 {
		return getSinglePlayerMessage(matchPlayers, groupID)
	} else {
		return getMultiPlayersMessage(matchPlayers, groupID, playerID2Name)
	}
}

func getSinglePlayerMessage(matchPlayers []*models.MatchPlayer, groupID int) string {
	mp := matchPlayers[0]
	sp := dao.GetSubPlayer(groupID, mp.PlayerID)

	winTimes, loseTimes := GetWinOrLoseTimesInRow(sp.PlayerID)
	data := map[string]interface{}{
		"Win":        mp.IsWin(),
		"Name":       sp.Name(),
		"HeroName":   dao.GetHeroName(mp.HeroID),
		"MatchID":    mp.MatchID,
		"MatchLevel": mp.SkillString(),
		"Kills":      mp.Kills,
		"Deaths":     mp.Deaths,
		"Assists":    mp.Assists,
		"Duration":   mp.DurationMinutes(),
		"winTimes":   winTimes,
		"loseTimes":  loseTimes,
	}
	logrus.Infof("获取模板入参: %+v", data)
	m, err := templates.GetSingleMessage(data)
	if err != nil {
		logrus.Errorf("模板错误: %+v", err)
	}
	return m
}

func getMultiPlayersMessage(matchPlayers []*models.MatchPlayer, groupID int, playerID2Name map[string]string) string {
	mp := matchPlayers[0]
	pretty := joinKda(matchPlayers, groupID)
	if mp.IsWin() {
		return fmt.Sprintf(multiWinMsgTemplate, joinChineseWords(matchPlayers, playerID2Name), num2ChineseWord(len(matchPlayers)), mp.MatchID, mp.SkillString(), pretty)
	} else {
		pretty += appendRidicule(matchPlayers, playerID2Name)
		return fmt.Sprintf(multiFailMsgTemplate, joinChineseWords(matchPlayers, playerID2Name), num2ChineseWord(len(matchPlayers)), mp.MatchID, mp.SkillString(), pretty)
	}
}

func joinKda(matchPlayers []*models.MatchPlayer, groupID int) string {
	pretty := ""
	for _, mp := range matchPlayers {
		sp := dao.GetSubPlayer(groupID, mp.PlayerID)
		kda := util.GetKda(mp.Kills, mp.Deaths, mp.Assists)
		pretty += fmt.Sprintf("%s玩%s KDA: %s (%d, %d, %d)\n", sp.Name(), dao.GetHeroName(mp.HeroID), kda, mp.Kills, mp.Deaths, mp.Assists)
	}
	return pretty[:len(pretty)-1]
}

func appendRidicule(matchPlayers []*models.MatchPlayer, playerID2Name map[string]string) string {
	maxLoseMp := new(models.MatchPlayer)
	maxLoseTimes := 0
	for _, mp := range matchPlayers {
		_, loseTimes := GetWinOrLoseTimesInRow(mp.PlayerID)
		if loseTimes >= maxLoseTimes {
			maxLoseTimes = loseTimes
			maxLoseMp = mp
		}
	}
	if maxLoseTimes >= 5 {
		return fmt.Sprintf("\n\n黑哥TV已经带不动这个连跪%d把的菜逼%s了", maxLoseTimes, playerID2Name[maxLoseMp.PlayerID])
	}
	return ""
}

func getNewMatchPlayersByMatchId(subNewMatchPlayers []*models.MatchPlayer) map[int64][]*models.MatchPlayer {
	result := map[int64][]*models.MatchPlayer{}
	for _, mp := range subNewMatchPlayers {
		result[mp.MatchID] = append(result[mp.MatchID], mp)
	}
	return result
}

func getNewMatchPlayersByGroupId(newMatchPlayers []*models.MatchPlayer) map[int][]*models.MatchPlayer {
	result := map[int][]*models.MatchPlayer{}
	for _, mp := range newMatchPlayers {
		// 待通知的订阅群组
		allSub := dao.ListSubPlayers(mp.PlayerID)
		// 逐个群通知
		for _, sp := range allSub {
			result[sp.GroupID] = append(result[sp.GroupID], mp)
		}
	}
	return result
}

func detectAndSaveNewMatches(playerIDs []string) (result []*models.MatchPlayer) {
	for _, pid := range playerIDs {
		matchPlayers := opendota.GetMatchPlayers(pid)
		for _, mp := range matchPlayers {
			mp.PlayerID = pid
			m := dao.GetMatchPlayer(mp.MatchID, mp.PlayerID)
			if m != nil {
				// 比赛已存在
				continue
			}

			logrus.Printf("探测到新的比赛：%d", mp.MatchID)
			result = append(result, &mp)
			dao.SaveMatchPlayer(&mp)
		}
	}
	return result
}

func GetWinOrLoseTimesInRow(playerID string) (winTimes, loseTimes int) {
	matches := dao.ListRecentMatchPlayers(playerID)
	if len(matches) == 0 {
		return 0, 0
	}
	firstWin := matches[0].IsWin()
	for _, match := range matches {
		if firstWin && match.IsWin() {
			winTimes++
		}
		if firstWin && !match.IsWin() {
			break
		}
		if !firstWin && !match.IsWin() {
			loseTimes++
		}
		if !firstWin && match.IsWin() {
			break
		}
	}
	return
}

func joinChineseWords(matchPlayers []*models.MatchPlayer, playerID2Name map[string]string) string {
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

func num2ChineseWord(i int) string {
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
