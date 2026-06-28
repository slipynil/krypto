package manager

import (
	"strings"

	"github.com/slipynil/krypto/internal/api"
	"github.com/slipynil/krypto/internal/models"
	"github.com/slipynil/krypto/internal/storage"
)

type Manager struct {
	api     *api.ApiService
	storage *storage.StorageService
}

func New(apiSvc *api.ApiService, storageSvc *storage.StorageService) *Manager {
	return &Manager{
		api:     apiSvc,
		storage: storageSvc,
	}
}

// GetCoins - метод, который скрывает сложность от main.go
func (m *Manager) GetCoins() ([]models.Coin, error) {
	coins, err := m.api.GetCryptoIDs()
	if err != nil {
		return nil, err
	}

	err = m.storage.SaveData(coins)
	return coins, err
}

// FindCoins метод ищет подходящюю криптовалюту по запросу
func (m *Manager) FindCoins(query string) ([]models.Coin, error) {
	allCoins, err := m.GetCoins() // Берет из кеша или API
	if err != nil {
		return nil, err
	}

	var result []models.Coin
	for _, coin := range allCoins {
		if strings.Contains(strings.ToLower(coin.ID), strings.ToLower(query)) {
			result = append(result, coin)
		}
		if len(result) >= 20 { // Ограничиваем вывод 20-ю результатами
			break
		}
	}
	return result, nil
}

// выводит информацию о криптовалюте
func (m *Manager) GetInfo(id, name string, curr models.Currency) (*models.Price, error) {
	price, err := m.api.GetPriceInfo(id, curr)
	if err != nil {
		return nil, err
	}
	price.Name = name
	price.CoinID = id
	return price, nil
}
