package dao

import (
	"context"
	"github.com/Bpazy/behappy/berrors"
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/dto"
	"github.com/Bpazy/behappy/ent"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB
var client *ent.Client

func InitDB() {
	initGorm()
	initEnt()

	berrors.Must(db.AutoMigrate(
		&dto.MatchPlayerDto{},
	))

	if err := client.Schema.Create(context.Background()); err != nil {
		logrus.Fatalf("failed creating schema resources: %v", err)
	}
}

func initEnt() {
	_client, err := ent.Open("mysql", config.GetConfig().DataSource.Url)
	if err != nil {
		logrus.Fatalf("failed opening connection to mysql: %v", err)
	}
	client = _client
}

func initGorm() {
	_db, err := gorm.Open(mysql.Open(config.GetConfig().DataSource.Url), &gorm.Config{
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

	db = _db
}
