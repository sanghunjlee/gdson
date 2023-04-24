package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sessionState int

const (
	idleState sessionState = iota
	conditionInputState
	dialogueInputState
	movementInputState
)

var promptStyle = lipgloss.NewStyle().
	Width(10).
	AlignHorizontal(lipgloss.Right).
	PaddingRight(1).
	Foreground(lipgloss.Color("20"))

var helpStyle = list.DefaultStyles().HelpStyle.
	PaddingLeft(4).
	PaddingBottom(1).
	Foreground(lipgloss.Color("240"))

var quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)

type model struct {
	state           sessionState
	width           int
	height          int
	focused         int
	mainMenu        menuList
	conditionInputs conditionForm
	dialogueInputs  []tea.Model
	movementInputs  []tea.Model
	err             error
}

func InitModel() model {

	m := InitMenu()
	m.Focus()

	c := InitConditionForm()

	return model{
		state:           idleState,
		focused:         0,
		mainMenu:        m,
		conditionInputs: c,
		err:             nil,
	}
}

func (m *model) SetSize(w int, h int) {
	m.width = w
	m.height = h
}
