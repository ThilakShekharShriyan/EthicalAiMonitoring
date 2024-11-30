package api

import (
	"encoding/json"
	"net/http"
	"toxicity-analysis/domain"
	"toxicity-analysis/service"
)

type Handlers struct {
	Service *service.AnalyzerService
}

func NewHandlers(service *service.AnalyzerService) *Handlers {
	return &Handlers{Service: service}
}

func (h *Handlers) AnalyzeToxicityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.ToxicityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	resp, err := h.Service.AnalyzeText(req.Text)
	if err != nil {
		http.Error(w, "Failed to analyze text", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
