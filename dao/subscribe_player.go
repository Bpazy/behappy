package dao

import (
	"github.com/Bpazy/really/models"
	"github.com/sirupsen/logrus"
)

func ListAllPlayerIDs() (pids []string) {
	var sps []models.SubscribePlayer
	db.Distinct("player_id").Find(&sps)
	if len(sps) == 0 {
		logrus.Infof("没有订阅的玩家")
	}

	for _, sp := range sps {
		pids = append(pids, sp.PlayerID)
	}
	return pids
}

func ListGroupIDs() []int {
	var groupIDs []int
	db.Model(&models.SubscribePlayer{}).Distinct("group_id").Find(&groupIDs)
	return groupIDs
}

func ListSubPlayers(playerID string) []models.SubscribePlayer {
	var allSub []models.SubscribePlayer
	if err := db.Where("player_id = ?", playerID).Find(&allSub).Error; err != nil {
		logrus.Info("没有订阅的玩家")
		return []models.SubscribePlayer{}
	}
	return allSub
}

func GetSubPlayer(groupID int, playerID string) *models.SubscribePlayer {
	query := &models.SubscribePlayer{
		GroupID:  groupID,
		PlayerID: playerID,
	}
	savedSP := models.SubscribePlayer{}
	if err := db.Where(query).First(&savedSP).Error; err != nil {
		return nil
	}
	return &savedSP
}

func GetSubPlayerMapByGroupId(groupID int) (m map[string]string) {
	var sps []models.SubscribePlayer
	if err := db.Where("group_id = ?", groupID).Find(&sps).Error; err != nil {
		return
	}

	m = map[string]string{}
	for _, sp := range sps {
		m[sp.PlayerID] = sp.Name()
	}
	return
}

func SaveSubPlayer(sp *models.SubscribePlayer) {
	db.Create(sp)
}

func UpdateSubPlayer(sp *models.SubscribePlayer) {
	db.Save(sp)
}
