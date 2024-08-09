package latea

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/humbornjo/wango/pkg/config"
)

func (m model) headerRender() string {
	nameStyle := lipgloss.NewStyle().
		Foreground(config.ClrFontNomo).
		Align(lipgloss.Center).
		Width(config.LayoutWidth)
	header := nameStyle.Render(config.CoolName())
	return header
}

func (m model) footerRender() string {
	var odd []string
	var even []string

	for i, man := range config.Manual {
		renderedItem := subtleStyle.Bold(true).Render(man.Key) +
			subtleStyle.Bold(false).Render(man.Usage)
		if i%2 == 0 {
			even = append(even, renderedItem)
		} else {
			odd = append(odd, renderedItem)
		}
	}

	for i := range len(odd) - 1 {
		odd[i] += "\n"
	}

	for i := range len(even) - 1 {
		even[i] += "\n"
	}

	manbarStyle := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Width(config.LayoutWidth / 2).
		Height((len(config.Manual) + 1) / 2)

	leftbar := manbarStyle.Render(even...)
	rightbar := manbarStyle.Render(odd...)

	footer := lipgloss.JoinHorizontal(lipgloss.Left, leftbar, rightbar)
	return footer
}

func (m model) centerRender(page string) string {
	padWidth := (winWidth - config.LayoutWidth) / 2
	padHeight := (winHeight - config.LayoutHeight) / 2
	centered := centerStyle.
		Padding(padHeight, padWidth).
		Align(lipgloss.Center).
		Render(page)
	return centered
}

func (m model) bodyRender() {

}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Bold(false).Render("* " + label)
	}
	return fmt.Sprintf("  %s", label)
}

func choicesView(title string, choices []config.Choice) string {
	view := choiceTitleStyle.Bold(true).Render(title) + "\n"

	for _, choice := range choices {
		view += fmt.Sprintf("%s\n", checkbox(choice.Label, choice.Choosen))
	}
	view = view[:len(view)-1]

	return view
}
