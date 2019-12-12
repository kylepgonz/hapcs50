package main

import (
	"github.com/go-chi/chi"
)

func getRoutes() chi.Router {
	// We're using chi as the router. You'll want to read
	// the documentation https://github.com/go-chi/chi
	// so that you can capture parameters like /events/5
	// or /api/events/4 -- where you want to get the
	// event id (5 and 4, respectively).

	r := chi.NewRouter()
	r.Get("/", indexController)
	r.Get("/about", aboutHandler)
	r.Get("/events/{eventID:[0-9]+}", eventDetailController)
	r.Get("/events/{eventID:[0-9]+}/{action:[a-z-]+}", eventDetailController)
	r.Post("/events/{eventID:[0-9]+}", eventDetailController)
	r.Get("/events/new", eventCreateController)
	r.Get("/api/events", apiEventListController)
	r.Get("/api/events/{eventID:[0-9]+}", apiEventController)
	r.Post("/events/new", eventCreateController)
	r.Get("/search", searchController)
	r.Post("/search", searchController)
	r.Get("/search/category", searchCategoryController)
	r.Post("/search/category", searchCategoryController)

	addStaticFileServer(r, "/static/", "staticfiles")
	return r
}
