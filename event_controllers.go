package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	//"io/ioutil"
)

type eventDetailContextData struct {
	Event        Event
	FormErrors   string
	FormMessages string
}

type eventCreateContextData struct {
	FormErrors   string
	FormMessages string
}

func eventCreateController(w http.ResponseWriter, r *http.Request) {
	errorsWeFound := ""
	messagesWeFound := ""
	if r.Method == http.MethodPost {
		//get data on event from create page
		r.ParseForm()
		title := strings.ToLower(r.PostFormValue("title"))
		location := strings.ToLower(r.PostFormValue("location"))
		image := r.PostFormValue("image")
		date, _ := time.Parse("2006-01-02T15:04", r.PostFormValue("date"))
		category := r.PostFormValue("category")
		description := r.PostFormValue("description")
		id := getMaxEventID() + 1

		today := time.Now()
		if today.After(date) {
			errorsWeFound += "Wrong date!"
		} else if id == -1 {
			errorsWeFound += "No more event ids!"
		} else if categoryExists[category] == false {
			// println(category)
			// println("test")
			errorsWeFound += "Wrong category!"
		} else {
			newEvent := Event{
				ID:          id,
				Title:       title,
				Date:        date,
				Image:       image,
				Location:    location,
				Category:    category,
				Description: description,
				Attending:   []string{},
			}
			addEvent(newEvent)
			_, foundEvent := getEventByID(id)
			if foundEvent == false {
				http.Error(w, "Error in event creation!", http.StatusInternalServerError)
				return
			}
			messagesWeFound += "Event creation successful!"
		}
	}
	eventData := eventCreateContextData{
		FormErrors:   errorsWeFound,
		FormMessages: messagesWeFound,
	}
	tmpl["event-new"].Execute(w, eventData)
}

func eventDetailController(w http.ResponseWriter, r *http.Request) {
	eventIDstring := chi.URLParam(r, "eventID")
	eventID, err := strconv.Atoi(eventIDstring)
	if err != nil {
		http.Error(w, "Invalid event ID!", http.StatusBadRequest)
		return
	}

	errorsWeFound := ""
	messagesWeFound := ""

	if r.Method == http.MethodPost {
		//try to RSVP person
		r.ParseForm()

		email := strings.ToLower(r.PostFormValue("email"))
		if strings.HasSuffix(email, "@yale.edu") {
			err := addAttendee(eventID, email)
			messagesWeFound += "RSVP successful!"
			if err != nil {
				http.Error(w, "Invalid event ID!", http.StatusInternalServerError)
				return
			}
		} else {
			errorsWeFound += "Bad email address!"
		}
	}

	event, foundEvent := getEventByID(eventID)
	if foundEvent == false {
		http.Error(w, "No such event!", http.StatusNotFound)
		return
	}

	eventData := eventDetailContextData{
		Event:        event,
		FormErrors:   errorsWeFound,
		FormMessages: messagesWeFound,
	}

	tmpl["event-detail"].Execute(w, eventData)
}
