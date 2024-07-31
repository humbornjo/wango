package latea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var inputWidth textinput.Model = textinput.New()

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
		m.ws.width = msg.Width
		m.ws.height = msg.Height
		return m, nil
	}

	var cmd tea.Cmd

	inputWidth, cmd = inputWidth.Update(message)
	return m, cmd
}

func (m model) View() string {
	text := fmt.Sprintf("window width: %d, height %d\n", m.ws.width, m.ws.height)
	return inputWidth.View() + "\n" + text
}
