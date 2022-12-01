package handlers

import (
	"events-service/pkg/entities"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
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
		fmt.Println(result.Error)
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
		fmt.Println(result.Error)
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
		log.Fatalln(err)
	}

	var event entities.Event
	json.Unmarshal(body, &event)

	// Append to the Books table
	if result := h.DB.Create(&event); result.Error != nil {
		fmt.Println(result.Error)
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

func (h handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {}
