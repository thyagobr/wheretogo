package handlers

import (
	"net/http"
	
	"github.com/thyagobr/wheretogo/internal/models"
	"github.com/thyagobr/wheretogo/internal/db"
	"github.com/thyagobr/wheretogo/internal/dtos"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	result := db.DB.Preload("Place").Preload("Place.Tags").Find(&events)
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
