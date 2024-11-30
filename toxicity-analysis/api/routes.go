package api

import (
	"net/http"
	"toxicity-analysis/service"
)

func NewRouter(service *service.AnalyzerService) *http.ServeMux {
	handlers := NewHandlers(service)
	mux := http.NewServeMux()
	mux.HandleFunc("/analyze", handlers.AnalyzeToxicityHandler)
	return mux
}
