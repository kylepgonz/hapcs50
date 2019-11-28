package main

import (
	"net/http"
	"strings"
	//"io/ioutil"
)

var searchEvents []Event
var filteredEvents []Event

type searchContextData struct {
	Events []Event
	Query  string
}

func searchController(w http.ResponseWriter, r *http.Request) {
	// errorsWeFound := ""
	// messagesWeFound := ""
	queries, hadQueries := r.URL.Query()["q"]
	if r.Method == http.MethodPost {
		//get data on event from create page
		// fmt.Fprintf(w, "pwede")
		r.ParseForm()
		category := r.PostFormValue("category")
		filteredEvents = make([]Event, 0)
		for _, events := range searchEvents {
			if strings.Contains(strings.ToLower(events.Category), strings.ToLower(category)) {
				filteredEvents = append(filteredEvents, events)
			}
		}
		contextData := searchContextData{
			Events: filteredEvents,
			Query:  queries[0],
		}
		tmpl["event-search"].Execute(w, contextData)
	} else {

		searchEvents = allEvents
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
	}

}

func searchCategoryController(w http.ResponseWriter, r *http.Request) {

	searchEvents = allEvents
	queries, hadQueries := r.URL.Query()["q"]

	if hadQueries {
		searchEvents = make([]Event, 0)
		for _, events := range allEvents {
			if strings.Contains(strings.ToLower(events.Category), strings.ToLower(queries[0])) {
				searchEvents = append(searchEvents, events)
			}
		}
	}

	contextData := searchContextData{
		Events: searchEvents,
		Query:  queries[0],
	}

	tmpl["event-category"].Execute(w, contextData)
}
