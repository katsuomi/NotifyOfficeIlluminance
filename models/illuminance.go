package models

// Illuminance is Illuminance models property
type Illuminance struct {
	ID    uint   `json:"id" binding:"required"`
	Illuminance  int `json:"illuminance" binding:"required"`	
}
