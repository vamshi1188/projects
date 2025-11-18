package models

type Service struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	DurationMinutes int    `json:"durationMinutes"`
	Price           int    `json:"price"`
	Description     string `json:"description,omitempty"`
}

type Booking struct {
	ID        int    `json:"id"`
	ServiceID int    `json:"serviceId"`
	Customer  string `json:"customer"`
	Phone     string `json:"phone"`
	Status    string `json:"status,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version,omitempty"`
}
