package dto

import (
	"time"

	"github.com/slipynil/krypto/internal/models"
)

// PriceDataUSD содержит конкретные показатели для монеты в USD
type DataUSD struct {
	USD         float64 `json:"usd"`
	MarketCap   float64 `json:"usd_market_cap"`
	Volume24h   float64 `json:"usd_24h_vol"`
	Change24h   float64 `json:"usd_24h_change"`
	LastUpdated int64   `json:"last_updated_at"`
}

// PriceDataRUB содержит конкретные показатели для монеты в RUB
type DataRUB struct {
	RUB         float64 `json:"rub"`
	MarketCap   float64 `json:"rub_market_cap"`
	Volume24h   float64 `json:"rub_24h_vol"`
	Change24h   float64 `json:"rub_24h_change"`
	LastUpdated int64   `json:"last_updated_at"`
}

func (d DataRUB) ToModel() *models.Price {
	return &models.Price{
		CurrencyName:  "RUB",
		Value:         d.RUB,
		MarketCap:     d.MarketCap,
		Volume24h:     d.Volume24h,
		GrowthRate:    (1 - d.Change24h) * 100,
		LastUpdatedAt: time.Unix(d.LastUpdated, 0),
	}

}
func (d DataUSD) ToModel() *models.Price {
	return &models.Price{
		CurrencyName:  "USD",
		Value:         d.USD,
		MarketCap:     d.MarketCap,
		Volume24h:     d.Volume24h,
		GrowthRate:    (1 - d.Change24h) * 100,
		LastUpdatedAt: time.Unix(d.LastUpdated, 0),
	}
}
