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
	r.HandleFunc("/notifications/{id}", h.GetNotification).Methods("GET")
	r.HandleFunc("/notifications", h.AddNotification).Methods("POST")
	r.HandleFunc("/notifications/{id}", h.DeleteNotification).Methods("DELETE")

	// event routes
	r.HandleFunc("/events", h.GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", h.GetEvent).Methods("GET")
	r.HandleFunc("/events", h.AddEvent).Methods("POST")
	r.HandleFunc("/events/{id}", h.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", h.DeleteEvent).Methods("DELETE")

	log.Println("API is running!")
	http.ListenAndServe(":8000", r)
}
