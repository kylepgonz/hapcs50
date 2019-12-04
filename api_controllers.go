package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	//"io/ioutil"
	"github.com/go-chi/chi"
)

func apiEventListController(w http.ResponseWriter, r *http.Request) {
	type eventListType struct {
		Events []Event `json:"events"`
	}
	responseData := eventListType{Events: allEvents}
	jsonResponseData, err := json.MarshalIndent(responseData, "", " ")
	if err != nil {
		http.Error(w, "Could not serialize json data!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponseData)
}

func apiEventController(w http.ResponseWriter, r *http.Request) {
	type eventListType struct {
		Events Event `json:"events"`
	}
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

	responseData := eventListType{Events: event}
	jsonResponseData, err := json.MarshalIndent(responseData, "", " ")
	if err != nil {
		http.Error(w, "Could not serialize json data!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponseData)
}
