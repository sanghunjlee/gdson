package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		m.conditionInputs.SetSize(msg.Width, msg.Height)
		m.mainMenu.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.state {
			case idleState:
				i, ok := m.mainMenu.menu.SelectedItem().(item)
				if ok {
					m.mainMenu.choice = string(i)
					m.mainMenu.Blur()
					switch string(i) {
					case "Condition":
						m.state = conditionInputState
						cmd = m.conditionInputs.Focus()
					case "Dialogue":
						m.state = dialogueInputState
					case "Movement":
						m.state = movementInputState
					}
				}
				return m, cmd
			}
		}
	}
	m.mainMenu, cmd = m.mainMenu.Update(msg)
	cmds = append(cmds, cmd)

	m.conditionInputs, cmd = m.conditionInputs.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)

}
