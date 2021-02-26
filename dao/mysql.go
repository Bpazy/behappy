package dao

import (
	"github.com/Bpazy/really/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB

func InitDB() {
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
