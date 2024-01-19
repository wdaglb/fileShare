package core

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

func initDb() {
	var err error
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  false,
		})
	DB, err = gorm.Open(sqlite.Open(Config.Storage.DbFile), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panic(err)
	}
	//DB.AutoMigrate()
}
