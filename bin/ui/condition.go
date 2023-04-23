package ui

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func numValidator(s string) error {
	re := regexp.MustCompile(`^[0-9,:\[\]]*$`)
	if match := re.MatchString(s); !match {
		return fmt.Errorf("not valid input")
	}
	if strings.Count(s, ":") > 1 {
		return fmt.Errorf("invalid slicing")
	}
	if strings.Contains(s, ",,") {
		return fmt.Errorf("invalid array")
	}
	return nil
}

func monthValidator(s string) error {
	if err := numValidator(s); err != nil {
		return err
	}
	if len(s) > 0 {
		if _, err := strconv.ParseUint(s[len(s)-1:], 10, 8); err == nil {
			if len(s) > 1 {
				if m, err := strconv.ParseUint(s[len(s)-2:], 10, 8); err == nil {
					if m > 12 {
						return fmt.Errorf("invalid month")
					}
				}
			}
		}
	}
	return nil
}

func dayValidator(s string) error {
	if err := numValidator(s); err != nil {
		return err
	}
	if len(s) > 0 {
		if _, err := strconv.ParseUint(s[len(s)-1:], 10, 8); err == nil {
			if len(s) > 1 {
				if d, err := strconv.ParseUint(s[len(s)-2:], 10, 8); err == nil {
					if d > 31 {
						return fmt.Errorf("invalid day")
					}
				}
			}
		}
	}
	return nil
}

func hourValidator(s string) error {
	if err := numValidator(s); err != nil {
		return err
	}
	if len(s) > 0 {
		if _, err := strconv.ParseUint(s[len(s)-1:], 10, 8); err == nil {
			if len(s) > 1 {
				if d, err := strconv.ParseUint(s[len(s)-2:], 10, 8); err == nil {
					if d > 24 {
						return fmt.Errorf("invalid day")
					}
				}
			}
		}
	}
	return nil
}

func weekdayValidator(s string) error {
	return nil
}

type ConditionForm struct {
	Width        int
	focus        bool
	quit         bool
	index        int
	title        string
	desc         string
	weekdayInput textinput.Model
	monthInput   textinput.Model
	dayInput     textinput.Model
	hourInput    textinput.Model
}

func InitConditionForm() ConditionForm {
	t := "Condition"
	d := `You can input multiple values using commas (,). Optinoally you can enclose the values with square brackets.
For numbers, you can input a range using a slice symbol (:).
Examples:
	(1) 1,2,3,4,5
	(2) [1,2,3,4,5]
	(3) [1:5]`

	weekday := textinput.New()
	weekday.Focus()
	weekday.Placeholder = ""
	weekday.CharLimit = 156
	weekday.Prompt = "Weekday: "
	weekday.PromptStyle = promptStyle
	weekday.Validate = weekdayValidator

	month := textinput.New()
	month.Placeholder = ""
	month.CharLimit = 156
	month.Prompt = "Month: "
	month.PromptStyle = promptStyle
	month.Validate = monthValidator

	day := textinput.New()
	day.Placeholder = ""
	day.CharLimit = 156
	day.Prompt = "Day: "
	day.PromptStyle = promptStyle
	day.Validate = dayValidator

	hour := textinput.New()
	hour.Placeholder = ""
	hour.CharLimit = 156
	hour.Prompt = "Hour: "
	hour.PromptStyle = promptStyle
	hour.Validate = hourValidator

	return ConditionForm{
		focus:        false,
		quit:         false,
		index:        0,
		title:        t,
		desc:         d,
		weekdayInput: weekday,
		monthInput:   month,
		dayInput:     day,
		hourInput:    hour,
	}
}

func (f *ConditionForm) Focus() tea.Cmd {
	f.focus = true
	return nil
}

func (f *ConditionForm) Blur() {
	f.focus = false
	f.weekdayInput.Blur()
	f.monthInput.Blur()
	f.dayInput.Blur()
	f.hourInput.Blur()
}

func (f *ConditionForm) nextInput() {
	f.index = (f.index + 1) % 4
}

func (f *ConditionForm) prevInput() {
	f.index--
	if f.index < 0 {
		f.index = 3
	}
}

func (f ConditionForm) Update(msg tea.Msg) (ConditionForm, tea.Cmd) {
	if !f.focus {
		return f, nil
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		f.Width = msg.Width
		f.weekdayInput.Width = msg.Width
		f.monthInput.Width = msg.Width
		f.dayInput.Width = msg.Width
		f.hourInput.Width = msg.Width
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
		}
	}

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	f.weekdayInput.Blur()
	f.monthInput.Blur()
	f.dayInput.Blur()
	f.hourInput.Blur()
	switch f.index {
	case 0:
		cmd = f.weekdayInput.Focus()
	case 1:
		cmd = f.monthInput.Focus()
	case 2:
		cmd = f.dayInput.Focus()
	case 3:
		cmd = f.hourInput.Focus()
	}
	cmds = append(cmds, cmd)

	f.weekdayInput, cmd = f.weekdayInput.Update(msg)
	cmds = append(cmds, cmd)
	f.monthInput, cmd = f.monthInput.Update(msg)
	cmds = append(cmds, cmd)
	f.dayInput, cmd = f.dayInput.Update(msg)
	cmds = append(cmds, cmd)
	f.hourInput, cmd = f.hourInput.Update(msg)
	cmds = append(cmds, cmd)

	return f, tea.Batch(cmds...)
}
func (f ConditionForm) View() string {
	v := titleStyle.Width(f.Width).Render(f.title) + "\n"
	v += descStyle.Width(f.Width).Render(f.desc) + "\n"
	v += f.weekdayInput.View() + "\n"
	v += f.monthInput.View() + "\n"
	v += f.dayInput.View() + "\n"
	v += f.hourInput.View() + "\n"
	return v
}
