package service

import (
	"toxicity-analysis/client"
	"toxicity-analysis/domain"

)

type AnalyzerService struct {
	OllamaClient *client.OllamaClient
}

func NewAnalyzerService(client *client.OllamaClient) *AnalyzerService {
	return &AnalyzerService{OllamaClient: client}
}

func (a *AnalyzerService) AnalyzeText(text string) (*domain.ToxicityResponse, error) {

	return a.OllamaClient.AnalyzeToxicity(text)
}
