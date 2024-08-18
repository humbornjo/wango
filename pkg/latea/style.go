package latea

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/humbornjo/wango/pkg/config"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
)

var (
	choiceTitleStyle   = lipgloss.NewStyle().Foreground(config.ClrFontHard)
	checkboxStyle      = lipgloss.NewStyle().Foreground(config.ClrFontFocus)
	subtleStyle        = lipgloss.NewStyle().Foreground(config.ClrFontDimed)
	centerStyle        = lipgloss.NewStyle().Width(winWidth).Height(winHeight)
	inputStyle         = lipgloss.NewStyle()
	boldStyle          = lipgloss.NewStyle().Bold(true)
	boxChubbyStyle     = BoxStyle(config.BoxWidth, config.BoxHeightLong)
	boxSkinnyStyle     = BoxStyle(config.BoxWidth, config.BoxHeightShort)
	boxChubbyHalfStyle = BoxStyle(config.BoxWidthHalf, config.BoxHeightLong)
	boxSkinnyHalfStyle = BoxStyle(config.BoxWidthHalf, config.BoxHeightShort)
)

func BoxStyle(width int, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Left).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(config.ClrLayoutBorder).
		Foreground(config.ClrFontNomo).
		Padding(0, 1, 0, 1).
		Width(width).
		Height(height)
}

func TextinputStyle(charlimit int, prompt string) textinput.Model {
	ti := textinput.New()
	ti.Width = charlimit
	ti.CharLimit = charlimit
	ti.Prompt = prompt
	return ti
}
