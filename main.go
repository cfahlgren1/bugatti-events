package main

import (
	"log"
	"net/http"

	"events-service/pkg/db"
	"events-service/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	DB := db.Init()
	h := handlers.New(DB)
	r := mux.NewRouter()

	r.HandleFunc("/events", h.GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", h.GetEvent).Methods("GET")
	r.HandleFunc("/events", h.AddEvent).Methods("POST")
	r.HandleFunc("/events/{id}", h.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", h.DeleteEvent).Methods("DELETE")

	log.Println("API is running!")
	http.ListenAndServe(":8000", r)
}
