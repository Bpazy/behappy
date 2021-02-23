package really

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
)

func initDB() *gorm.DB {
	userHomeDir, _ := os.UserHomeDir()
	dbPath := filepath.Join(userHomeDir, ".really.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	err = db.AutoMigrate(&MatchPlayer{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&SubscribePlayer{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Hero{})
	if err != nil {
		panic(err)
	}
	return db
}

// MatchPlayer 表 match_player 对应的结构
type MatchPlayer struct {
	MatchID      int64 `json:"match_id"`
	PlayerID     string
	PlayerSlot   int  `json:"player_slot"`
	RadiantWin   bool `json:"radiant_win"`
	Duration     int  `json:"duration"`
	GameMode     int  `json:"game_mode"`
	LobbyType    int  `json:"lobby_type"`
	HeroID       int  `json:"hero_id"`
	StartTime    int  `json:"start_time"`
	Version      int  `json:"version"`
	Kills        int  `json:"kills"`
	Deaths       int  `json:"deaths"`
	Assists      int  `json:"assists"`
	Skill        int  `json:"skill"`
	LeaverStatus int  `json:"leaver_status"`
	PartySize    int  `json:"party_size"`
	gorm.Model
}

func (m MatchPlayer) HeroName() string {
	var hero Hero
	db.Where(&Hero{
		ID: m.HeroID,
	}).Find(&hero)

	if hero.ID != 0 {
		return hero.LocalizedName
	}
	return strconv.Itoa(m.HeroID)
}

func (m MatchPlayer) IsWin() bool {
	return m.PlayerSlot < 127 && m.RadiantWin
}

func (m MatchPlayer) MatchResultString() string {
	if m.IsWin() {
		return "胜利"
	}
	return "失败"
}

func (m MatchPlayer) SkillString() string {
	switch m.Skill {
	case 3:
		return "Very High"
	case 2:
		return "High"
	}
	return "Normal"
}

// SubscribePlayer 订阅 dotamax 更新
type SubscribePlayer struct {
	GroupId  int    // 群ID
	PlayerId string // 玩家ID
	Alias    string // 别名
	gorm.Model
}

func (sp SubscribePlayer) Name() string {
	if sp.Alias != "" {
		return sp.Alias
	}
	return sp.PlayerId
}

type Hero struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	LocalizedName string `json:"localized_name"`
	gorm.Model
}

type SteamApiResult struct {
	Result struct {
		Heroes []Hero `json:"heroes"`
		Status int    `json:"status"`
		Count  int    `json:"count"`
	} `json:"result"`
}
