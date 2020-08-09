package api

import (
	"github.com/dandelabs/ghostbuster-backend-libs/dandelog"
	"github.com/dandelabs/ghostbuster-backend-libs/db"
	"github.com/dandelabs/ghostbuster-backend-webhooks/config"
	"os"
)

var cfg *config.Config

func init() {
	method := "init:"
	environment := os.Getenv("environment")
	config.SetEnvironment(environment)
	cfg = &config.Cfg
	err := db.StartDBCon(cfg.DB.URL)
	if err != nil {
		dandelog.Error.Println(method + " It could not database connection:" + err.Error())
		os.Exit(3)
	}

}
