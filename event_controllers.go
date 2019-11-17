package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	//"io/ioutil"
)

// type indexContData struct {
// 	Events    Event
// 	Attending []string
// }

func eventDetailController(w http.ResponseWriter, r *http.Request) {
	eventIDstring := chi.URLParam(r, "eventID")
	eventID, err := strconv.Atoi(eventIDstring)
	if err != nil {
		http.Error(w, "Invalid event ID!", http.StatusBadRequest)
		return
	}
	event, foundEvent := getEventByID(eventID)
	if foundEvent == false {
		http.Error(w, "No such event!", http.StatusNotFound)
		return
	}
	// eventData := indexContData{
	// 	Events:    event,
	// 	Attending: event.Attending,
	// }
	tmpl["event-detail"].Execute(w, event)
}
