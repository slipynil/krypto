package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/slipynil/krypto/internal/models"
)

var (
	footerStyle = lipgloss.NewStyle().
			Padding(1, 0, 0, 0).
			Foreground(lipgloss.Color("240")) // Темно-серый текст

	keyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")). // Яркий розовый для клавиш
			Bold(true)
)

func (m Model) View() tea.View {
	var content string
	switch m.state {
	case loadView:
		return tea.NewView(m.spinner.View())
	case searchView:
		content = m.list.View()
	case selectedCoinView:
		content = m.renderSelectedCoinView()
	case errorView:
		content = fmt.Sprintf("Ошибка: %v\n\npress q to exit", m.err)
	}

	return tea.NewView(content)
}

func (m Model) renderSelectedCoinView() string {
	p := m.selectedCoin
	if p == nil {
		return "Загрузка данных..."
	}

	// 1. Определяем символ валюты (можно расширить логику, если валют будет больше)
	currencySymbol := "$"
	if m.currency == models.CURRENCY_RUB {
		currencySymbol = "₽"
	}

	// 2. Стили
	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Width(16)
	valueStyle := lipgloss.NewStyle().Bold(true)

	growthColor := "#00FF00"
	if p.GrowthRate < 0 {
		growthColor = "#FF0000"
	}
	growthStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(growthColor)).Bold(true)

	// 3. Хелпер для строк
	row := func(label string, value string) string {
		return lipgloss.JoinHorizontal(lipgloss.Left, labelStyle.Render(label), valueStyle.Render(value))
	}

	// 4. Формируем тело с учетом выбранной валюты
	details := lipgloss.JoinVertical(lipgloss.Left,
		row("Цена:", fmt.Sprintf("%s%.2f", currencySymbol, p.Value)),
		row("Изменение:", growthStyle.Render(fmt.Sprintf("%.2f%%", p.GrowthRate))),
		row("Капитализация:", fmt.Sprintf("%s%s", currencySymbol, formatLargeNumber(p.MarketCap))),
		row("Объем (24ч):", fmt.Sprintf("%s%s", currencySymbol, formatLargeNumber(p.Volume24h))),
		row("Обновлено:", p.LastUpdatedAt.Format("15:04:05")),
	)

	// 5. Оборачиваем в карточку
	card := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Width(40). // Фиксируем ширину, чтобы рамка была ровной
		Render(details)

	// 6. Заголовок
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7D56F4")).
		Bold(true).
		Render(fmt.Sprintf("--- %s (%s) ---", p.Name, string(m.currency)))

	// 7. Подсказки
	line := footerStyle.Render(
		fmt.Sprintf("Press %s to quit • %s to change currency",
			keyStyle.Render("q"),
			keyStyle.Render("c")),
	)

	return lipgloss.JoinVertical(lipgloss.Center, title, card) + line
}

func formatLargeNumber(val float64) string {
	switch {
	case val >= 1_000_000_000:
		return fmt.Sprintf("%.2f млрд", val/1_000_000_000)
	case val >= 1_000_000:
		return fmt.Sprintf("%.2f млн", val/1_000_000)
	default:
		return fmt.Sprintf("%.2f", val)
	}
}
