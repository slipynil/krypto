package manager

import (
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
	coins, err := m.storage.LoadData()

	if err == nil && len(coins) > 0 {
		return coins, nil
	}

	coins, err = m.api.GetCryptoIDs()
	if err != nil {
		return nil, err
	}
	err = m.storage.SaveData(coins)
	return coins, err
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
