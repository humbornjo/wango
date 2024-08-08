package latea

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
	dotChar           = " • "
	width             = 96
	columnWidth       = 30
)

var (
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F7B5CA"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
)

func BoxStyle(width int, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Left).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFEBD4")).
		Foreground(lipgloss.Color("#FAFAFA")).
		Margin(0, 1, 0, 0).
		Padding(1, 1, 1, 3).
		Width(width).
		Height(height)
}

func PlainStyle(width int, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("241")).
		Margin(0, 0).
		Padding(0, 1).
		Width(width).
		Height(height)
}
