package main

import (
	"net/http"
	"time"
	"strconv"
	//"io/ioutil"
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
	pageContent := []string{}
	reportDate := ""
	pages, foundPages := r.URL.Query()["report"]

	if foundPages {
		i, err := strconv.Atoi(pages[0])
		if err == nil {
				switch i {
				case 2:
					pageContent = append(pageContent,
													"https://www.manchesteryz.org/wp-content/uploads/2018/10/UNDER-CONSTRUCTION.jpg",)
													//"https://live.staticflickr.com/65535/48992791498_8928daea1c_b.jpg",
													//"https://live.staticflickr.com/65535/48993530992_35de02a005_b.jpg",
													//"https://live.staticflickr.com/65535/48992791483_45b8c0b78b_b.jpg")
					reportDate = "November 07, 2019"
				default:
					pageContent = append(pageContent,
													"https://www.manchesteryz.org/wp-content/uploads/2018/10/UNDER-CONSTRUCTION.jpg")
				}
			}
	} else {
			pageContent = append(pageContent,
											"https://live.staticflickr.com/65535/48993530967_dfac43e066_b.jpg",
											"https://live.staticflickr.com/65535/48992791498_8928daea1c_b.jpg",
											"https://live.staticflickr.com/65535/48993530992_35de02a005_b.jpg",
											"https://live.staticflickr.com/65535/48992791483_45b8c0b78b_b.jpg")
			reportDate = "October 31, 2019"
	}

	type indexContextData2 struct {
		Report []string
		Content  string
	}
	contextData2 := indexContextData2{
		Report: pageContent,
		Content:  reportDate,
	}

	tmpl["layout"].Execute(w, contextData2)
}
