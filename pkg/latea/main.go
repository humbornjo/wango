package latea

import (
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
	inputClrBg  textinput.Model = TextinputStyle(10, ":\n")

	winWidth  int
	winHeight int
	ius       = []InputUnit{
		&TextinputUnit{&inputWidth, "Width    "},
		&TextinputUnit{&inputHeight, "Height   "},
		&TextinputUnit{&inputSize, "Tile Size"},
		&TextinputUnit{&inputPath, "Save Path"},
		&TextinputUnit{&inputClrBg, "Background Color (HEX)"},
		&SingleChoiceUnit{config.ChoicesMode, "Mode"},
		&SingleChoiceUnit{config.ChoicesShader, "Shader"},
		&MultiChoiceUnit{config.ChoicesFilter, "Filter"},
	}
)

func (m model) Init() tea.Cmd {
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
		case "enter":
			m.Generate()
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		winWidth = msg.Width
		winHeight = msg.Height
		return m, nil
	}

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
