package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
	"github.com/slipynil/krypto/internal/api"
	"github.com/slipynil/krypto/internal/models"
)

func main() {
	apiKey := loadEnv()
	apiService := api.NewService(apiKey, models.CURRENCY_RUB)
	result, err := apiService.GetPriceInfo("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(result)
}

func loadEnv() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if len(apiKey) == 0 {
		log.Fatal("API_KEY not exist")
	}
	return apiKey
}
