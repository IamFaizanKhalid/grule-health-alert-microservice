package models

import "time"

// Person struct
type Person struct {
	Name        string    `json:"name"`
	ID          int64     `json:"id"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Weight      float64   `json:"weight"` // kg
	Height      float64   `json:"height"` // cm
	Bmi         float64   `json:"bmi"`
}
