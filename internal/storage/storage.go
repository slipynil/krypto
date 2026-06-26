package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/slipynil/krypto/internal/dto"
)

type StorageService struct {
	filepath string
}

// constructor
func NewService() (*StorageService, error) {
	filepath, err := setupDataPath("cryptoIDs.json")
	if err != nil {
		return nil, err
	}
	return &StorageService{
		filepath: filepath,
	}, nil
}

// функция setup которая возвращает filepath для файла сохранения
func setupDataPath(filename string) (string, error) {
	// получаем системную папку
	baseDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	// создаем системный путь
	appDir := filepath.Join(baseDir, "krypto")

	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		if err := os.MkdirAll(appDir, 0755); err != nil {
			return "", err
		}
	}

	return filepath.Join(appDir, filename), nil
}

func (s *StorageService) SaveData(coins []dto.Coin) error {
	data, err := json.MarshalIndent(coins, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filepath, data, 0644)
}

func (s *StorageService) LoadData() ([]dto.Coin, error) {
	file, err := os.Open(s.filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var coins []dto.Coin
	if err := json.NewDecoder(file).Decode(&coins); err != nil {
		return nil, err
	}
	return coins, nil
}
