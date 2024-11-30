package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"toxicity-analysis/domain"
)

type OllamaClient struct {
	BaseURL string
	Client  *http.Client
}

func NewOllamaClient(baseURL string) *OllamaClient {
	return &OllamaClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (o *OllamaClient) AnalyzeToxicity(prompt string) (*domain.ToxicityResponse, error) {
	requestBody, _ := json.Marshal(domain.ToxicityRequest{Text: prompt})
	resp, err := o.Client.Post(o.BaseURL+"/analyse", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("ollama API returned an error")
	}

	var result domain.ToxicityResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
