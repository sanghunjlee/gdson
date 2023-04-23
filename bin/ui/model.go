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

var titleStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("220")).
	Margin(0, 2).
	Padding(1, 2).
	Align(lipgloss.Center).
	Bold(true)

var descStyle = lipgloss.NewStyle().
	PaddingLeft(4).
	PaddingBottom(1).
	Foreground(lipgloss.Color("240"))
var (
	labelStyle = lipgloss.NewStyle().
			Width(10).
			AlignHorizontal(lipgloss.Right).
			PaddingRight(1).
			Foreground(lipgloss.Color("20"))
	helpStyle = list.DefaultStyles().HelpStyle.
			PaddingLeft(4).
			PaddingBottom(1).
			Foreground(lipgloss.Color("240"))
	quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type model struct {
	state           sessionState
	width           int
	mainMenu        menuList
	focused         int
	conditionInputs ConditionForm
	dialogueInputs  []tea.Model
	movementInputs  []tea.Model
	err             error
}

func InitModel() model {

	m := InitMenu()

	c := InitConditionForm()

	return model{
		state:           idleState,
		focused:         0,
		mainMenu:        m,
		conditionInputs: c,
		err:             nil,
	}
}
