package really

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"time"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.DataSource.Url), &gorm.Config{
		Logger: logger.New(
			logrus.StandardLogger(), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Warn, // Log level
				Colorful:      false,       // Disable color
			},
		),
	})
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}

// CREATE TABLE `match_players` (
//   `match_id` bigint(20) DEFAULT NULL,
//   `player_id` varchar(100) DEFAULT NULL,
//   `player_slot` bigint(20) DEFAULT NULL,
//   `radiant_win` tinyint(1) DEFAULT NULL,
//   `duration` bigint(20) DEFAULT NULL,
//   `game_mode` bigint(20) DEFAULT NULL,
//   `lobby_type` bigint(20) DEFAULT NULL,
//   `hero_id` bigint(20) DEFAULT NULL,
//   `start_time` bigint(20) DEFAULT NULL,
//   `version` bigint(20) DEFAULT NULL,
//   `kills` bigint(20) DEFAULT NULL,
//   `deaths` bigint(20) DEFAULT NULL,
//   `assists` bigint(20) DEFAULT NULL,
//   `skill` bigint(20) DEFAULT NULL,
//   `leaver_status` bigint(20) DEFAULT NULL,
//   `party_size` bigint(20) DEFAULT NULL,
//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
//   `created_at` datetime(3) DEFAULT NULL,
//   `updated_at` datetime(3) DEFAULT NULL,
//   `deleted_at` datetime(3) DEFAULT NULL,
//   PRIMARY KEY (`id`),
//   KEY `idx_match_id` (`match_id`),
//   KEY `idx_player_id` (`player_id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type MatchPlayer struct {
	MatchID      int64  `json:"match_id" gorm:"index"`
	PlayerID     string `gorm:"index"`
	PlayerSlot   int    `json:"player_slot"`
	RadiantWin   bool   `json:"radiant_win"`
	Duration     int    `json:"duration"`
	GameMode     int    `json:"game_mode"`
	LobbyType    int    `json:"lobby_type"`
	HeroID       int    `json:"hero_id"`
	StartTime    int    `json:"start_time"`
	Version      int    `json:"version"`
	Kills        int    `json:"kills"`
	Deaths       int    `json:"deaths"`
	Assists      int    `json:"assists"`
	Skill        *int   `json:"skill"`
	LeaverStatus int    `json:"leaver_status"`
	PartySize    int    `json:"party_size"`
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
	if m.PlayerSlot < 127 {
		return m.RadiantWin
	}
	return !m.RadiantWin
}

func (m MatchPlayer) MatchResultString() string {
	if m.IsWin() {
		return "胜利"
	}
	return "失败"
}

func (m MatchPlayer) SkillString() string {
	if m.Skill == nil {
		return "Unknown"
	}
	switch *m.Skill {
	case 3:
		return "Very High"
	case 2:
		return "High"
	}
	return "Normal"
}

//  CREATE TABLE `subscribe_players` (
//   `group_id` bigint(20) DEFAULT NULL,
//   `player_id` varchar(100) DEFAULT NULL,
//   `alias` varchar(100) DEFAULT NULL,
//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
//   `created_at` datetime(3) DEFAULT NULL,
//   `updated_at` datetime(3) DEFAULT NULL,
//   `deleted_at` datetime(3) DEFAULT NULL,
//   PRIMARY KEY (`id`),
//   KEY `idx_group_id` (`group_id`),
//   KEY `idx_player_id` (`player_id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type SubscribePlayer struct {
	GroupId  int    `gorm:"index"` // 群ID
	PlayerId string `gorm:"index"` // 玩家ID
	Alias    string // 别名
	gorm.Model
}

func (sp SubscribePlayer) Name() string {
	if sp.Alias != "" {
		return sp.Alias
	}
	return sp.PlayerId
}

// CREATE TABLE `heros` (
//   `name` varchar(200) DEFAULT NULL,
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `localized_name` varchar(200) DEFAULT NULL,
//   `created_at` datetime(3) DEFAULT NULL,
//   `updated_at` datetime(3) DEFAULT NULL,
//   `deleted_at` datetime(3) DEFAULT NULL,
//   PRIMARY KEY (`id`),
//   KEY `idx_heros_id` (`id`)
// ) ENGINE=InnoDB AUTO_INCREMENT=130 DEFAULT CHARSET=utf8mb4;
type Hero struct {
	Name          string `json:"name"`
	ID            int    `json:"id" gorm:"index"`
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

// CREATE TABLE `jokes` (
//   `content` longtext NOT NULL,
//   `limited_hero_id` bigint(20) unsigned DEFAULT NULL,
//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
//   `created_at` datetime(3) DEFAULT NULL,
//   `updated_at` datetime(3) DEFAULT NULL,
//   `deleted_at` datetime(3) DEFAULT NULL,
//   PRIMARY KEY (`id`),
//   KEY `idx_hero_id` (`limited_hero_id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
type Joke struct {
	Content       string
	LimitedHeroId int `gorm:"index"`
	gorm.Model
}
