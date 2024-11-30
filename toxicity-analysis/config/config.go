package config

type Config struct {
	OllamaBaseURL string
	Port          string
}

func LoadConfig() *Config {
	return &Config{
		OllamaBaseURL: "http://localhost:8991",
		Port:          "8080",
	}
}
