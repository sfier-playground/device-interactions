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
	Devices   []Device    `json:"devices" validate:"gte=1,lte=2,dive"`
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
	Latitude  decimal.Decimal `json:"latitude" validate:"required,dgte=-85.05115,dlte=85"`
	Longitude decimal.Decimal `json:"longitude" validate:"required,dgte=-180,dlte=180"`
}

// Device ...
type Device struct {
	ID   string `json:"id" validate:"uuid_rfc4122"`
	Name string `json:"name" validate:"device_name,max=15"`
}
