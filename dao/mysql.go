package dao

import (
	"context"
	"github.com/Bpazy/behappy/config"
	"github.com/Bpazy/behappy/ent"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var client *ent.Client

func InitDB() {
	initEnt()

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
