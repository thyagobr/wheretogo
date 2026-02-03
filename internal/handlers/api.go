package handlers

import (
	"encoding/json"
	"net/http"
)

type ApiResponse[T any] struct {
	Data T `json:"data"`
}

func respondJson(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

