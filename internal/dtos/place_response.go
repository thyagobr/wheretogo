package dtos

import "github.com/thyagobr/wheretogo/internal/models"

type PlaceResponse struct {
	ID      uint        	`json:"id"`
	Name    string      	`json:"name"`
	Address string 				`json:"address"`
	Country string 				`json:"country"`
	City    string 				`json:"city"`
	Tags 		[]TagResponse `json:"tags"`
}

func ToPlaceResponse(place models.Place) PlaceResponse {
	tags := make([]TagResponse, len(place.Tags))
	for i, tag := range place.Tags {
		tags[i] = TagResponse{
			ID:   tag.ID,
			Text: tag.Text,
		}
	}

	return PlaceResponse{
		ID:      place.ID,
		Name:    place.Name,
		Address: place.Address,
		Country: place.Country,
		City:    place.City,
		Tags:    tags,
	}
}
