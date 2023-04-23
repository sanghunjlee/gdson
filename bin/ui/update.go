package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) nextInput() {
	switch m.state {
	case conditionInputState:
		m.focused = (m.focused + 1) % len(m.conditionInputs)
	case dialogueInputState:
		m.focused = (m.focused + 1) % len(m.dialogueInputs)
	case movementInputState:
		m.focused = (m.focused + 1) % len(m.movementInputs)
	}
}

func (m *model) prevInput() {
	m.focused--
	if m.focused < 0 {
		switch m.state {
		case conditionInputState:
			m.focused = len(m.conditionInputs) - 1
		case dialogueInputState:
			m.focused = len(m.dialogueInputs) - 1
		case movementInputState:
			m.focused = len(m.movementInputs) - 1
		}
	}
}

func (m model) refocus() {
	switch m.state {
	case conditionInputState:
		for i := range m.conditionInputs {
			m.conditionInputs[i].Blur()
		}
		m.conditionInputs[m.focused].Focus()
	case dialogueInputState:
		for i := range m.dialogueInputs {
			m.dialogueInputs[i].Blur()
		}
		m.dialogueInputs[m.focused].Focus()
	case movementInputState:
		for i := range m.movementInputs {
			m.movementInputs[i].Blur()
		}
		m.movementInputs[m.focused].Focus()
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		switch m.state {
		case idleState:
			m.mainMenu.SetWidth(msg.Width)
		case conditionInputState:
			for i := range m.conditionInputs {
				m.conditionInputs[i].Width = msg.Width
			}
		case dialogueInputState:
			for i := range m.dialogueInputs {
				m.dialogueInputs[i].Width = msg.Width
			}
		case movementInputState:
			for i := range m.movementInputs {
				m.movementInputs[i].Width = msg.Width
			}
		}
		return m, nil

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
					case "Dialogue":
						m.state = dialogueInputState
					case "Movement":
						m.state = movementInputState
					}
				}
			case conditionInputState:
				if m.focused == len(m.conditionInputs)-1 {
					m.state = idleState
					return m, nil
				}
				m.nextInput()
			case dialogueInputState:
				if m.focused == len(m.dialogueInputs)-1 {
					m.state = idleState
					return m, nil
				}
				m.nextInput()
			case movementInputState:
				if m.focused == len(m.movementInputs)-1 {
					m.state = idleState
					return m, nil
				}
				m.nextInput()
			}
		case tea.KeyShiftTab:
			m.prevInput()
		case tea.KeyTab:
			m.nextInput()
		}
		m.refocus()
	}

	m.mainMenu, cmd = m.mainMenu.Update(msg)
	cmds = append(cmds, cmd)

	for i := range m.conditionInputs {
		m.conditionInputs[i], cmd = m.conditionInputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)

}
