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
		m.width = msg.Width
		m.conditionInputs.Width = msg.Width
		m.mainMenu.Width = msg.Width
		m.mainMenu.menu.SetWidth(msg.Width)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.state {
			case idleState:
				i, ok := m.mainMenu.SelectedItem().(item)
				if ok {
					m.mainMenuChoice = string(i)
					switch string(i) {
					case "Condition":
						m.state = conditionInputState
						m.conditionInputs.Focus()
					case "Dialogue":
						m.state = dialogueInputState
					case "Movement":
						m.state = movementInputState
					}
				}
			case conditionInputState:
				if m.conditionInputs.index == 3 {
					m.conditionInputs.quit = true
					return m, nil
				}
				m.conditionInputs.nextInput()
			}
		case tea.KeyShiftTab:
			switch m.state {
			case conditionInputState:
				m.conditionInputs.prevInput()
			}
		case tea.KeyTab:
			switch m.state {
			case conditionInputState:
				m.conditionInputs.nextInput()
			}
		}
	}
	m.mainMenu, cmd = m.mainMenu.Update(msg)
	cmds = append(cmds, cmd)

	m.conditionInputs, cmd = m.conditionInputs.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)

}
