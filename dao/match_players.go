package dao

import "github.com/Bpazy/really/models"

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

func SaveMatchPlayer(mp *models.MatchPlayer) {
	db.Create(mp)
}
