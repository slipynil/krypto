# Krypto

<p>
  <a><img alt="go version 1.26.4" src="https://img.shields.io/badge/go-1.26.4-blue?logo=go"></a>
  <a href="https://docs.coingecko.com"><img width="20" src="https://avatars.githubusercontent.com/u/7111837?s=200&v=4" alt="coin gecko"></a>
  <a href="https://charm.land/libs/"><img width="20" alt="bubble tea v2" src="https://avatars.githubusercontent.com/u/57376114?s=48&v=4"</a>
</p>
<p><b>Krypto cli</b> - приложение для отслеживания рынка криптовалюты</p>

---

## Installation

```bash
go install github.com/slipynil/krypto@latest
```

## Tutorial

#### Design
Krypto предоставляет дружелюбный и простой интерфейс для для работы с CoinGecko API благодаря пакету <b>bubble tea v2</b>

<img alt="demo running" width="600" src="https://github.com/slipynil/krypto/blob/master/assets/krypto.gif?raw=true">

#### Flags
Существует 1 флаг `--curr`, отвечает за валюту, чтобы показывать цену.

```bash
# Запуск с валютой по умолчанию (USD)
krypto

# Запуск с указанием конкретной валюты
krypto --curr RUB
```

#### API endpoints
`/coins/list` — получает список всех доступных криптовалют

`/simple/price` — узнать актуальную цену монеты.

## Project Structure
```text
krypto/
├── internal/
│   ├── api/      # Транспортный уровень (CoinGecko client)
│   ├── storage/  # Локальное хранилище (кэширование монет)
│   └── tui/      # Интерфейс на Bubble Tea
└── cmd/
    └── main.go   # Точка входа
```
---
