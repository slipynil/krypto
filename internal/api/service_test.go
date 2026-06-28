package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/slipynil/krypto/internal/models"
)

func TestGetPriceInfo_Success(t *testing.T) {
	data := []byte(`
{
    "ethereum": {
        "rub": 118872,
        "rub_market_cap": 14338438567941.64,
        "rub_24h_vol": 1254605779106.2979,
        "rub_24h_change": 1.844360094654535,
        "last_updated_at": 1782410327
    }
}
	`)
	// фековый сервер
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(data)
		}))
	defer ts.Close()

	svc := NewService("fake-key")
	svc.baseURL = ts.URL

	price, err := svc.GetPriceInfo("ethereum", models.CURRENCY_RUB)
	if err != nil {
		t.Fatal(err)
	}

	expectedValue := 118872.0
	if price.Value != expectedValue {
		t.Errorf("ожидалось Value %f, получили %f", expectedValue, price.Value)
	}

	if price.GrowthRate == 0 {
		t.Error("GrowthRate не должен быть 0")
	}

}

func TestGetPriceInfo_InvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{ "bad_json": }`))
		}))
	defer ts.Close()

	svc := NewService("fake-key")
	svc.baseURL = ts.URL
	_, err := svc.GetPriceInfo("bitcoin", models.CURRENCY_RUB)
	if err == nil {
		t.Error("Ожидали ошибку при битом JSON, но получили nil")
	}

}

func TestGetCryptoIDs_Success(t *testing.T) {
	data := []byte(`[{
    "id": "official-trump",
    "symbol": "trump",
    "name": "Official Trump",
    "platforms": {
      "solana": "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN"
    }
  },
  {
    "id": "ondo-finance",
    "symbol": "ondo",
    "name": "Ondo",
    "platforms": {
      "ethereum": "0xfaba6f8e4a5e8ab82f62fe7c39859fa577269be3"
    }
  }]`)
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(data)
		}))
	defer ts.Close()

	svc := NewService("fake-key")
	svc.baseURL = ts.URL

	result, err := svc.GetCryptoIDs()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("ID_Validation", func(t *testing.T) {
		if result[0].ID != "official-trump" {
			t.Errorf("expected ID official-trump, got %s", result[0].ID)
		}
	})
	t.Run("Symbol_Validation", func(t *testing.T) {
		if result[0].Symbol != "trump" {
			t.Errorf("expected Symbol trump, got %s", result[0].Symbol)
		}
	})
	t.Run("Name_Validation", func(t *testing.T) {
		if result[0].Name != "Official Trump" {
			t.Errorf("expacted Name 'Official Trump', got %s", result[0].Name)
		}
	})
	t.Run("Platform_Validation", func(t *testing.T) {
		if _, ok := result[0].Platforms["solana"]; !ok {
			t.Error("solana missing")
		}
	})
}

func TestGetCryptoIDs_InvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{ "bad_json": }`))
		}))
	defer ts.Close()

	svc := NewService("fake-key")
	svc.baseURL = ts.URL
	_, err := svc.GetCryptoIDs()
	if err == nil {
		t.Error("Ожидали ошибку при битом JSON, но получили nil")
	}
}

func TestCheckSwapCurrency(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("unable to load .env: %v", err)
	}
	apiKey := os.Getenv("API_KEY")
	svc := NewService(apiKey)
	_, err := svc.GetPriceInfo("ethereum", models.CURRENCY_RUB)
	if err != nil {
		t.Fatalf("failed to get price info for RUB currency: %v", err)
	}
	_, err = svc.GetPriceInfo("ethereum", models.CURRENCY_USD)
	if err != nil {
		t.Fatalf("failed to get price info for USD currency: %v", err)
	}
}
