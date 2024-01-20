package repository

import (
	"time"

	"github.com/shopspring/decimal"
)

type DeviceInteractionModel struct {
	InteractionID string          `gorm:"column:interaction_id"`
	Latitude      decimal.Decimal `gorm:"column:latitude"`
	Longitude     decimal.Decimal `gorm:"column:longitude"`
	DeviceID      string          `gorm:"column:device_id"`
	DeviceName    string          `gorm:"column:device_name"`
	Timestamp     time.Time       `gorm:"column:timestamp"`
	CreatedAt     time.Time       `gorm:"column:created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at"`
}
