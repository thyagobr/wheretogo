package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/thyagobr/wheretogo/internal/models"
  "github.com/thyagobr/wheretogo/internal/db"
	"github.com/thyagobr/wheretogo/internal/dtos"
	"github.com/thyagobr/wheretogo/internal/clients"
)

func GetPlaces(w http.ResponseWriter, r *http.Request) { var places []models.Place
	result := db.DB.Preload("Tags").Find(&places)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve places", http.StatusInternalServerError)
		return
	}

	placeResponses := make([]dtos.PlaceResponse, len(places))
	for i, place := range places {
		placeResponses[i] = dtos.ToPlaceResponse(place)
	}
	apiResp := ApiResponse[dtos.PlacesResponse]{
		Data: dtos.PlacesResponse{
			Places: placeResponses,
		},
	}

	respondJson(w, http.StatusOK, apiResp)
}

func GetPlace(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var place models.Place
	result := db.DB.Preload("Tags").First(&place, id)
	if result.Error != nil {
		http.Error(w, "Place not found", http.StatusNotFound)
		return
	}

	placeResponse := dtos.ToPlaceResponse(place)

	apiResp := ApiResponse[dtos.PlacesResponse]{
		Data: dtos.PlacesResponse{
			Place: &placeResponse,
		},
	}

	respondJson(w, http.StatusOK, apiResp)
}

func GetPlaceEvents(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var events []models.Event
	result := db.DB.Where("place_id = ?", id).Find(&events)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
		return
	}

	eventResponses := make([]dtos.EventResponse, len(events))
	for i, event := range events {
		eventResponses[i] = dtos.ToEventResponse(event)
	}

	apiResp := ApiResponse[dtos.EventsResponse] {
		Data: dtos.EventsResponse {
			Events: eventResponses,
		},
	}

	respondJson(w, http.StatusOK, apiResp)
}

type CreatePlaceRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Country string `json:"country"`
	City    string `json:"city"`
	Tags    []CreateTagRequest `json:"tags"`
}

type CreateTagRequest struct {
	Text string `json:"text"`
}

func CreatePlace(w http.ResponseWriter, r *http.Request) {
	var createPlaceReq CreatePlaceRequest
	err := json.NewDecoder(r.Body).Decode(&createPlaceReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	place := models.Place{
		Name:    createPlaceReq.Name,
		Address: createPlaceReq.Address,
		Country: createPlaceReq.Country,
		City:    createPlaceReq.City,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := db.DB.Create(&place)
	if result.Error != nil {
		http.Error(w, "Failed to create place", http.StatusInternalServerError)
		return
	}

	tags := make([]models.Tag, len(createPlaceReq.Tags))
	for i, tag := range createPlaceReq.Tags {
		tags[i] = models.Tag{
			Text:    			tag.Text,
			TaggableID: 	place.ID,
			TaggableType: "Place",
		}
	}
	if len(tags) > 0 {
		result = db.DB.Create(&tags)
		if result.Error != nil {
			http.Error(w, "Failed to create tags", http.StatusInternalServerError)
			return
		}
	}

	var createdPlace models.Place
	db.DB.Preload("Tags").First(&createdPlace, place.ID)

	placeResponse := dtos.ToPlaceResponse(createdPlace)

	apiResp := ApiResponse[dtos.PlacesResponse]{
		Data: dtos.PlacesResponse{
			Place: &placeResponse,
		},
	}

	respondJson(w, http.StatusCreated, apiResp)
}

// params: map[city:[berlin] country:[germany] limit:[1] name:[Wohzimmer]]
func SearchAddress(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	searchReq := dtos.SearchAddressRequest{
		Name:    queryParams.Get("name"),
		City:    queryParams.Get("city"),
		Country: queryParams.Get("country"),
		Limit:   10,
	}

	openMapsClient := clients.NewOpenMapsClient()
	openMapsResp, err := openMapsClient.SearchAddress(searchReq)
	if err != nil {
		http.Error(w, "Failed to search address", http.StatusInternalServerError)
		return
	}

	apiResp := ApiResponse[[]clients.OpenMapsResponse]{
		Data: openMapsResp,
	}

	respondJson(w, http.StatusOK, apiResp)
}
