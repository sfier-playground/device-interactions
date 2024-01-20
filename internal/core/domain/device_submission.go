package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type DeviceSubmission struct {
	Timestamp time.Time
	Location  GeoLocation
	Devices   []Device
}

type GeoLocation struct {
	Latitude  decimal.Decimal
	Longitude decimal.Decimal
}

type Device struct {
	interactionID string
	DeviceID      string
	Name          string
}

func (d *Device) SetInteractionID() {
	d.interactionID = uuid.NewString()
}

func (d *Device) GetInteractionID() string {
	return d.interactionID
}
