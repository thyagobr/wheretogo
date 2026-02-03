package dtos

import "github.com/thyagobr/wheretogo/internal/models"
import "time"

func ToEventResponse(event models.Event) EventResponse {
	return EventResponse{
		ID:          event.ID,
		Name:        event.Name,
		StartsAt:    event.StartsAt,
		EndsAt:      event.EndsAt,
		Description: event.Description,
		Public:      event.Public,
		Place:       event.Place,
	}
}

type EventResponse struct {
	ID          uint   `json:"id"`
	Name				string `json:"name"`
	StartsAt    time.Time `json:"startsAt"`
	EndsAt      *time.Time `json:"endsAt"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Place       models.Place `json:"place"`
}

type EventsResponse struct {
	Event *EventResponse   `json:"event,omitempty"`
	Events []EventResponse `json:"events,omitempty"`
}
