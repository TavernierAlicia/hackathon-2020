package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var log *zap.Logger

func testdb() {
	log, _ = zap.NewProduction()

	defer log.Sync()
	//// IMPORT CONFIG ////
	viper.SetConfigName("conf")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Unable to load config file", zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
	}

	//// DEFINE CONFIG VARIABLES FROM JSON FILE ////
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.pass")
	dbname := viper.GetString("database.dbname")

	//// DB CONNECTION ////
	pathSQL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
	db, err := sqlx.Connect("mysql", pathSQL)
	if err != nil {
		log.Error("failed to connect database", zap.String("database", dbname),
			zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
		return

	} else {
		log.Info("Connexion etablished ", zap.String("database", dbname),
			zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
	}

	/*
		data := &Results{}
		//// Providers ////
		err = db.Select(&data.Result, testquery)
		if err != nil {
			log.Error("failed to request database ", zap.String("database", err.Error()),
				zap.String("query_name", "testquery"))
		} else {
			log.Info("Request Succeed ", zap.String("database", dbname),
				zap.String("query_name", "testquery"))
		}
		fmt.Println(data)
	*/
}
