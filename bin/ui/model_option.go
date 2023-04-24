package ui

import "github.com/charmbracelet/bubbles/textinput"

type optionItem struct {
	Say  textinput.Model
	Next int
}

func NewOptionItem() optionItem {
	s := textinput.New()
	return optionItem{
		Say: s,
	}
}
