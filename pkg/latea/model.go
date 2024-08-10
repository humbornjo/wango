package latea

import (
	"image/color"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/humbornjo/wango/pkg/config"
	"github.com/humbornjo/wango/pkg/render"
)

type WinSize struct {
	width, height int
}

type model struct {
	wang          render.Wang
	width, height int
	path          string
	texture       string
	palette       color.Palette
	seed          int
	stage         int
}

func InitModel() model {
	return model{}
}

type InputUnit interface {
	Focus() tea.Cmd
	Update(tea.Msg) tea.Cmd
	Blur() tea.Cmd
	View() string
}

// single choice only panel unit
type SingleChoiceUnit struct {
	data  []*config.Choice
	title string
}

func (u *SingleChoiceUnit) Focus() tea.Cmd {
	for _, choice := range u.data {
		choice.Selected = choice.Choosen
	}
	return nil
}

func (u *SingleChoiceUnit) Update(message tea.Msg) tea.Cmd {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "k":
			var idx int
			var offset int
			var length int = len(u.data)
			if msg.String() == "j" {
				offset = 1
			} else {
				offset = -1
			}
			for i, choice := range u.data {
				if choice.Selected {
					idx = i
				}
				choice.Selected = false
			}

			u.data[(idx+offset+length)%length].Selected = true
		case " ":
			for _, choice := range u.data {
				choice.Choosen = choice.Selected
			}
		}
	}
	return nil
}

func (u *SingleChoiceUnit) View() string {
	return choicesView(u.title, u.data)
}

func (u *SingleChoiceUnit) Blur() tea.Cmd {
	for _, choice := range u.data {
		choice.Selected = false
	}
	return nil
}

// multi choice only panel unit
type MultiChoiceUnit struct {
	data  []*config.Choice
	title string
}

func (u *MultiChoiceUnit) Focus() tea.Cmd {
	for _, choice := range u.data {
		if choice.Choosen {
			choice.Selected = choice.Choosen
			break
		}
	}
	return nil
}

func (u *MultiChoiceUnit) Update(message tea.Msg) tea.Cmd {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "k":
			var idx int
			var offset int
			var length int = len(u.data)
			if msg.String() == "j" {
				offset = 1
			} else {
				offset = -1
			}
			for i, choice := range u.data {
				if choice.Selected {
					idx = i
				}
				choice.Selected = false
			}

			u.data[(idx+offset+length)%length].Selected = true
		case " ":
			for _, choice := range u.data {
				choice.Choosen = !choice.Choosen
			}
		}
	}
	return nil
}

func (u *MultiChoiceUnit) View() string {
	return choicesView(u.title, u.data)
}

func (u *MultiChoiceUnit) Blur() tea.Cmd {
	for _, choice := range u.data {
		choice.Selected = false
	}
	return nil
}

// text input panel unit
type TextinputUnit struct {
	data  *textinput.Model
	title string
}

func (u *TextinputUnit) Focus() tea.Cmd {
	return u.data.Focus()
}

func (u *TextinputUnit) Update(message tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	*u.data, cmd = u.data.Update(message)
	return cmd
}

func (u *TextinputUnit) Blur() tea.Cmd {
	u.data.Blur()
	return nil
}

func (u *TextinputUnit) View() string {
	return u.title + u.data.View()
}
