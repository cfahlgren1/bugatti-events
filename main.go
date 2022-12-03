package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"events-service/pkg/db"
	"events-service/pkg/handlers"
)

func main() {

	DB := db.Init()
	h := handlers.New(DB)
	r := mux.NewRouter()

	// notification routes
	r.HandleFunc("/notifications", h.GetAllNotifications).Methods("GET")
	r.HandleFunc("/notifications", h.AddNotification).Methods("POST")

	// get notifications from event
	r.HandleFunc("/notifications/{event}", h.GetNotificationsByEvent).Methods("GET")

	// event routes
	r.HandleFunc("/events", h.GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", h.GetEvent).Methods("GET")
	r.HandleFunc("/events", h.AddEvent).Methods("POST")
	r.HandleFunc("/events/{id}", h.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", h.DeleteEvent).Methods("DELETE")

	// phone number routes
	r.HandleFunc("/sms", h.AddPhoneNumberToEvent).Methods("POST")
	r.HandleFunc("/sms/{event}", h.ShowPhoneNumbersForEvent).Methods("GET")

	log.Println("API is running!")
	http.ListenAndServe(":8000", r)
}
