package models

// UpdatePersonRequest is request model for /text API
type UpdatePersonRequest struct {
	Height float64 `json:"height"`
	Weight  float64 `json:"width"`
}
