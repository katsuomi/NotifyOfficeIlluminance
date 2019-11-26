package models

// Illuminance is Illuminance models property
type Illuminance struct {
	ID    uint   `json:"id" binding:"required"`
	illuminance  string `json:"illuminance" binding:"required"`	
}
