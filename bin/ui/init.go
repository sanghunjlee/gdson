package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (g model) Init() tea.Cmd {
	return textinput.Blink
}
