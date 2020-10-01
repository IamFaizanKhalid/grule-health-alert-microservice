package models

import "time"

// Person struct
type Person struct {
	Name        string    `json:"name"`
	ID          int64     `json:"id"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Weight      float64   `json:"weight"` // kg
	Height      float64   `json:"height"` // cm
	Bmi         float64   `json:"bmi"` // weight (kg) / [height (m)]2
}

func (p *Person) UpdateBMI() float64 {
	p.Bmi = p.Weight/ ((p.Height/100)*(p.Height/100))
	return p.Bmi
}
