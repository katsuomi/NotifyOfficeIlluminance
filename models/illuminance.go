package models

import (
	"time"
)

// Illuminance is Illuminance models property
type Illuminance struct {
	ID          uint      `json:"id"`
	Illuminance int       `json:"illuminance" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}
