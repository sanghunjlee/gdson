package ui

func (m model) View() string {
	switch m.state {
	case idleState:
		return m.mainMenu.View()
	case conditionInputState:
		return m.conditionInputs.View()
	case dialogueInputState:
		return m.dialogueInputs.View()
	}
	return ""
}
