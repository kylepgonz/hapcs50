package main

import (
	"net/http"
	"strconv"
	"time"
	//"io/ioutil"
	"math/rand"
)

type indexContextData struct {
	Events    []Event
	AllEvents []Event
	Today     time.Time
}

var randomEvents []Event

func indexController(w http.ResponseWriter, r *http.Request) {
	randomEvents = allEvents
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allEvents), func(i, j int) { randomEvents[i], randomEvents[j] = randomEvents[j], randomEvents[i] })
	var eventsList = randomEvents[0:6]
	contextData := indexContextData{
		Events:    eventsList,
		AllEvents: allEvents,
		Today:     time.Now(),
	}

	tmpl["index"].Execute(w, contextData)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl["about"].Execute(w, "")
}

func eventsController(w http.ResponseWriter, r *http.Request) {

	contextData := indexContextData{
		Events: allEvents,
		Today:  time.Now(),
	}

	tmpl["index"].Execute(w, contextData)
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	pages, foundPages := r.URL.Query()["report"]

	if foundPages {
		i, err := strconv.Atoi(pages[0])
		if err == nil {
			switch i {
			case 2:
				// if the report 2
				tmpl["report2"].Execute(w, "")
			case 3:
				// if the report 2
				tmpl["report3"].Execute(w, "")
			default:
				// if anthing e;slse
			}
		}
	} else {
		//if there is no extra number
		tmpl["report1"].Execute(w, "")
	}

}
