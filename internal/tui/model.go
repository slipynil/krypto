package tui

import (
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/slipynil/krypto/internal/manager"
	"github.com/slipynil/krypto/internal/models"
)

const (
	loadView uint = iota
	searchView
	selectedCoinView
)

// Model: храним состояние
type Model struct {
	state        uint
	list         list.Model
	manager      *manager.Manager
	currency     models.Currency
	selectedCoin *models.Price
	width        int
	height       int
	spinner      spinner.Model
}

func NewModel(mgr *manager.Manager, currency models.Currency) *Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return &Model{
		state:    loadView,
		currency: currency,
		manager:  mgr,
		spinner:  s,
	}
}

type errMsg error

// Init: начальное состояние
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg {
			coins, err := m.manager.GetCoins()
			if err != nil {
				return errMsg(err)
			}
			return LoadedCoinsMsg{Coins: coins}
		})
}

// Update: Логика tui приложения
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// перехватываем специфичные события
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.state != loadView {
			m.list.SetSize(msg.Width, msg.Height)
		}

	case tea.KeyMsg:
		key := msg.String()
		switch m.state {

		case searchView:
			if key == "space" {
				coin := m.list.SelectedItem().(models.Coin)
				return m, m.getPriceCmd(coin.ID, coin.Name)
			}

		case selectedCoinView:
			switch key {
			case "q", "esc":
				m.state = searchView
				return m, nil
			}
		}

	// обрабатываем инфу о выбранной монете
	case LoadedCoinsMsg:
		items := make([]list.Item, len(msg.Coins))
		for i, c := range msg.Coins {
			items[i] = c
		}
		m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)

		if m.width > 0 && m.height > 0 {
			m.list.SetSize(m.width, m.height)
		}

		m.state = searchView
		return m, nil

	case GotPriceMsg:
		if msg.Err == nil {
			// TODO: Handle error
			m.selectedCoin = msg.Price
			m.state = selectedCoinView
		} else {
			panic(msg.Err)
		}
		return m, nil

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if m.state != loadView {
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

type GotPriceMsg struct {
	Price *models.Price
	Err   error
}

type LoadedCoinsMsg struct {
	Coins []models.Coin
}

func (m Model) getPriceCmd(id, name string) tea.Cmd {
	return func() tea.Msg {
		price, err := m.manager.GetInfo(id, name, m.currency)
		return GotPriceMsg{Price: price, Err: err}
	}
}
