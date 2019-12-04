package main

import (
	"encoding/json"
	"net/http"
	//"io/ioutil"
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
