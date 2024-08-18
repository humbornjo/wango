package latea

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/humbornjo/wango/pkg/config"
	"github.com/lucasb-eyer/go-colorful"
	"strings"
)

func (m model) headerRender() string {
	nameStyle := lipgloss.NewStyle().
		Foreground(config.ClrFontNomo).
		Align(lipgloss.Center).
		Width(config.LayoutWidth)
	header := nameStyle.Render(config.CoolName())
	return "\n\n\n" + header
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

func (m model) bodyRender() string {
	var upleft, upright, bottomleft, bottomright string

	{
		upleft = boxChubbyStyle.
			Align(lipgloss.Left).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Left,
					ius[0].View(), "\n",
					ius[1].View(), "\n",
					ius[2].View(), "\n",
					ius[3].View(),
				),
			)
	}

	{
		mode := boxSkinnyHalfStyle.
			Align(lipgloss.Left).
			Render(ius[5].View())
		shader := boxSkinnyHalfStyle.
			Align(lipgloss.Left).
			Render(ius[6].View())
		bottomleft = lipgloss.JoinHorizontal(lipgloss.Center, mode, shader)
	}

	leftbar := lipgloss.JoinVertical(lipgloss.Top, upleft, bottomleft)

	{
		colors := func() string {
			colors := colorGrid(config.BoxWidthHalf, 7)
			b := strings.Builder{}
			for _, x := range colors {
				for _, y := range x {
					s := lipgloss.NewStyle().SetString("  ").Background(lipgloss.Color(y))
					b.WriteString(s.String())
				}
				b.WriteRune('\n')
			}
			return b.String()
		}()
		upright = boxChubbyStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				colors,
				ius[4].View(),
			),
		)
	}

	{
		bottomright = boxSkinnyStyle.
			Align(lipgloss.Left).
			Render(ius[7].View())
	}

	rightbar := lipgloss.JoinVertical(lipgloss.Top, upright, bottomright)

	body := lipgloss.JoinHorizontal(lipgloss.Top, leftbar, rightbar)

	return body
}

func checkbox(choice *config.Choice) string {
	var selected string
	var choosen string
	if choice.Selected {
		selected = "* "
	} else {
		selected = "  "
	}
	if choice.Choosen {
		choosen = checkboxStyle.Bold(false).Render(choice.Label)
	} else {
		choosen = choice.Label
	}
	return choosen + selected
}

func choicesView(title string, choices []*config.Choice) string {
	view := boldStyle.Render(title) + "\n"
	view += "──────────" + "\n"

	for _, choice := range choices {
		view += fmt.Sprintf("%s\n", checkbox(choice))
	}
	view = view[:len(view)-1]

	return view
}

// Color grid
func colorGrid(xSteps, ySteps int) [][]string {
	x0y0, _ := colorful.Hex(string(config.ClrFontFocus))
	x1y0, _ := colorful.Hex(string(config.ClrFontHard))
	x0y1, _ := colorful.Hex(string(config.ClrFontDimed))
	x1y1, _ := colorful.Hex(string(config.ClrFontNomo))

	x0 := make([]colorful.Color, ySteps)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(ySteps))
	}

	x1 := make([]colorful.Color, ySteps)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(ySteps))
	}

	grid := make([][]string, ySteps)
	for x := 0; x < ySteps; x++ {
		y0 := x0[x]
		grid[x] = make([]string, xSteps)
		for y := 0; y < xSteps; y++ {
			grid[x][y] = y0.BlendLuv(x1[x], float64(y)/float64(xSteps)).Hex()
		}
	}

	return grid
}
