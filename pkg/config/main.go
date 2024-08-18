package config

import (
	"image/color"

	"github.com/charmbracelet/lipgloss"
)

const (
	WIDTH  = 2048
	HEIGHT = 1536
	SIZE   = 256
	PATH   = "./wang_tile.png"

	LayoutWidth    = 64
	LayoutHeight   = 24
	LayoutMargin   = 0
	LayoutBorder   = 1
	BoxWidth       = LayoutWidth/2 - LayoutBorder*2
	BoxWidthHalf   = (BoxWidth - LayoutBorder*2) / 2
	BoxHeightLong  = 12
	BoxHeightShort = 8
)

var (
	ClrLayoutBorder = lipgloss.Color("#FFEBD4")
	ClrFontNomo     = lipgloss.Color("#FAFAFA")
	ClrFontHard     = lipgloss.Color("#F0A8D0")
	ClrFontFocus    = lipgloss.Color("#FFC6C6")
	ClrFontDimed    = lipgloss.Color("#626262")

	ClrBackground = color.RGBA{}
)

type Man struct {
	Key   string
	Usage string
}

type Choice struct {
	Label    string
	Choosen  bool
	Selected bool
}

var ChoicesMode = []*Choice{
	{"up         ", true, false},
	{"down       ", false, false},
	{"exact      ", false, false},
}

var ChoicesShader = []*Choice{
	{"moist      ", true, false},
}

var ChoicesFilter = []*Choice{
	{"identical  ", true, false},
	{"noise      ", false, false},
}

var Manual = []Man{
	{" j/k       ", "select in box"},
	{" tab/S-tab ", "move in boxes"},
	{"space     ", "toggle confirmation"},
	{"enter     ", "start generating"},
	{"esc/C-c   ", "quit"},
}

func CoolName() string {
	return "" +
		"░█░░░█░█▀▀▄░█▀▀▄░█▀▀▀░▄▀▀▄░\n" +
		"░▀▄█▄▀░█▄▄█░█░▒█░█░▀▄░█░░█░\n" +
		"░░▀░▀░░▀░░▀░▀░░▀░▀▀▀▀░░▀▀░░ "
}
