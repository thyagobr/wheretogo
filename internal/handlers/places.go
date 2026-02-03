package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/thyagobr/wheretogo/internal/models"
  "github.com/thyagobr/wheretogo/internal/db"
	"github.com/thyagobr/wheretogo/internal/dtos"
)

type ApiResponse[T any] struct {
	Data T `json:"data"`
}

type PlacesResponse struct {
	Places []dtos.PlaceResponse `json:"places"`
}

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

	apiResp := ApiResponse[PlacesResponse]{
		Data: PlacesResponse{
			Places: placeResponses,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(apiResp)
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
	result := db.DB.First(&place, id)
	if result.Error != nil {
		http.Error(w, "Place not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(place)
}

type CreatePlaceRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Country string `json:"country"`
	City    string `json:"city"`
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
	}

	result := db.DB.Create(&place)
	if result.Error != nil {
		http.Error(w, "Failed to create place", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(place)
}
