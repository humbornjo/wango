package config

import ()

const (
	LayoutWidth  = 64
	LayoutHeight = 24
	LayoutMargin = 0
	LayoutBorder = 1
)

type Man struct {
	Key   string
	Usage string
}

var ChoicesMode = []string{"  up", " fit", "down"}

var Manual = []Man{
	{" j/k       ", "select in box"},
	{" tab/S-tab ", "move in boxes"},
	{"enter     ", "start generating"},
	{"esc       ", "quit"},
}

func CoolName() string {
	return "" +
		"░█░░░█░█▀▀▄░█▀▀▄░█▀▀▀░▄▀▀▄░\n" +
		"░▀▄█▄▀░█▄▄█░█░▒█░█░▀▄░█░░█░\n" +
		"░░▀░▀░░▀░░▀░▀░░▀░▀▀▀▀░░▀▀░░ "
}
