package main

import (
	"net/http"
	"strings"
	//"io/ioutil"
	"time"
	// "github.com/jinzhu/now"
)

var searchEvents []Event
var filteredEvents []Event
var filterDate time.Time

type searchContextData struct {
	Events     []Event
	Query      string
	CatFilter  string
	DateFilter string
}

func searchController(w http.ResponseWriter, r *http.Request) {

	queries, hadQueries := r.URL.Query()["q"]
	if r.Method == http.MethodPost {
		//get data on event from create page

		r.ParseForm()
		category := r.PostFormValue("category")
		reset := r.PostFormValue("reset")
		date := r.PostFormValue("date")
		// fmt.Println("TEST")
		// fmt.Println(reset)

		// Reset button, set category to blank
		if reset == "reset" {
			category = ""
			date = ""
		}

		// This is getting the end of day. We will add days depending on the filter
		// to get the end date/time
		filterDate = time.Now()
		y, m, d := time.Now().Date()
		t := time.Date(y, m, d, 23, 59, 59, 99, filterDate.Location())

		switch date {
		case "Today":
			filterDate = time.Now()
			// if the report 2
		case "Tomorrow":
			filterDate = t.AddDate(0, 0, 1)
			// if the report 2
		case "ThisWeek":
			// if the report 2
			weekday := int(time.Now().Weekday())

			if weekday == 0 {
				filterDate = t
			} else {
				filterDate = t.AddDate(0, 0, (7 - weekday))
			}

		case "ThisWeekend":
			// if the report 2
			// filterDate := now.EndOfWeek()
		case "NextWeek":
			filterDate = t.AddDate(0, 0, 7)
		case "NextMonth":
			filterDate = t.AddDate(0, 1, 0)
		default:
			filterDate = time.Time{}
		}

		// Category and date search, if the date entered is zero, it just appends. if it's not zero, it has to be within the
		// time range specified
		filteredEvents = make([]Event, 0)
		for _, events := range searchEvents {
			if strings.Contains(strings.ToLower(events.Category), strings.ToLower(category)) {
				if filterDate.IsZero() == true || (events.Date.After(time.Now()) && events.Date.Before(filterDate)) {
					filteredEvents = append(filteredEvents, events)
				}
			}
		}

		contextData := searchContextData{
			Events:     filteredEvents,
			Query:      queries[0],
			CatFilter:  category,
			DateFilter: date,
		}
		tmpl["event-search"].Execute(w, contextData)

	} else {

		// If not category search, keyword search
		searchEvents = allEvents
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
			Events:     searchEvents,
			Query:      queries[0],
			CatFilter:  "",
			DateFilter: "",
		}
		tmpl["event-search"].Execute(w, contextData)
	}

}

func searchCategoryController(w http.ResponseWriter, r *http.Request) {
	queries, hadQueries := r.URL.Query()["q"]
	if r.Method == http.MethodPost {
		//get data on event from create page
		r.ParseForm()
		reset := r.PostFormValue("reset")
		date := r.PostFormValue("date")

		// Reset button, set category to blank
		if reset == "reset" {
			date = ""
		}
		// filterDate := ""
		filterDate = time.Now()
		y, m, d := time.Now().Date()
		t := time.Date(y, m, d, 23, 59, 59, 99, filterDate.Location())

		filterDate = time.Now()
		switch date {
		case "Today":
			filterDate = time.Now()
			// if the report 2
		case "Tomorrow":
			filterDate = t.AddDate(0, 0, 1)
			// if the report 2
		case "ThisWeekend":
			// if the report 2
			// filterDate := now.Sunday()
		case "ThisWeek":
			weekday := int(time.Now().Weekday())

			if weekday == 0 {
				filterDate = t
			} else {
				filterDate = t.AddDate(0, 0, (7 - weekday))
			}

		case "NextWeek":
			filterDate = t.AddDate(0, 0, 7)
		case "NextMonth":
			filterDate = t.AddDate(0, 1, 0)
		default:
			filterDate = time.Time{}
		}

		// if the date entered is zero, it just appends. if it's not zero, it has to be within the
		// time range specified
		filteredEvents = make([]Event, 0)
		for _, events := range searchEvents {
			if filterDate.IsZero() == true || (events.Date.After(time.Now()) && events.Date.Before(filterDate)) {
				filteredEvents = append(filteredEvents, events)
			}
		}

		contextData := searchContextData{
			Events:     filteredEvents,
			Query:      queries[0],
			CatFilter:  "",
			DateFilter: date,
		}

		tmpl["event-category"].Execute(w, contextData)

		// Separate controller for category link pages!
	} else {

		searchEvents = allEvents
		if hadQueries {
			searchEvents = make([]Event, 0)
			for _, events := range allEvents {
				if strings.Contains(strings.ToLower(events.Category), strings.ToLower(queries[0])) {
					searchEvents = append(searchEvents, events)
				}
			}
		}

		contextData := searchContextData{
			Events:     searchEvents,
			Query:      queries[0],
			CatFilter:  "",
			DateFilter: "",
		}

		tmpl["event-category"].Execute(w, contextData)
	}
}
