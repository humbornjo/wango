package latea

import (
	// "fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/humbornjo/wango/pkg/config"
)

var inputWidth textinput.Model = textinput.New()
var inputHeight textinput.Model = textinput.New()
var inputSize textinput.Model = textinput.New()
var inputPath textinput.Model = textinput.New()

var inputClrs textinput.Model = textinput.New()
var inputClrBg textinput.Model = textinput.New()

var inputs []any = []any{inputWidth, inputHeight}

var (
	winWidth  int
	winHeight int
)

func (m model) Init() tea.Cmd {
	inputWidth.CharLimit = 32
	return tea.Batch(
		tea.SetWindowTitle("wango"),
		textinput.Blink,
	)
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "tab":
			cmds := []tea.Cmd{}
			cmds = append(cmds, inputWidth.Focus())
			return m, tea.Batch(cmds...)
		}
	case tea.WindowSizeMsg:
		winWidth = msg.Width
		winHeight = msg.Height
		return m, nil
	}

	var cmd tea.Cmd

	inputWidth, cmd = inputWidth.Update(message)
	return m, cmd
}

func (m model) View() (page string) {
	page += "\n\n\n"
	{
		nameStyle := lipgloss.NewStyle().
			Align(lipgloss.Center).Width(72)
		page += nameStyle.Render(config.CoolName()) + "\n\n"
	}
	{
		inputStyle := BoxStyle(33, 12)
		input := inputStyle.
			Align(lipgloss.Left).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Left,
					inputWidth.View(), "\n",
					inputHeight.View(), "\n",
					inputSize.View(), "\n",
					inputPath.View(),
				),
			)

		choiceStyle := BoxStyle(15, 8)
		mode := choiceStyle.
			Align(lipgloss.Center).
			Render(choicesView("Mode", config.ChoicesMode, 1))
		shader := choiceStyle.
			Align(lipgloss.Center).
			Render(choicesView("Shader", config.ChoicesMode, 2))
		choices := lipgloss.JoinHorizontal(lipgloss.Center, mode, shader)
		leftbar := lipgloss.JoinVertical(lipgloss.Top, input, choices)

		filter := BoxStyle(32, 12).Render("\n\n")
		color := BoxStyle(32, 8).Render("")
		rightbar := lipgloss.JoinVertical(lipgloss.Top, filter, color)

		body := lipgloss.JoinHorizontal(lipgloss.Top, leftbar, rightbar)

		page += body + "\n"
	}

	{
		textStyle := PlainStyle(64, 2)
		// text := fmt.Sprintf(
		// 	"window width: %d, height %d\n",
		// 	winWidth,
		// 	winHeight,
		// )
		help := textStyle.Align(lipgloss.Center).Render("j/k: select in box" + dotStyle + "tab/shift+tab: move between boxes" + dotStyle + "enter: generate" + dotStyle + "esc: quit")
		footer := lipgloss.JoinVertical(lipgloss.Center, help)
		page += textStyle.Align(lipgloss.Center).Render(footer)
	}

	return page
}
