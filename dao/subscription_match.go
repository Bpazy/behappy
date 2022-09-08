package dao

import (
	"context"
	"github.com/Bpazy/behappy/dto"
	"github.com/Bpazy/behappy/ent"
	"github.com/Bpazy/behappy/ent/subscriptionmatch"
	"time"
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
	PlayerID string `json:"player_id"`
	Count    int64  `json:"count"`
}

func GetMatchesCount(playerIds []string, start, end time.Time) (result []PlayerMatchCount) {
	client.SubscriptionMatch.Query().
		Where(
			subscriptionmatch.And(
				subscriptionmatch.PlayerIDIn(playerIds...),
				subscriptionmatch.CreateTimeGTE(start),
				subscriptionmatch.CreateTimeLT(end),
			),
		).
		GroupBy(subscriptionmatch.FieldPlayerID).
		Aggregate(ent.Count()).
		ScanX(context.TODO(), &result)
	return
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
