package main

import (
	"flag"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/joho/godotenv"
	"github.com/slipynil/krypto/internal/api"
	"github.com/slipynil/krypto/internal/manager"
	"github.com/slipynil/krypto/internal/models"
	"github.com/slipynil/krypto/internal/storage"
	"github.com/slipynil/krypto/internal/tui"
)

func main() {
	p := tea.NewProgram(setupTui())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

func setupTui() tui.Model {
	apiKey := loadEnv()
	apiSvc := api.NewService(apiKey)
	storageSvc, _ := storage.NewService()
	mgr := manager.New(apiSvc, storageSvc)
	return *tui.NewModel(mgr, getCurr())
}

func getCurr() models.Currency {
	currency := flag.String("curr", "USD", "флаг для валюты, в которой будет все считаться\nexpample: krypto --curr RUB")
	flag.Parse()
	return models.Currency(*currency)
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
