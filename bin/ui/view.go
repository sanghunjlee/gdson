package ui

import "fmt"

const (
	inputHelp string = `You can input multiple values using commas (,). Optinoally you can enclose the values with square brackets.
For numbers, you can input a range using a slice symbol (:).
Examples:
	(1) 1,2,3,4,5
	(2) [1,2,3,4,5]
	(3) [1:5]
	`
)

func (m model) View() string {
	switch m.state {
	case idleState:
		return m.mainMenu.View()
	case conditionInputState:
		return fmt.Sprintf(`%s
%s
%s
%s
%s
%s
		`,
			titleStyle.Width(m.width).Render(m.mainMenuChoice),
			helpStyle.Width(m.width).Render(inputHelp),
			m.conditionInputs[weekday].View(),
			m.conditionInputs[month].View(),
			m.conditionInputs[day].View(),
			m.conditionInputs[hour].View(),
		) + "\n"
	case dialogueInputState:
	}
	return ""
}
