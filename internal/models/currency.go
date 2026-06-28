package models

import "time"

type Currency string

const (
	CURRENCY_RUB = Currency("RUB")
	CURRENCY_USD = Currency("USD")
)

type Price struct {
	CoinID        string
	Name          string
	Value         float64
	MarketCap     float64
	Volume24h     float64
	GrowthRate    float64
	LastUpdatedAt time.Time
}
