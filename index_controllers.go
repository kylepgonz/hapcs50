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

var randomEvents []int

var eventsList []Event

func indexController(w http.ResponseWriter, r *http.Request) {

	eventsList = make([]Event, 6)
	// fmt.Println(allEvents)
	randomEvents = make([]int, len(allEvents))
	// fmt.Print(randomEvents)
	for i := 0; i < len(allEvents); i++ {
		randomEvents[i] = allEvents[i].ID
	}
	// fmt.Println(randomEvents)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randomEvents), func(i, j int) { randomEvents[i], randomEvents[j] = randomEvents[j], randomEvents[i] })
	// fmt.Println(randomEvents)
	// eventsList := randomEvents[0:6]

	for i := 1; i <= 6; i++ {
		enterEvent, _ := getEventByID(randomEvents[i])
		eventsList[i-1] = enterEvent
	}

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
	pages, foundPages := r.URL.Query()["sprint"]

	if foundPages {
		i, err := strconv.Atoi(pages[0])
		if err == nil {
			switch i {
			case 1:
				// if the report 2
				tmpl["report1"].Execute(w, "")
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
	}

}
