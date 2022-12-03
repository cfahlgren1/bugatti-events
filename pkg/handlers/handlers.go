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
	validator "github.com/asaskevich/govalidator"
	"github.com/satori/go.uuid"

	"events-service/pkg/sms"
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

// AddNotification creates a new notification in the database.
func (h handler) AddNotification(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	
	// Unmarshal the request body into a Notification struct
    var notification entities.Notification
    json.Unmarshal(body, &notification)

    // Get the values from the Notification struct
    message := notification.Message
    eventID := notification.EventID

    // Generate a new UUID for the Notification struct
    id := uuid.NewV4()

    // Create a new Notification struct with the values from the request
    newNotification := entities.Notification{
        ID:      id.String(),
        Message: message,
        EventID: eventID,
    }

    // Validate the notification struct
    if valid, err := validator.ValidateStruct(notification); valid != true {
        log.Println(err)
        http.Error(w, fmt.Sprintf("invalid request %s", err), http.StatusBadRequest)
        return
    }

	sms.SendSms([]string{}, message)

	// Look up the associated Event using the EventID field
    var event entities.Event
    if result := h.DB.Where("id = ?", eventID).First(&event); result.Error != nil {
        log.Println(result.Error)
        http.Error(w, "Invalid EventID", http.StatusBadRequest)
        return
    }

    // Set the Event field in the Notification
    notification.Event = event

    // Append to the Notifications table
    if result := h.DB.Create(&newNotification); result.Error != nil {
        log.Println(result.Error)
        http.Error(w, "Failed to create notification", http.StatusInternalServerError)
        return
    }

    // Send a 201 created response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    log.Println("Created new notification", &notification.ID)
    json.NewEncoder(w).Encode("Created")
}

func (h handler) GetNotificationsByEvent(w http.ResponseWriter, r *http.Request) {
	// Get the event ID from the path parameter
	eventID := mux.Vars(r)["event"]

	// Query the notifications for the given event ID
	var notifications []entities.Notification
	if result := h.DB.Where("event_id = ?", eventID).Find(&notifications); result.Error != nil {
		log.Fatal(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notifications)
}

// Handle processes a signup request and adds the provided phone number to the specified event.
func (h handler) AddPhoneNumberToEvent(w http.ResponseWriter, r *http.Request) {
    // Parse and validate the request body.
    var req entities.SignupRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if valid, err := validator.ValidateStruct(req); err != nil || !valid {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Check if the event exists.
    var event entities.Event
    if result := h.DB.First(&event, "id = ?", req.EventID); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

	// Check if the phone number exists, if not create it
	phoneNumber := entities.PhoneNumber{
		ID: req.PhoneNumber,
		Number: req.PhoneNumber,
	}
	
	if result := h.DB.First(&phoneNumber, "number = ?", req.PhoneNumber); result.Error != nil {
		if result := h.DB.Create(&phoneNumber); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Created phone number", phoneNumber)
	}

    // Check if the phone number is already signed up for the event.
    var existingRelationship entities.Relationship
    if result := h.DB.First(&existingRelationship, "event_id = ? AND phone_number_id = ?", req.EventID, req.PhoneNumber); result.Error == nil {
        http.Error(w, "phone number is already signed up for this event", http.StatusConflict)
        return
    }

    // Create a new relationship between the event and the phone number.
	relationship := entities.Relationship{
		ID: uuid.NewV4().String(), 
		EventID: req.EventID,
		PhoneNumberID: req.PhoneNumber,
	}
	
    if result := h.DB.Create(&relationship); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

	// return successful response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Println("Created new event", &event.ID)
	json.NewEncoder(w).Encode("Created")
}

func (h handler) ShowPhoneNumbersForEvent(w http.ResponseWriter, r *http.Request) {
    // Parse the event ID from the request path.
    eventID := mux.Vars(r)["event"]

    // Fetch the event from the database.
    var event entities.Event
    if result := h.DB.First(&event, "id = ?", eventID); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

	// Fetch the phone numbers associated with the event from the database.
	var phoneNumbers []entities.PhoneNumber
	if result := h.DB.Table("phone_numbers").Joins("LEFT JOIN relationships ON phone_numbers.id = relationships.phone_number_id").Where("relationships.event_id = ?", event.ID).Find(&phoneNumbers); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

    // Convert the phone numbers to a slice of strings and return them as a response.
    numbers := make([]string, len(phoneNumbers))
    for i, number := range phoneNumbers {
        numbers[i] = number.Number
    }
    json.NewEncoder(w).Encode(numbers)
}