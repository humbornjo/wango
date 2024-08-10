package latea

import (
	// "fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/humbornjo/wango/pkg/config"
)

var (
	inputWidth  textinput.Model = TextinputStyle(6, ": ")
	inputHeight textinput.Model = TextinputStyle(6, ": ")
	inputSize   textinput.Model = TextinputStyle(6, ": ")
	inputPath   textinput.Model = TextinputStyle(26, ": ")
	inputClrBg  textinput.Model = TextinputStyle(10, ": ")

	winWidth  int
	winHeight int
	ius       = []InputUnit{
		&TextinputUnit{&inputWidth, "width"},
		&TextinputUnit{&inputHeight, "height"},
		&TextinputUnit{&inputSize, "tile size"},
		&TextinputUnit{&inputPath, "save path"},
		&TextinputUnit{&inputClrBg, "background color"},
		&SingleChoiceUnit{config.ChoicesMode, "mode"},
		&SingleChoiceUnit{config.ChoicesShader, "shader"},
	}
)

func (m model) Init() tea.Cmd {
	inputWidth.CharLimit = 32
	return tea.Batch(
		tea.SetWindowTitle("wango"),
		textinput.Blink,
		inputWidth.Focus(),
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
			cmds = append(cmds, ius[m.stage].Blur())
			m.stage = (m.stage + 1) % len(ius)
			cmds = append(cmds, ius[m.stage].Focus())
			return m, tea.Batch(cmds...)
		case "shift+tab":
			cmds := []tea.Cmd{}
			cmds = append(cmds, ius[m.stage].Blur())
			m.stage = (m.stage - 1 + len(ius)) % len(ius)
			cmds = append(cmds, ius[m.stage].Focus())
			return m, tea.Batch(cmds...)
		case "enter": // TODO:
		}

	case tea.WindowSizeMsg:
		winWidth = msg.Width
		winHeight = msg.Height
		return m, nil
	}

	// cmd := ius[m.stage].TypeAction(message)
	var cmd tea.Cmd
	cmd = ius[m.stage].Update(message)
	return m, cmd
}

func (m model) View() (page string) {
	page += m.headerRender() + "\n\n"
	page += m.bodyRender() + "\n"
	page += m.footerRender()

	return m.centerRender(
		lipgloss.JoinVertical(
			lipgloss.Top,
			page,
		),
	)
}
