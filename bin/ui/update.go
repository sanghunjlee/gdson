package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		m.conditionInputs.SetSize(msg.Width, msg.Height)
		m.mainMenu.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			if m.state != idleState {
				m.state = idleState
				m.conditionInputs.Blur()
				m.dialogueInputs.Blur()

				cmd := m.mainMenu.Focus()
				return m, cmd
			}
			return m, tea.Quit
		}
	}

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.mainMenu, cmd = m.mainMenu.Update(msg)
	if m.mainMenu.IsDone() {
		m.mainMenu.Blur()
		switch m.mainMenu.Choice {
		case "Condition":
			m.state = conditionInputState
			cmd = m.conditionInputs.Focus()
			return m, cmd
		case "Dialogue":
			m.state = dialogueInputState
			cmd = m.dialogueInputs.Focus()
			return m, cmd
		}
	}
	cmds = append(cmds, cmd)

	m.conditionInputs, cmd = m.conditionInputs.Update(msg)
	if m.conditionInputs.IsDone() {
		m.conditionInputs.Blur()
		m.state = idleState
		cmd = m.mainMenu.Focus()
		return m, cmd
	}
	cmds = append(cmds, cmd)

	m.dialogueInputs, cmd = m.dialogueInputs.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
