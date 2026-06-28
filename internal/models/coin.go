package models

type Coin struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms"`
}

func (c Coin) Title() string {
	return c.Symbol
}

func (c Coin) Description() string {
	return c.Name
}

func (c Coin) FilterValue() string {
	return c.Name
}
