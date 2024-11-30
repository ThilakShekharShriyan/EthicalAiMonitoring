package main

import (
	"log"
	"net/http"
	"toxicity-analysis/api"
	"toxicity-analysis/client"
	"toxicity-analysis/config"
	"toxicity-analysis/service"
	//"fmt"
)

func main(){

	//Fetches the config from config.go
	cfg := config.LoadConfig()

	//fmt.Println(cfg.OllamaBaseURL,cfg.Port)

	//Creates an OllamaClient
	ollamaClient := client.NewOllamaClient(cfg.OllamaBaseURL)
	analyzerService := service.NewAnalyzerService(ollamaClient)
	router := api.NewRouter(analyzerService)
	log.Printf("Starting server on port %s...", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
