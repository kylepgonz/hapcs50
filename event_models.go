package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

var allEvents []Event

// A "mutex" is used to say "hey, I'm using these data,
// don't touch them while I'm using them!" We use the
// mutex when adding events or attendees.
var allEventsMutex = &sync.Mutex{}
var categoryExists = map[string]bool{
	"Sport":         true,
	"Career + Biz":  true,
	"Food + Drink":  true,
	"Music + Dance": true,
	"Spiritual":     true,
	"Art":           true,
	"Charity":       true,
}

// Event - encapsulates information about an event
type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Location    string    `json:"location"`
	Image       string    `json:"image"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Attending   []string  `json:"attending"`
}

// getEventByID - returns the event in `allEvents` that has
// the specified id and a boolean indicating whether or not
// it was found. If it is not found, returns an empty event
// and false.
func getEventByID(id int) (Event, bool) {
	for _, event := range allEvents {
		if event.ID == id {
			return event, true
		}
	}
	return Event{}, false
}

// getMaxEventID returns the maximum of all
// the ids of the events in allEvents
func getMaxEventID() int {
	maxID := -1
	for _, event := range allEvents {
		if event.ID > maxID {
			maxID = event.ID
		}
	}
	return maxID
}

// Adds an attendee to an event
func addAttendee(id int, email string) error {
	// When get a "lock", we are saying that we're
	// going to edit some data and we don't want
	// anybody else to use those data while we're
	// writing it. Recall, in go there might be
	// multiple threads (goroutines).
	allEventsMutex.Lock()
	defer allEventsMutex.Unlock()
	for i := 0; i < len(allEvents); i++ {
		if allEvents[i].ID == id {
			allEvents[i].Attending = append(allEvents[i].Attending, email)
			log.Println(allEvents[i].Attending)
			return nil
		}
	}
	return errors.New("No such event")
}

func addEvent(event Event) {
	allEventsMutex.Lock()
	event.ID = getMaxEventID() + 1
	allEvents = append(allEvents, event)
	allEventsMutex.Unlock()
}

// init is run once when this file is first loaded. See
// https://golang.org/doc/effective_go.html#init
// https://medium.com/golangspec/init-functions-in-go-eac191b3860a
func init() {
	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic("Could not load timezone database on your system!")
	}

	defaultEvents := []Event{
		Event{
			ID:          1,
			Title:       "SOM House Party",
			Date:        time.Date(2019, 10, 17, 16, 30, 0, 0, newYork),
			Image:       "http://i.imgur.com/pXjrQ.gif",
			Location:    "Kyle's house",
			Category:    "",
			Attending:   []string{"kyle.jensen@yale.edu", "kim.kardashian@yale.edu"},
			Description: "Something fun!",
		},
		Event{
			ID:        2,
			Title:     "BBQ party for hackers and nerds",
			Date:      time.Date(2019, 10, 19, 19, 0, 0, 0, newYork),
			Image:     "http://i.imgur.com/7pe2k.gif",
			Location:  "Sharon Oster's house",
			Category:  "",
			Attending: []string{"kyle.jensen@yale.edu", "kim.kardashian@yale.edu"},
		},
		Event{
			ID:        3,
			Title:     "BBQ for managers",
			Date:      time.Date(2019, 12, 2, 18, 0, 0, 0, newYork),
			Image:     "http://i.imgur.com/CJLrRqh.gif",
			Location:  "Barry Nalebuff's house",
			Category:  "",
			Attending: []string{"kim.kardashian@yale.edu"},
		},
		Event{
			ID:        4,
			Title:     "Cooking lessons for the busy business student",
			Date:      time.Date(2019, 12, 21, 19, 0, 0, 0, newYork),
			Image:     "http://i.imgur.com/02KT9.gif",
			Location:  "Yale Farm",
			Category:  "",
			Attending: []string{"homer.simpson@yale.edu"},
		},
		Event{
			ID:          5,
			Title:       "Singing and Giving",
			Date:        time.Date(2019, 12, 21, 19, 0, 0, 0, newYork),
			Image:       "https://media.giphy.com/media/JD0PRlEsWNkeQ/giphy.gif",
			Location:    "Yale School of Music",
			Category:    "",
			Attending:   []string{"homer.simpson@yale.edu"},
			Description: "Attend a chorale presentation from Yale School of music and proceeds will be donated to charities in New Haven",
		},
		Event{
			ID:          6,
			Title:       "Crack the Case",
			Date:        time.Date(2019, 12, 21, 19, 0, 0, 0, newYork),
			Image:       "https://i.imgur.com/VzV1mIV.gif",
			Location:    "SOM Beinecke",
			Category:    "Career + Business",
			Attending:   []string{"homer.simpson@yale.edu"},
			Description: "Learn to case like an MBB consultant. This crash course will get you in tip top casing shape in no time!",
		},
		Event{
			ID:          7,
			Title:       "Draw a Dino at the Dino-snore",
			Date:        time.Date(2019, 12, 21, 19, 0, 0, 0, newYork),
			Image:       "https://media.giphy.com/media/3oEhn6803hZKJNeMhy/giphy.gif",
			Location:    "Peabody Museum",
			Category:    "Art",
			Attending:   []string{"homer.simpson@yale.edu"},
			Description: "Join us for a family sleepover at the Peabody. This night will be filled with artsy inspiration of our prehistoric predators. Activities include: Dino face painting, Brontosaurus sketching and paleo cooking classes.",
		},
		Event{
			ID:          8,
			Title:       "Meditate and Chill",
			Date:        time.Date(2019, 12, 21, 19, 0, 0, 0, newYork),
			Image:       "https://media.giphy.com/media/3oKGzC42QlXAnjijHa/giphy.gif",
			Location:    "Yale Farm",
			Category:    "Spiritual",
			Attending:   []string{"homer.simpson@yale.edu"},
			Description: "If you want to learn effective ways to meditate, this event is for you.",
		},
		Event{
			ID:          9,
			Title:       "Jiggle and Jive",
			Date:        time.Date(2019, 12, 21, 19, 0, 0, 0, newYork),
			Image:       "http://giphygifs.s3.amazonaws.com/media/10gZNwuUuer5aU/giphy.gif",
			Location:    "Divinity School",
			Category:    "Music + dance",
			Attending:   []string{"homer.simpson@yale.edu"},
			Description: "If you like to jiggle what your moma gave ya, don’t miss this pow-wow! We’ll enjoy some live tunes to booty shake to, courtesy of the New Haven Jive Tribe",
		},
		Event{
			ID:          10,
			Title:       "Turkey Basting Bonanza",
			Date:        time.Date(2019, 12, 21, 19, 0, 0, 0, newYork),
			Image:       "http://giphygifs.s3.amazonaws.com/media/JBP8eDB1rhIv6/giphy.gif",
			Location:    "SOM Quad",
			Category:    "Food + drink",
			Attending:   []string{"homer.simpson@yale.edu"},
			Description: "Bring your baster to the business school for this big bird cookoff!",
		},
		Event{
			ID:          11,
			Title:       "Mutt Strut",
			Date:        time.Date(2019, 12, 21, 19, 0, 0, 0, newYork),
			Image:       "https://media.giphy.com/media/cLcxtL1z8t8oo/giphy.gif",
			Location:    "East Rock Park",
			Category:    "Sport",
			Attending:   []string{"homer.simpson@yale.edu"},
			Description: "Join us for the 5 mile mutt strut. Just bring your pooch & pozzie  to the park for playful parade.",
		},
	}
	allEvents = append(allEvents, defaultEvents...)
}
