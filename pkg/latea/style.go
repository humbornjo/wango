package latea

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/humbornjo/wango/pkg/config"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
)

var (
	checkboxStyle = lipgloss.NewStyle().Foreground(config.ClrFontFocus)
	subtleStyle   = lipgloss.NewStyle().Foreground(config.ClrFontDimed)
	centerStyle   = lipgloss.NewStyle().Width(winWidth).Height(winHeight)
	inputStyle    = lipgloss.NewStyle()
)

func BoxStyle(width int, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Left).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(config.ClrLayoutBorder).
		Foreground(config.ClrFontNomo).
		Padding(1, 1, 1, 1).
		Width(width).
		Height(height)
}
