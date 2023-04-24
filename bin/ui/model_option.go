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
	return optionItem{
		Say:  s,
		Next: n,
	}
}

func (o *optionItem) Focus() tea.Cmd {
	o.focus = true
	return o.Say.Focus()
}

func (o *optionItem) Blur() {
	o.focus = false
	o.Say.Blur()
	o.Next.Blur()
}

func (o optionItem) Update(msg tea.Msg) (optionItem, tea.Cmd) {
	if !o.focus {
		return o, nil
	}

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	o.Say, cmd = o.Say.Update(msg)
	cmds = append(cmds, cmd)

	o.Next, cmd = o.Next.Update(msg)
	cmds = append(cmds, cmd)

	return o, tea.Batch(cmds...)
}
