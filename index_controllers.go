package main

import (
	"net/http"
	"time"
)

func indexController(w http.ResponseWriter, r *http.Request) {

	type indexContextData struct {
		Events []Event
		Today  time.Time
	}

	contextData := indexContextData{
		Events: allEvents,
		Today:  time.Now(),
	}

	tmpl["index"].Execute(w, contextData)
}
