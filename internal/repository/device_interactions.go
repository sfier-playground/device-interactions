package repository

import (
	"os"
	"time"

	"github.com/sifer169966/device-interactions/internal/core/domain"
	"github.com/sifer169966/device-interactions/pkg/apperror"
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
	query, err := os.ReadFile("./internal/repository/device_interactions.sql")
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

func (r *DeviceInteractions) CreateMany(in domain.DeviceSubmission) error {
	deviceInteractionModels := make([]*DeviceInteractionModel, len(in.Devices))
	now := time.Now().UTC()
	for i := range in.Devices {
		deviceInteractionModels[i] = &DeviceInteractionModel{
			InteractionID: in.Devices[i].GetInteractionID(),
			Latitude:      in.Location.Latitude,
			Longitude:     in.Location.Longitude,
			DeviceID:      in.Devices[i].DeviceID,
			DeviceName:    in.Devices[i].Name,
			Timestamp:     in.Timestamp,
			CreatedAt:     now,
			UpdatedAt:     now,
		}
	}

	err := r.db.Create(deviceInteractionModels).Error
	if err != nil {
		logger.Error("could not insert device_interactions", "error", err)
		return apperror.NewInternalServerError()
	}
	return nil
}
