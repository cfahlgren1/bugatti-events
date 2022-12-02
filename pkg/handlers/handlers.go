package handlers

import (
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"events-service/pkg/entities"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	var events []entities.Event

	if result := h.DB.Find(&events); result.Error != nil {
		log.Fatal(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func (h handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := vars["id"]

	// Find Event by Id
	var event entities.Event

	if result := h.DB.First(&event, "id = ?", id); result.Error != nil {
		log.Fatal(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)

}

// Create an Event in the database
func (h handler) AddEvent(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var event entities.Event
	json.Unmarshal(body, &event)

	// Append to the Books table
	if result := h.DB.Create(&event); result.Error != nil {
		log.Fatal(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Println("Created new event", &event.ID)
	json.NewEncoder(w).Encode("Created")
}

func (h handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the event by Id
	var event entities.Event

	if result := h.DB.First(&event, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that event
	h.DB.Delete(&event)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

// Update an Event in the database
func (h handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := vars["id"]

	// Read the request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// Update the Event
	var event entities.Event
	json.Unmarshal(body, &event)

	// Check if the Event exists
	if result := h.DB.First(&event, "id = ?", id); result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Update the Event
	event.UpdatedAt = time.Now()
	if result := h.DB.Save(&event); result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send a 200 OK response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("Updated event", &event.ID)
	json.NewEncoder(w).Encode("OK")
}

// GetAllNotifications returns all notifications in the database.
func (h handler) GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	var notifications []entities.Notification

	if result := h.DB.Preload("Event").Find(&notifications); result.Error != nil {
		log.Fatal(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notifications)
}

// GetNotification returns the notification with the given ID.
func (h handler) GetNotification(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := vars["id"]

	// Find notification by Id
	var notification entities.Notification

	if result := h.DB.First(&notification, "id = ?", id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notification)
}

// AddNotification creates a new notification in the database.
func (h handler) AddNotification(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var notification entities.Notification
	json.Unmarshal(body, &notification)

	// Append to the Notifications table
	if result := h.DB.Create(&notification); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Println("Created new notification", &notification.ID)
	json.NewEncoder(w).Encode("Created")
}

// DeleteNotification deletes the notification with the given ID.
func (h handler) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the notification by Id
	var notification entities.Notification

	if result := h.DB.First(&notification, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that notification
	h.DB.Delete(&notification)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
