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
		StringsX(context.TODO())
}

func ListSubPlayersByPlayerId(playerID string) []dto.SubscriptionDto {
	//var allSub []models.SubscriptionDto
	//if err := db.Where("player_id = ?", playerID).Find(&allSub).Error; err != nil {
	//	logrus.Info("没有订阅的玩家")
	//	return []models.SubscriptionDto{}
	//}
	//return allSub
	return convert(client.Subscription.
		Query().
		Where(subscription.SteamID(playerID)).
		AllX(context.TODO()))
}

func ListSubPlayersByGroupId(groupID int) []dto.SubscriptionDto {
	return convert(client.Subscription.
		Query().
		Where(subscription.GroupID(groupID)).
		AllX(context.TODO()))
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
		IntsX(context.TODO())
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
		FirstX(context.TODO())
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
		SaveX(context.TODO())
}

func UpdateSubPlayer(s *ent.Subscription) {
	s.Update().SaveX(context.TODO())
}
