package dao

import (
	"github.com/Bpazy/behappy/models"
	"github.com/sirupsen/logrus"
)

func GetMatchPlayer(matchID int64, playerID string) *models.MatchPlayer {
	if matchID == 0 || playerID == "" {
		return nil
	}
	s := map[string]interface{}{
		"match_id":  matchID,
		"player_id": playerID,
	}
	mp := models.MatchPlayer{}
	if err := db.Where(s).First(&mp).Error; err != nil {
		return nil
	}
	return &mp
}

func ListRecentMatchPlayers(playerID string) []*models.MatchPlayer {
	s := map[string]interface{}{
		"player_id": playerID,
	}
	var mp []*models.MatchPlayer
	if err := db.Where(s).Order("created_at desc").Limit(30).Find(&mp).Error; err != nil {
		logrus.Errorf("查询最近场次失败: %+v", err)
		return mp
	}
	return mp
}

func SaveMatchPlayer(mp *models.MatchPlayer) {
	db.Create(mp)
}
