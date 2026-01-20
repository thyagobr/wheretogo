package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thyagobr/wheretogo/internal/models"
)

func GetPlaces(w http.ResponseWriter, r *http.Request) {
	places := []models.Place{
		{
			ID:      1,
			Name:    "Eiffel Tower",
			Address: "Champ de Mars, 5 Avenue Anatole France, 75007 Paris",
			Country: "France",
			City:    "Paris",
			//Description: "An iconic symbol of France, the Eiffel Tower is a wrought-iron lattice tower on the Champ de Mars in Paris.",
		},
		{
			ID:      2,
			Name:    "Statue of Liberty",
			Address: "Liberty Island, New York, NY 10004",
			Country: "USA",
			City:    "New York",
			//Description: "A gift from France to the United States, the Statue of Liberty is a symbol of freedom and democracy.",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(places)
}
