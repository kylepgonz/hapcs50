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
	FormErrors   []string
	FormMessages string
	Redirect     []int
}

func eventCreateController(w http.ResponseWriter, r *http.Request) {
	errorsWeFound := make([]string, 0)
	email := make([]string, 0)
	messagesWeFound := ""
	redirect := make([]int, 0)
	id := 0

	if r.Method == http.MethodPost {
		//get data on event from create page
		r.ParseForm()
		title := strings.ToLower(r.PostFormValue("title"))
		location := strings.ToLower(r.PostFormValue("location"))
		image := r.PostFormValue("image")
		email = append(email, r.PostFormValue("useremail"))
		date, err := time.Parse("2006-01-02T15:04", r.PostFormValue("date"))
		category := r.PostFormValue("category")
		description := r.PostFormValue("description")
		id = getMaxEventID() + 1

		today := time.Now()
		if today.After(date) {
			errorsWeFound = append(errorsWeFound, "Wrong date!")
		}
		if err != nil {
			errorsWeFound = append(errorsWeFound, "Wrong date format!")
		}
		// if strings.HasSuffix(email[0], "@yale.edu") == false {
		// 	errorsWeFound = append(errorsWeFound, "@yale.edu only!")
		// }
		if id == -1 {
			errorsWeFound = append(errorsWeFound, "No more event ids!")
		}
		if categoryExists[category] == false {
			// println(category)
			// println("test")
			errorsWeFound = append(errorsWeFound, "Wrong category!")
		}
		if len(title) < 5 || len(title) > 50 {
			errorsWeFound = append(errorsWeFound, "Bad Title! (5-50 Chars only)")
		}
		if len(location) < 5 || len(location) > 50 {
			errorsWeFound = append(errorsWeFound, "Bad Location! (5-50 Chars only)")
		}
		if len(image) != 0 {
			imageFileType := image[(len(image) - 3):]
			if gifTypeAllowed[imageFileType] == false {
				errorsWeFound = append(errorsWeFound, "Bad Image Type! '.png', '.jpg', '.jpeg', '.gif', '.gifv' only")
			}
		} else {
			errorsWeFound = append(errorsWeFound, "Bad Image Type! '.png', '.jpg', '.jpeg', '.gif', '.gifv' only")
		}

		if len(errorsWeFound) == 0 {
			newEvent := Event{
				ID:          id,
				Title:       title,
				Date:        date,
				Image:       image,
				Location:    location,
				Category:    category,
				Description: description,
				Attending:   email,
			}
			addEvent(newEvent)
			_, foundEvent := getEventByID(id)
			if foundEvent == false {
				http.Error(w, "Error in event creation!", http.StatusInternalServerError)
				return
			}
			messagesWeFound += "Event has been successfully created! Redirecting you..."
			redirect = append(redirect, id)
		}
	}
	eventData := eventCreateContextData{
		FormErrors:   errorsWeFound,
		FormMessages: messagesWeFound,
		Redirect:     redirect,
	}
	if messagesWeFound != "" {
		foo := "/events/" + strconv.Itoa(id)
		println(foo)
		time.Sleep(2 * time.Second)
		http.Redirect(w, r, foo, http.StatusFound)
		return
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
	confCode := ""
	if r.Method == http.MethodPost {
		r.ParseForm()

		email := strings.ToLower(r.PostFormValue("email"))
		if strings.HasSuffix(email, "@yale.edu") {
			err := addAttendee(eventID, email)
			if err != nil {
				http.Error(w, "Invalid event ID!", http.StatusInternalServerError)
				return
			}
			confCode = getConfirmationCode(email)
		} else {
			errorsWeFound += "Bad email address! '@yale.edu only!'"
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
		FormMessages: confCode,
	}

	tmpl["event-detail"].Execute(w, eventData)
}
