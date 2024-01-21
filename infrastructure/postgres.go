package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/sifer169966/go-logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// implement any connection to your infrastructure such as database, rabbitMQ etc.

type PostgreSQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Debug    bool
}

func NewPostgreSQLConnection(cfg PostgreSQLConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	logLevel := gormLogger.Error
	if cfg.Debug {
		logLevel = gormLogger.Info
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(logLevel),
	})
	if err != nil {
		logger.Fatal("could not connection to postgresql", "error", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("could not get sqldb connection", "error", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err = sqlDB.PingContext(ctx)
	if err != nil {
		logger.Fatal("could not ping to postgresql server", "error", err)
	}
	return db
}
