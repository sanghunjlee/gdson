package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type choiceModel struct {
	focus bool
	width int

	cursor   int
	selected string
	choices  []string

	Prompt      string
	PromptStyle lipgloss.Style
	ChoiceStyle lipgloss.Style
}

func NewChoiceModel(args ...string) choiceModel {
	if args == nil {
		args = []string{""}
	}
	return choiceModel{
		focus:       false,
		cursor:      0,
		selected:    "",
		choices:     args,
		Prompt:      "",
		PromptStyle: lipgloss.NewStyle().PaddingLeft(1),
		ChoiceStyle: lipgloss.NewStyle().PaddingLeft(1),
	}
}

func (m *choiceModel) Focus() tea.Cmd {
	m.focus = true
	return nil
}

func (m *choiceModel) Blur() {
	m.focus = false
}

func (m *choiceModel) SetWidth(w int) {
	m.width = w
}

func (m choiceModel) Update(msg tea.Msg) (choiceModel, tea.Cmd) {
	if !m.focus {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.selected = m.choices[m.cursor]
			return m, nil
		case tea.KeyRight:
			m.cursor++
			if m.cursor > len(m.choices) {
				m.cursor = len(m.choices) - 1
			}
		case tea.KeyLeft:
			m.cursor--
			if m.cursor < 0 {
				m.cursor = 0
			}
		}
	}

	return m, nil
}

func (m choiceModel) View() string {
	prompt := m.PromptStyle.Render(m.Prompt)
	var content string
	for i := 0; i < len(m.choices); i++ {
		var s string
		if m.cursor == i {
			s = "(•) "
		} else {
			s = "( ) "
		}
		s += m.choices[i]
		content += m.ChoiceStyle.Render(s)
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, prompt, content)
}
