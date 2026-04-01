package dtos

import "github.com/thyagobr/wheretogo/internal/models"

type PlaceResponse struct {
	ID      uint          `json:"id"`
	Name    string        `json:"name"`
	Address string        `json:"address"`
	Country string        `json:"country"`
	City    string        `json:"city"`
	Tags    []TagResponse `json:"tags"`
}

type PlacesResponse struct {
	Places []PlaceResponse `json:"places"`
	Place  *PlaceResponse  `json:"place,omitempty"`
}

type UpdatePlaceRequest struct {
	Name    *string             `json:"name,omitempty"`
	Address *string             `json:"address,omitempty"`
	Country *string             `json:"country,omitempty"`
	City    *string             `json:"city,omitempty"`
	Tags    *[]UpdateTagRequest `json:"tags,omitempty"`
}

type UpdateTagRequest struct {
	Text *string `json:"text,omitempty"`
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
