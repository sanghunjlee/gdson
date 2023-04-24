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
	focus  bool
	quit   bool
	width  int
	height int
	title  string
	desc   string
	Choice string
	menu   list.Model
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
		focus: false,
		title: t,
		desc:  d,
		menu:  l,
	}
}

func (m *menuList) Focus() tea.Cmd {
	m.focus = true
	return nil
}
func (m *menuList) Blur() {
	m.focus = false
	m.quit = false
}

func (m *menuList) IsDone() bool {
	return m.quit
}

func (m *menuList) SetSize(w int, h int) {
	m.width = w
	m.height = h
}

func (m menuList) Update(msg tea.Msg) (menuList, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			i, ok := m.menu.SelectedItem().(item)
			if ok {
				m.Choice = string(i)
				m.quit = true
				return m, nil
			}
		}
	}

	var cmd tea.Cmd

	m.menu, cmd = m.menu.Update(msg)

	return m, cmd
}

func (m menuList) View() string {
	var availHeight = m.height

	title := m.titleView()
	desc := m.descView()

	availHeight -= lipgloss.Height(title) + lipgloss.Height(desc)

	content := lipgloss.NewStyle().Height(availHeight).Render(m.menu.View())
	return lipgloss.JoinVertical(lipgloss.Left, title, desc, content)
}

func (m menuList) titleView() string {
	titleStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Margin(1, 2).
		Padding(1).
		BorderForeground(lipgloss.Color("220")).
		Align(lipgloss.Center).
		Bold(true)

	gap := titleStyle.GetHorizontalPadding() +
		titleStyle.GetHorizontalMargins() +
		titleStyle.GetHorizontalBorderSize()

	return titleStyle.Width(m.width - gap).Render(m.title)
}

func (m menuList) descView() string {
	descStyle := lipgloss.NewStyle().
		PaddingLeft(4).
		PaddingBottom(1).
		Foreground(lipgloss.Color("240"))

	gap := descStyle.GetHorizontalPadding()

	return descStyle.Width(m.width - gap).Render(m.desc)
}
