package entities

import (
	"time"
)

// Event struct (Model)
type Event struct {
	ID          string    `json:"-"`
	Name        string    `json:"name`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-,omitempty"`
}

type Notification struct {
    ID        string    `json:"-"`
    Message   string    `json:"message" validate:"required"`
    SentAt    time.Time `json:"sent_at"`
    Event     Event     `json:"-" gorm:"foreignkey:EventID`
    EventID   string    `json:"event_id" validate:"required"`
    CreatedAt time.Time `json:"-"`
    UpdatedAt time.Time `json:"-,omitempty"`
    DeletedAt time.Time `json:"-,omitempty"`
}
