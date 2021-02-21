package really

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"
)

func initDB() *sql.DB {
	userHomeDir, _ := os.UserHomeDir()
	dbPath := filepath.Join(userHomeDir, ".really.db")
	_, dbNotExistsErr := os.Open(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	if dbNotExistsErr != nil {
		sqlStmt := `
			CREATE TABLE "match_player" (
			  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			  "match_id" text NOT NULL,
			  "player_id" text NOT NULL,
			  "hero" text NOT NULL,
			  "match_mode" text NOT NULL,
			  "match_result" TEXT NOT NULL,
			  "match_kda" TEXT NOT NULL,
			  "match_level" TEXT NOT NULL,
			  "create_time" TEXT NOT NULL,
			  "modify_time" TEXT NOT NULL
			);
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			panic(err)
		}
	}
	return db
}

type MatchPlayer struct {
	Id          string    // 主键
	MatchId     string    // 比赛ID
	PlayerId    string    // 玩家ID
	Hero        string    // 选择的英雄
	MatchMode   string    // 比赛模式
	MatchResult string    // 比赛结果
	MatchKDA    string    // 比赛KDA
	MatchLevel  string    // 比赛级别
	CreateTime  time.Time // 记录创建时间
	ModifyTime  time.Time // 记录修改时间
}
