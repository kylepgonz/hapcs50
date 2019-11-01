package main

import (
	"net/http"
	"time"
)

type indexContextData struct {
	Events []Event
	Today  time.Time
}

func indexController(w http.ResponseWriter, r *http.Request) {

	contextData := indexContextData{
		Events: allEvents,
		Today:  time.Now(),
	}

	tmpl["index"].Execute(w, contextData)
}

func reportHandler(w http.ResponseWriter, r *http.Request) {

	type indexContextData2 struct {
		Title string
		Content  string
	}
	contextData2 := indexContextData2{
		Title: "Report",
		Content:  "Here's our report!",
	}

	tmpl["layout"].Execute(w, contextData2)
}
