package appInit

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connection struct {
	DB *gorm.DB
}

const dbParams = "?charset=utf8mb4&parseTime=true&timeout=5s&rejectReadOnly=true&loc=Asia%2FCalcutta"

func InitializeDatabase(ctx context.Context) *Connection {
	gormLogger := logger.Default
	gormLogger = gormLogger.LogMode(logger.Info)
	userName := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	dbName := viper.GetString("mysql.database")

	dsn := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + dbParams
	log.Printf("Connecting to database: %s", dsn)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatalf("Unable to make gorm connection | Error: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Unable to get sqlDB from gormDB | Error: %v", err)
	}

	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConnections"))
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConnections"))

	retries := 3
	for retries > 0 {
		err = sqlDB.Ping()
		if err != nil {
			log.Printf("Unable to ping database server: %s, waiting 2 seconds before trying %d more times", err.Error(), retries)
			time.Sleep(time.Second * 2)
			retries--
		} else {
			err = nil
			break
		}
	}
	if err != nil {
		panic("Db not initialised")
	}

	conn := Connection{
		DB: gormDB,
	}
	return &conn
}
