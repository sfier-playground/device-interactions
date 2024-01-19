package handler

import (
	"time"
)

// DeviceSubmissionRequest ...
type DeviceSubmissionRequest struct {
	Timestamp time.Time   `json:"timestamp" validate:"required"`
	Location  GeoLocation `json:"location" validate:"required"`
	Devices   []Device    `json:"devices" validate:"required"`
}

// GeoLocation ...
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Device ...
type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
