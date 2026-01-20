package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thyagobr/wheretogo/internal/models"
  "github.com/thyagobr/wheretogo/internal/db"
)

func GetPlaces(w http.ResponseWriter, r *http.Request) {
	var places []models.Place

	result := db.DB.Find(&places)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve places", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(places)
}
