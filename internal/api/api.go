package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/slipynil/krypto/internal/dto"
	"github.com/slipynil/krypto/internal/models"
)

type ApiService struct {
	apiKey   string
	currency models.Currency
	client   *http.Client
	baseURL  string
}

// constructor
func NewService(apiKey string, curr models.Currency) *ApiService {
	return &ApiService{
		apiKey:   apiKey,
		currency: curr,
		client:   &http.Client{},
		baseURL:  "https://api.coingecko.com/api/v3",
	}
}

func (a *ApiService) GetCryptoIDs() ([]models.Coin, error) {
	u, err := url.JoinPath(a.baseURL, "coins", "list")
	if err != nil {
		return nil, err
	}

	resp, err := a.client.Get(u)
	if err != nil {
		return nil, err
	}

	var result []models.Coin
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, err
}

// возвращает информацию о криптовалюте
func (a *ApiService) GetPriceInfo(id string) (*models.Price, error) {
	u, _ := url.JoinPath(a.baseURL, "simple", "price")
	u = a.priceQueryRequest(u, id)

	// костыль пока не выйдет go1.27.1
	switch a.currency {
	case models.CURRENCY_RUB:
		resp := map[string]dto.DataRUB{}
		if err := a.fetchData(u, &resp); err != nil {
			return nil, err
		}
		return checkResponse(resp, id)

	case models.CURRENCY_USD:
		resp := map[string]dto.DataUSD{}
		if err := a.fetchData(u, &resp); err != nil {
			return nil, err
		}
		return checkResponse(resp, id)

	default:
		return nil, fmt.Errorf("currency this type %s not exist", a.currency)
	}
}

// Вспомогательная функция-метод для [GetPriceInfo].
// Возвращает url с настроенными query параметрами
func (a *ApiService) priceQueryRequest(u, id string) string {
	parsed, _ := url.Parse(u)

	params := url.Values{}
	params.Add("vs_currencies", string(a.currency))
	params.Add("ids", id)
	params.Add("include_last_updated_at", "true")
	params.Add("include_24hr_change", "true")
	params.Add("include_market_cap", "true")
	params.Add("include_24hr_vol", "true")

	parsed.RawQuery = params.Encode()
	return parsed.String()
}

// метод просто делает запрос, dto определяется снаружи.
// ОБЯЗАТЕЛЬНО taget с указателем
func (a *ApiService) fetchData(url string, target any) error {
	resp, err := a.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(&target)
}

type PriceMapper interface {
	ToModel() *models.Price
}

func checkResponse[T PriceMapper](data map[string]T, id string) (*models.Price, error) {
	if data, ok := data[id]; !ok {
		return nil, fmt.Errorf("coin not found in response")
	} else {
		return data.ToModel(), nil
	}
}
