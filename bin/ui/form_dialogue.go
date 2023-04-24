package ui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

type dialogueForm struct {
	focus  bool
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

	opt := NewChoiceModel("Y", "N")
	opt.Prompt = "Option"

	return dialogueForm{

		Name:         name,
		Say:          say,
		OptionChoice: opt,
	}
}
