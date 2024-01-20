package repository

import (
	"os"

	"github.com/sifer169966/device-interactions/internal/core/domain"
	"gorm.io/gorm"

	"github.com/sifer169966/go-logger"
)

/*
	|--------------------------------------------------------------------------
	| The Repository Adaptor
	|--------------------------------------------------------------------------
	|
	| An Adapter will initiate the interaction with the Application through
	| a Port, using specific technology that means you can choose
	| any technology you want for your application or business logic.
	|
*/

type DeviceInteractions struct {
	db *gorm.DB
}

func NewDeviceInteractions(db *gorm.DB) *DeviceInteractions {
	query, err := os.ReadFile("./device_interactions.sql")
	if err != nil {
		logger.Fatal("could not found the sql file", "error", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("could not get sqldb for table preparation", "error", err)
	}
	_, err = sqlDB.Exec(string(query))
	if err != nil {
		logger.Fatal("could not exec the sql file", "error", err)
	}
	return &DeviceInteractions{
		db: db,
	}
}

func (r *DeviceInteractions) CreateMany(d domain.DeviceSubmission) error {

	return nil
}
