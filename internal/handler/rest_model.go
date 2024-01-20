package handler

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/sifer169966/device-interactions/internal/core/domain"
)

// DeviceSubmissionRequest ...
type DeviceSubmissionRequest struct {
	Timestamp time.Time   `json:"timestamp" validate:"required"`
	Location  GeoLocation `json:"location" validate:"required"`
	Devices   []Device    `json:"devices" validate:"required"`
}

func (t DeviceSubmissionRequest) toDeviceSubmissionDomain() domain.DeviceSubmission {
	deviceLenght := len(t.Devices)
	devices := make([]domain.Device, deviceLenght)
	for i := 0; i < deviceLenght; i++ {
		devices[i].DeviceID = t.Devices[i].ID
		devices[i].Name = t.Devices[i].Name
	}
	return domain.DeviceSubmission{
		Timestamp: t.Timestamp,
		Location: domain.GeoLocation{
			Latitude:  t.Location.Latitude,
			Longitude: t.Location.Longitude,
		},
		Devices: devices,
	}
}

// GeoLocation ...
type GeoLocation struct {
	Latitude  decimal.Decimal `json:"latitude"`
	Longitude decimal.Decimal `json:"longitude"`
}

// Device ...
type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
