package ui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type dialogueForm struct {
	focus  bool
	quit   bool
	width  int
	height int
	index  int

	Id           int
	Name         textinput.Model
	Say          textarea.Model
	OptionChoice choiceModel
	Options      []optionItem
	Next         textinput.Model
}

func NewDialogueForm() dialogueForm {
	name := textinput.New()

	say := textarea.New()

	choice := NewChoiceModel("Y", "N")
	choice.Prompt = "Option"

	opts := make([]optionItem, 4)
	opts[0] = NewOptionItem()
	opts[1] = NewOptionItem()
	opts[2] = NewOptionItem()
	opts[3] = NewOptionItem()

	next := textinput.New()

	return dialogueForm{

		Name:         name,
		Say:          say,
		OptionChoice: choice,
		Options:      opts,
		Next:         next,
	}
}

func (f *dialogueForm) Focus() tea.Cmd {
	f.focus = true
	return f.Name.Focus()
}

func (f *dialogueForm) Blur() {
	f.focus = false
	f.quit = false
	f.Name.Blur()
	f.Say.Blur()
	f.OptionChoice.Blur()
	for i := range f.Options {
		f.Options[i].Blur()
	}
	f.Next.Blur()
}

func (f dialogueForm) Update(msg tea.Msg) (dialogueForm, tea.Cmd) {
	if !f.focus {
		return f, nil
	}

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	f.Name, cmd = f.Name.Update(msg)
	cmds = append(cmds, cmd)
	f.Say, cmd = f.Say.Update(msg)
	cmds = append(cmds, cmd)
	f.OptionChoice, cmd = f.OptionChoice.Update(msg)
	cmds = append(cmds, cmd)
	for i := range f.Options {
		f.Options[i], cmd = f.Options[i].Update(msg)
		cmds = append(cmds, cmd)
	}
	f.Next, cmd = f.Next.Update(msg)
	cmds = append(cmds, cmd)

	return f, tea.Batch(cmds...)
}
