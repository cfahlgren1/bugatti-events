package entities

import (
	"time"
)

// Event struct (Model)
type Event struct {
	ID          string    `json:"id"`
	Name        string    `json:"name`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}

// Notification struct (Model)
type Notification struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	SentAt    time.Time `json:"sent_at"`
	Event   Event    `json:"event_id" gorm:"foreignkey:EventID"`
	EventID string
}