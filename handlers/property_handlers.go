package handlers

import (
	"encoding/json"
	"net/http"

	"thugcorp.io/nomado/models"
)

type PropertyHandler struct {
	// TODO
}

func (h *PropertyHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *PropertyHandler) GetTopProperties(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the logic to retrieve top properties
	properties := []models.Property{
		{ID: "1", Name: "Luxury Villa", Price: 1500000, Location: "Beverly Hills"},
		{ID: "2", Name: "Modern Apartment", Price: 800000, Location: "New York"},
		{ID: "3", Name: "Cozy Cottage", Price: 300000, Location: "Lake Tahoe"},
	}
	h.writeJSONResponse(w, properties)
}