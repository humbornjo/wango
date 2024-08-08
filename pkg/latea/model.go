package latea

import (
	"image/color"
	"reflect"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/humbornjo/wango/pkg/render"
)

type WinSize struct {
	width, height int
}

type model struct {
	wang          render.Wang
	width, height int
	path          string
	mode          string
	texture       string
	palette       color.Palette
	seed          int
	winsize       WinSize
}

func InitModel() model {
	return model{}
}

type InputUnit struct {
	data      any
	selection int // 0 indicate not select, 1 single select, 2 multi select
	metaType  string
	index     int
}

func (iu *InputUnit) Action(message tea.Msg) tea.Cmd {
	switch reflect.TypeOf(iu.data) {
	// ran into a choices unit
	case reflect.TypeOf([]string{}):
		// array := iu.data.([]string)
		// single selection
		if iu.selection == 1 {

		}

		// multi selection
		if iu.selection == 2 {

		}
	// ran into a textinput unit
	case reflect.TypeOf(textinput.New()):
		text := iu.data.(textinput.Model)
		text, cmd := text.Update(message)
		return cmd
	}
	return nil
}
