package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"thugcorp.io/nomado/data"
	"thugcorp.io/nomado/logger"

)

type PropertyHandler struct {
	Storage data.PropertyStorage
	Logger  *logger.Logger
}

func (h *PropertyHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.Logger.Error("Failed to encode JSON response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *PropertyHandler) GetTopProperties(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the logic to retrieve top properties
	properties, err := h.Storage.GetTopProperties()
	if err != nil {
		h.Logger.Errorf("Failed to get top properties: %v", err)
		http.Error(w, "Failed to retrieve top properties", http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, properties)
}

func (h *PropertyHandler) GetRandomProperties(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the logic to retrieve random properties
	properties, err := h.Storage.GetRandomProperties()
	if err != nil {
		h.Logger.Error("Failed to get random properties", err)
		http.Error(w, "Failed to retrieve random properties", http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, properties)
}

func (h *PropertyHandler) GetPropertyByID(w http.ResponseWriter, r *http.Request) {
	properties, err := h.Storage.GetPropertyByID(r.URL.Query().Get("id"))
	if err != nil {
		h.Logger.Errorf("Failed to get property by ID: %v", err)
		http.Error(w, "Failed to retrieve property", http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, properties)
}

func (h *PropertyHandler) SearchPropertiesByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name query parameter is required", http.StatusBadRequest)
		return
	}

	properties, err := h.Storage.SearchPropertiesByName(name)
	if err != nil {
		h.Logger.Errorf("Failed to search properties by name: %v", err)
		http.Error(w, "Failed to search properties", http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, properties)
}

func (h *PropertyHandler) GetPropertiesByLocation(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		http.Error(w, "Location query parameter is required", http.StatusBadRequest)
		return
	}

	properties, err := h.Storage.GetPropertiesByLocation(location)
	if err != nil {
		h.Logger.Errorf("Failed to get properties by location: %v", err)
		http.Error(w, "Failed to retrieve properties by location", http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, properties)
}
func (h *PropertyHandler) GetPropertiesByPriceRange(w http.ResponseWriter, r *http.Request) {
	minPriceStr := r.URL.Query().Get("minPrice")
	maxPriceStr := r.URL.Query().Get("maxPrice")
	if minPriceStr == "" || maxPriceStr == "" {
		http.Error(w, "minPrice and maxPrice query parameters are required", http.StatusBadRequest)
		return
	}

	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		http.Error(w, "Invalid minPrice value", http.StatusBadRequest)
		return
	}
	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		http.Error(w, "Invalid maxPrice value", http.StatusBadRequest)
		return
	}

	properties, err := h.Storage.GetPropertiesByPriceRange(minPrice, maxPrice)
	if err != nil {
		h.Logger.Errorf("Failed to get properties by price range: %v", err)
		http.Error(w, "Failed to retrieve properties by price range", http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, properties)
}


func (h *PropertyHandler) GetPropertiesByType(w http.ResponseWriter, r *http.Request){
	propertyType := r.URL.Query().Get("type")
	if propertyType == "" {
		http.Error(w, "Type query parameter is required", http.StatusBadRequest)
		return
	}

	properties, err := h.Storage.GetPropertiesByType(propertyType)
	if err != nil {
		h.Logger.Errorf("Failed to get properties by type: %v", err)
		http.Error(w, "Failed to retrieve properties by type", http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, properties)
}