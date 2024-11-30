package domain

// ToxicityRequest represents the input from the user
type ToxicityRequest struct {
	prompt string `json:"text"`
	promptResponse string `json:"text"`
}

// ToxicityResponse represents the API response
type ToxicityResponse struct {
	ToxicityScore float64 `json:"toxicity_score"`
	Label         string  `json:"label"`
}
