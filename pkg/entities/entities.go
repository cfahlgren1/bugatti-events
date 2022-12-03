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

type Notification struct {
    ID        string    `json:"id"`
    Message   string    `json:"message" validate:"required"`
    SentAt    time.Time `json:"sent_at"`
    Event     Event     `gorm:"foreignkey:EventID`
    EventID   string    `json:"event_id" validate:"required"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    DeletedAt time.Time `json:"deleted_at,omitempty"`
}
