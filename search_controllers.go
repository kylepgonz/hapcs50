package main

import (
	"net/http"
	"strings"
	//"io/ioutil"
)

var searchEvents []Event

type searchContextData struct {
	Events []Event
	Query  string
}

func searchController(w http.ResponseWriter, r *http.Request) {
	// errorsWeFound := ""
	// messagesWeFound := ""
	searchEvents = allEvents
	queries, hadQueries := r.URL.Query()["q"]
	// if foundPages {
	// 	fmt.Fprintf(w, pages[0])
	// }
	if hadQueries {
		searchEvents = make([]Event, 0)
		for _, events := range allEvents {
			if strings.Contains(strings.ToLower(events.Title), strings.ToLower(queries[0])) {
				searchEvents = append(searchEvents, events)
			} else if strings.Contains(strings.ToLower(events.Location), strings.ToLower(queries[0])) {
				searchEvents = append(searchEvents, events)
			}
		}
	}

	contextData := searchContextData{
		Events: searchEvents,
		Query:  queries[0],
	}

	tmpl["event-search"].Execute(w, contextData)

	// fmt.Fprintf("%v", searchEvents)
	// if r.Method == http.MethodPost {
	// 	//get data on event from create page
	// 	r.ParseForm()
	// 	title := strings.ToLower(r.PostFormValue("title"))
	// 	location := strings.ToLower(r.PostFormValue("location"))
	// 	image := r.PostFormValue("image")
	// 	date, _ := time.Parse("2006-01-02T15:04", r.PostFormValue("date"))
	// 	category := r.PostFormValue("category")
	// 	description := r.PostFormValue("description")
	// 	id := getMaxEventID() + 1
	//
	// 	today := time.Now()
	// 	if today.After(date) {
	// 		errorsWeFound += "Wrong date!"
	// 	} else if id == -1 {
	// 		errorsWeFound += "No more event ids!"
	// 	} else if categoryExists[category] == false {
	// 		// println(category)
	// 		// println("test")
	// 		errorsWeFound += "Wrong category!"
	// 	} else {
	// 		newEvent := Event{
	// 			ID:          id,
	// 			Title:       title,
	// 			Date:        date,
	// 			Image:       image,
	// 			Location:    location,
	// 			Category:    category,
	// 			Description: description,
	// 			Attending:   []string{},
	// 		}
	// 		addEvent(newEvent)
	// 		_, foundEvent := getEventByID(id)
	// 		if foundEvent == false {
	// 			http.Error(w, "Error in event creation!", http.StatusInternalServerError)
	// 			return
	// 		}
	// 		messagesWeFound += "Event creation successful!"
	// 	}
	// }
	// eventData := eventCreateContextData{
	// 	FormErrors:   errorsWeFound,
	// 	FormMessages: messagesWeFound,
	// }

}
