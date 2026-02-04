package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/thyagobr/wheretogo/internal/db"
	"github.com/thyagobr/wheretogo/internal/dtos"
	"github.com/thyagobr/wheretogo/internal/middlewares"
	"github.com/thyagobr/wheretogo/internal/models"
	"github.com/thyagobr/wheretogo/internal/utils"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	result := db.DB.Preload("Place").Preload("Place.Tags").Where("start_at > ?", time.Now()).Find(&events)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
		return
	}

	eventResponses := make([]dtos.EventResponse, len(events))
	for i, event := range events {
		eventResponses[i] = dtos.ToEventResponse(event)
	}

	apiResp := ApiResponse[dtos.EventsResponse]{
		Data: dtos.EventsResponse{
			Events: eventResponses,
		},
	}

	respondJson(w, http.StatusOK, apiResp)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if (err != nil) || (id == 0) {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	result := db.DB.Preload("Place").Preload("Place.Tags").First(&event, id)
	if result.Error != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	eventResponse := dtos.ToEventResponse(event)

	apiResp := ApiResponse[dtos.EventsResponse]{
		Data: dtos.EventsResponse{
			Event: &eventResponse,
		},
	}

	respondJson(w, http.StatusOK, apiResp)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var eventParams dtos.CreateEventRequest
	err := utils.DecodeJSON(r, &eventParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Make sure the place exists
	placeID := eventParams.PlaceID
	var place models.Place
	resultPlace := db.DB.First(&place, placeID)
	if resultPlace.Error != nil {
		http.Error(w, "Place not found", http.StatusBadRequest)
		return
	}

	user, ok := middlewares.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	event := models.Event{
		Name:        eventParams.Name,
		Description: eventParams.Description,
		StartsAt:    eventParams.StartsAt,
		EndsAt:      eventParams.EndsAt,
		PlaceID:     uint(placeID),
		Public:      eventParams.Public,
		UserID:      uint(user.ID),
	}

	result := db.DB.Create(&event)
	if result.Error != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	eventResponse := dtos.ToEventResponse(event)

	apiResp := ApiResponse[dtos.EventsResponse]{
		Data: dtos.EventsResponse{
			Event: &eventResponse,
		},
	}

	respondJson(w, http.StatusCreated, apiResp)
}
