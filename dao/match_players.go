package dao

import (
	"fmt"
	"github.com/Bpazy/behappy/models"
	"github.com/sirupsen/logrus"
	"time"
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

type PlayerMatchCount struct {
	PlayerID string
	Count    int64
}

func GetMatchesCount(playerIds []string) (result []PlayerMatchCount) {
	lastWeek := time.Now().Add(24 * time.Hour).Add(-24 * 7 * time.Hour).Format("2006-01-02")
	tx := db.Model(&models.MatchPlayer{}).
		Select("player_id, count(*) as count").
		Where("player_id in ? AND created_at >= ?", playerIds, lastWeek).
		Group("player_id").
		Find(&result)
	if tx.Error != nil {
		panic(fmt.Errorf("查询最近场次失败: %+v", tx.Error))
	}
	return result
}

func SaveMatchPlayer(mp *models.MatchPlayer) {
	db.Create(mp)
}
