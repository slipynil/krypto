package storage

import (
	"testing"

	"github.com/slipynil/krypto/internal/dto"
)

func TestStorageService_SaveAndLoad(t *testing.T) {
	// Создаем сервис
	svc, err := NewService()
	if err != nil {
		t.Fatalf("Не удалось создать сервис: %v", err)
	}

	// Тестовые данные
	expected := []dto.Coin{
		{ID: "bitcoin", Symbol: "btc", Name: "Bitcoin"},
		{ID: "ethereum", Symbol: "eth", Name: "Ethereum"},
	}

	// 1. Сохраняем
	err = svc.SaveData(expected)
	if err != nil {
		t.Fatalf("Ошибка при сохранении: %v", err)
	}

	// 2. Читаем
	actual, err := svc.LoadData()
	if err != nil {
		t.Fatalf("Ошибка при загрузке: %v", err)
	}

	// 3. Сравниваем
	if len(actual) != len(expected) {
		t.Errorf("Ожидали %d монет, получили %d", len(expected), len(actual))
	}
}
