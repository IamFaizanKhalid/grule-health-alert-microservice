package data

import (
	"time"

	"../models"
)

// People array
var People = []models.Person{
	{
		Name:        "Zaeem",
		ID:          1,
		DateOfBirth: time.Date(1998, 8, 2, 1, 2, 3, 4, time.Local),
		Weight:      60,   // kg
		Height:      150, // cm
		Bmi:         0,
	},
	{
		Name:        "Fahad",
		ID:          2,
		DateOfBirth: time.Date(1993, 8, 2, 1, 2, 3, 4, time.Local),
		Weight:      20,   // kg
		Height:      130, // cm
		Bmi:         0,
	},
}
