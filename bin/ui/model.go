package ui

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
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

type conditionInputId int

const (
	weekday conditionInputId = iota
	month
	day
	hour
)
const listHeight = 14

var (
	titleStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Bold(true)
	labelStyle = lipgloss.NewStyle().
			Width(10).
			AlignHorizontal(lipgloss.Right).
			PaddingRight(1).
			Foreground(lipgloss.Color("20"))
	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("120"))
	paginationStyle = list.DefaultStyles().PaginationStyle.
			PaddingLeft(4)
	helpStyle = list.DefaultStyles().HelpStyle.
			PaddingLeft(4).
			PaddingBottom(1).
			Foreground(lipgloss.Color("240"))
	quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}
	str := fmt.Sprintf("%d. %s", index+1, i)
	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}
	fmt.Fprint(w, fn(str))
}

type model struct {
	state           sessionState
	width           int
	mainMenu        list.Model
	mainMenuChoice  string
	focused         int
	conditionInputs []textinput.Model
	dialogueInputs  []textinput.Model
	movementInputs  []textinput.Model
	err             error
}

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

func InitModel() model {
	menuItems := []list.Item{
		item("Condition"),
		item("Dialogue"),
		item("Movement"),
	}
	const defaultWidth = 20
	l := list.New(menuItems, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Main Menu\nPlease select which node to add/edit/remove"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	var cinps []textinput.Model = make([]textinput.Model, 4)
	cinps[weekday] = textinput.New()
	cinps[weekday].Placeholder = ""
	cinps[weekday].Focus()
	cinps[weekday].CharLimit = 156
	cinps[weekday].Prompt = "Weekday: "
	cinps[weekday].PromptStyle = labelStyle
	cinps[weekday].Validate = weekdayValidator

	cinps[month] = textinput.New()
	cinps[month].Placeholder = ""
	cinps[month].CharLimit = 156
	cinps[month].Prompt = "Month: "
	cinps[month].PromptStyle = labelStyle
	cinps[month].Validate = monthValidator

	cinps[day] = textinput.New()
	cinps[day].Placeholder = ""
	cinps[day].CharLimit = 156
	cinps[day].Prompt = "Day: "
	cinps[day].PromptStyle = labelStyle
	cinps[day].Validate = dayValidator

	cinps[hour] = textinput.New()
	cinps[hour].Placeholder = ""
	cinps[hour].CharLimit = 156
	cinps[hour].Prompt = "Hour: "
	cinps[hour].PromptStyle = labelStyle
	cinps[hour].Validate = hourValidator

	return model{
		state:           idleState,
		focused:         0,
		mainMenu:        l,
		conditionInputs: cinps,
		err:             nil,
	}
}
