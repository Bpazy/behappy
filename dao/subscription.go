package dao

import (
	"context"
	"github.com/Bpazy/behappy/dto"
	"github.com/Bpazy/behappy/ent"
	"github.com/Bpazy/behappy/ent/subscription"
)

func ListAllPlayerIDs() (pids []string) {
	return client.Subscription.
		Query().
		Unique(true).
		Select(subscription.FieldSteamID).
		StringsX(context.Background())
}

func ListSubPlayersByPlayerId(playerID string) []dto.SubscriptionDto {
	return convert(client.Subscription.
		Query().
		Where(subscription.SteamID(playerID)).
		AllX(context.Background()))
}

func ListSubPlayersByGroupId(groupID int) []dto.SubscriptionDto {
	return convert(client.Subscription.
		Query().
		Where(subscription.GroupID(groupID)).
		AllX(context.Background()))
}

func convert(subscriptions []*ent.Subscription) []dto.SubscriptionDto {
	result := make([]dto.SubscriptionDto, len(subscriptions))
	for i, s := range subscriptions {
		result[i] = dto.SubscriptionDto{
			GroupID:  s.GroupID,
			PlayerID: s.SteamID,
			Alias:    s.Alias,
		}
	}
	return result
}

func ListAllGroupIds() []int {
	return client.Subscription.
		Query().
		Unique(true).
		Select(subscription.FieldGroupID).
		IntsX(context.Background())
}

func GetSubscription(groupID int, playerID string) *ent.Subscription {
	return client.Subscription.
		Query().
		Where(
			subscription.And(
				subscription.GroupID(groupID),
				subscription.SteamID(playerID),
			),
		).
		FirstX(context.Background())
}

func GetSubscriptionDto(groupID int, playerID string) *dto.SubscriptionDto {
	s := GetSubscription(groupID, playerID)
	return &dto.SubscriptionDto{
		GroupID:  s.GroupID,
		PlayerID: s.SteamID,
		Alias:    s.Alias,
	}
}

func GetSubPlayerMapByGroupId(groupID int) map[string]string {
	m := map[string]string{}
	for _, sp := range ListSubPlayersByGroupId(groupID) {
		m[sp.PlayerID] = sp.Name()
	}
	return m
}

func SaveSubPlayer(groupId int, steamId string, alias string) {
	client.Subscription.
		Create().
		SetGroupID(groupId).
		SetSteamID(steamId).
		SetAlias(alias).
		SaveX(context.Background())
}

func UpdateSubPlayer(s *ent.Subscription) {
	s.Update().SaveX(context.Background())
}
