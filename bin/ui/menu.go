package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var itemStyle = lipgloss.NewStyle().
	PaddingLeft(4)
var selectedItemStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	Foreground(lipgloss.Color("120"))
var paginationStyle = list.DefaultStyles().PaginationStyle.
	PaddingLeft(4)

const listHeight = 14

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

type menuList struct {
	Width   int
	focused bool
	title   string
	desc    string
	choice  string
	menu    list.Model
}

func InitMenu() menuList {
	t := "Main Menu"
	d := "Select the node that you'd like to add / edit / remove"
	menuItems := []list.Item{
		item("Condition"),
		item("Dialogue"),
		item("Movement"),
	}
	const defaultWidth = 20
	l := list.New(menuItems, itemDelegate{}, defaultWidth, listHeight)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return menuList{
		focused: false,
		title:   t,
		desc:    d,
		menu:    l,
	}
}

func (m menuList) View() string {
	v := titleStyle.Width(m.Width).Render(f.title) + "\n"
	v += descStyle.Width(f.Width).Render(f.desc) + "\n"
	v += f.weekdayInput.View() + "\n"
	v += f.monthInput.View() + "\n"
	v += f.dayInput.View() + "\n"
	v += f.hourInput.View() + "\n"
	return v
}
