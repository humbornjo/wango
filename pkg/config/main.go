package config

import "github.com/charmbracelet/lipgloss"

const (
	LayoutWidth  = 64
	LayoutHeight = 24
	LayoutMargin = 0
	LayoutBorder = 1
)

var (
	ClrLayoutBorder = lipgloss.Color("#FFEBD4")
	ClrFontNomo     = lipgloss.Color("#FAFAFA")
	ClrFontHard     = lipgloss.Color("#F0A8D0")
	ClrFontFocus    = lipgloss.Color("#FFC6C6")
	ClrFontDimed    = lipgloss.Color("#626262")
)

type Man struct {
	Key   string
	Usage string
}

type Choice struct {
	Label   string
	Choosen bool
}

var ChoicesMode = []Choice{
	{"      up", true},
	{"    down", false},
	{"   exact", false},
}

var ChoicesShader = []Choice{
	{"   moist", true},
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
