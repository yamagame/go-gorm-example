package model

import (
	"time"
)

type Address struct {
	Addrid     int64 `gorm:"primaryKey;autoIncrement:true"`
	Postalcode *string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
