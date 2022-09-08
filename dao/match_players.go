package dao

import (
	"context"
	"github.com/Bpazy/behappy/dto"
	"github.com/Bpazy/behappy/ent/subscriptionmatch"
)

func GetMatchPlayer(matchID int64, playerID string) *dto.MatchPlayerDto {
	if matchID == 0 || playerID == "" {
		return nil
	}
	sm := client.SubscriptionMatch.
		Query().
		Where(
			subscriptionmatch.And(
				subscriptionmatch.MatchID(matchID)),
			subscriptionmatch.PlayerID(playerID),
		).
		FirstX(context.TODO())
	if sm == nil {
		return nil
	}
	return dto.FromSubscriptionMatch(sm)
}

func ListRecentMatchPlayers(playerID string) []*dto.MatchPlayerDto {
	all := client.SubscriptionMatch.
		Query().
		Where(subscriptionmatch.PlayerID(playerID)).
		Limit(30).
		AllX(context.TODO())
	result := make([]*dto.MatchPlayerDto, len(all))
	for i, m := range all {
		result[i] = dto.FromSubscriptionMatch(m)
	}
	return result
}

type PlayerMatchCount struct {
	PlayerID string
	Count    int64
}

func GetMatchesCount(playerIds []string) (result []PlayerMatchCount) {
	//lastWeek := time.Now().Add(24 * time.Hour).Add(-24 * 7 * time.Hour)
	//t := time.Date(lastWeek.Year(), lastWeek.Month(), lastWeek.Day(), 0, 0, 0, 0, lastWeek.Location())

	//tx := db.Model(&dto.MatchPlayerDto{}).
	//	Select("player_id, count(*) as count").
	//	Where("player_id in ? AND created_at >= ?", playerIds, lastWeek).
	//	Group("player_id").
	//	Find(&result)
	//if tx.Error != nil {
	//	panic(fmt.Errorf("查询最近场次失败: %+v", tx.Error))
	//}
	//return result
	//client.SubscriptionMatch.Query().
	//	Where(
	//		subscriptionmatch.And(
	//			subscriptionmatch.PlayerIDIn(playerIds...),
	//			subscriptionmatch.CreateTimeGTE(t),
	//		),
	//	).
	//	GroupBy(subscriptionmatch.FieldPlayerID)
	return []PlayerMatchCount{}
}

func SaveMatchPlayer(mp *dto.MatchPlayerDto) {
	client.SubscriptionMatch.Create().
		SetMatchID(mp.MatchID).
		SetPlayerID(mp.PlayerID).
		SetPlayerSlot(mp.PlayerSlot).
		SetRadiantWin(mp.RadiantWin).
		SetDuration(mp.Duration).
		SetGameMode(mp.GameMode).
		SetLobbyType(mp.LobbyType).
		SetHeroID(mp.HeroID).
		SetStartTime(mp.StartTime).
		SetVersion(mp.Version).
		SetKills(mp.Kills).
		SetDeaths(mp.Deaths).
		SetAssists(mp.Assists).
		SetNillableSkill(mp.Skill).
		SetLeaverStatus(mp.LeaverStatus).
		SetPartySize(mp.PartySize).
		SaveX(context.TODO())
}
