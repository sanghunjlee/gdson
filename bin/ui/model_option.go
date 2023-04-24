package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type optionItem struct {
	focus bool
	Say   textinput.Model
	Next  textinput.Model
}

func NewOptionItem() optionItem {
	s := textinput.New()
	n := textinput.New()

	return optionItem{
		Say:  s,
		Next: n,
	}
}

func (f *optionItem) Focus() tea.Cmd {
	f.focus = true
	return f.Say.Focus()
}

func (f *optionItem) Blur() {
	f.focus = false
	f.Say.Blur()
	f.Next.Blur()
}
