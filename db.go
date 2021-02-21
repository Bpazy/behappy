package really

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func initDB() *gorm.DB {
	userHomeDir, _ := os.UserHomeDir()
	dbPath := filepath.Join(userHomeDir, ".really.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&MatchPlayer{})
	err = db.AutoMigrate(&SubscribePlayer{})
	if err != nil {
		panic(err)
	}
	return db
}

// MatchPlayer 表 match_player 对应的结构
type MatchPlayer struct {
	MatchId     string // 比赛ID
	PlayerId    string // 玩家ID
	Hero        string // 选择的英雄
	MatchMode   string // 比赛模式
	MatchResult string // 比赛结果
	MatchKDA    string // 比赛KDA
	MatchLevel  string // 比赛级别
	gorm.Model
}

func (m MatchPlayer) String() string {
	return fmt.Sprintf("英雄: %s, 比赛ID: %s, 比赛模式: %s, 结果: %s, KDA: %s, 等级: %s", m.Hero, m.MatchId, m.MatchMode, m.MatchResult, m.MatchKDA, m.MatchLevel)
}

// SubscribePlayer 订阅 dotamax 更新
type SubscribePlayer struct {
	GroupId  string // 群ID
	PlayerId string // 玩家ID
	Alias    string // 别名
	gorm.Model
}
