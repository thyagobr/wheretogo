package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/thyagobr/wheretogo/internal/db"
	"github.com/thyagobr/wheretogo/internal/dtos"
	"github.com/thyagobr/wheretogo/internal/models"
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
