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
		case "enter":
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
	page += m.headerRender() + "\n\n"
	{

		inputStyle := BoxStyle(30, 12)
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

		choiceStyle := BoxStyle(14, 8)
		mode := choiceStyle.
			Align(lipgloss.Center).
			Render(choicesView("Mode", config.ChoicesMode))
		shader := choiceStyle.
			Align(lipgloss.Center).
			Render(choicesView("Shader", config.ChoicesShader))
		choices := lipgloss.JoinHorizontal(lipgloss.Center, mode, shader)
		leftbar := lipgloss.JoinVertical(lipgloss.Top, input, choices)

		filter := BoxStyle(30, 12).Render("\n\n")
		color := BoxStyle(30, 8).Render("")
		rightbar := lipgloss.JoinVertical(lipgloss.Top, filter, color)

		body := lipgloss.JoinHorizontal(lipgloss.Top, leftbar, rightbar)

		page += body + "\n"
	}
	page += m.footerRender()

	return m.centerRender(page)
}
