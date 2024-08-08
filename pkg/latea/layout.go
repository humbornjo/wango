package latea

import (
	"fmt"
)

func (m model) footerRender() {

}

func (m model) cuisineRender() {

}

func (m model) mainPanelRender() {

}

func (m model) stagePaginatorRender() {

}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("* " + label)
	}
	return fmt.Sprintf("  %s", label)
}

func choicesView(title string, choices []string, idx int) string {

	tpl := "%s\n%s"

	boxes := ""
	for i, choice := range choices {
		boxes += fmt.Sprintf("%s\n", checkbox(choice, idx == i))
	}
	boxes = boxes[:len(boxes)-1]

	return fmt.Sprintf(tpl, title, boxes)
}
