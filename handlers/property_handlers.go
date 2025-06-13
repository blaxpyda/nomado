package handlers

import (
	"encoding/json"
	"net/http"

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
		h.Logger.Error("Failed to get top properties ", err)
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
