package latea

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/humbornjo/wango/pkg/config"
)

func (m model) headerRender() string {
	nameStyle := lipgloss.NewStyle().
		Align(lipgloss.Center).Width(config.LayoutWidth)
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
		Height(max(len(odd), len(even)))

	leftbar := manbarStyle.Render(even...)
	rightbar := manbarStyle.Render(odd...)

	footer := lipgloss.JoinHorizontal(lipgloss.Left, leftbar, rightbar)
	return footer
}

func (m model) centerRender(page string) string {
	return centerStyle.
		Padding((winHeight-config.LayoutHeight)/2, (winWidth-config.LayoutWidth)/2).
		Render(page)
}

func (m model) bodyRender() {

}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("* " + label)
	}
	return fmt.Sprintf("  %s", label)
}

func choicesView(title string, choices []string, idx int) string {

	tpl := "%s\n%s"

	boxes := ""
	for i, choice := range choices {
		boxes += fmt.Sprintf("%s\n", checkbox(choice, idx == i))
	}
	boxes = boxes[:len(boxes)-1]

	return fmt.Sprintf(tpl, title, boxes)
}
