//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type DeviceInteractions struct {
	InteractionID string `sql:"primary_key"`
	Latitude      float64
	Longitude     float64
	DeviceID      string
	DeviceName    string
	Timestamp     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}