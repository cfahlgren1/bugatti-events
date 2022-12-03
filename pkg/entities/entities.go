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
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-,omitempty"`
}

type PhoneNumber struct {
	ID         string `json:"id"`
	Number     string `json:"number"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  time.Time `json:"-,omitempty"`
}

type Relationship struct {
	ID        string `json:"-" sql:"DEFAULT:uuid_generate_v4()" gorm:"primary_key"`
	Event     Event `json:"event"`
	EventID   string `json:"event_id"`
	PhoneNumber PhoneNumber `json:"phone_number"`
	PhoneNumberID string `json:"phone_number_id"`
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

// SignupRequest is the request body for the SignupHandler.
type SignupRequest struct {
    EventID string `json:"id" validate:"required"`
    PhoneNumber string `json:"phone_number" validate:"required"`
}
