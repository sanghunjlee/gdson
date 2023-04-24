package ui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type dialogueForm struct {
	focus  bool
	quit   bool
	width  int
	height int
	index  int

	title string
	desc  string

	Id           int
	Name         textinput.Model
	Say          textarea.Model
	OptionChoice choiceModel
	Options      []optionItem
	Next         textinput.Model
}

func NewDialogueForm() dialogueForm {
	t := "Dialogue"
	d := ``

	name := textinput.New()

	say := textarea.New()

	choice := NewChoiceModel("Y", "N")
	choice.Prompt = "Option"

	opts := make([]optionItem, 4)
	opts[0] = NewOptionItem()
	opts[1] = NewOptionItem()
	opts[2] = NewOptionItem()
	opts[3] = NewOptionItem()

	next := textinput.New()

	return dialogueForm{
		title:        t,
		desc:         d,
		Name:         name,
		Say:          say,
		OptionChoice: choice,
		Options:      opts,
		Next:         next,
	}
}

func (f *dialogueForm) Focus() tea.Cmd {
	f.focus = true
	return f.Name.Focus()
}

func (f *dialogueForm) Blur() {
	f.focus = false
	f.quit = false
	f.Name.Blur()
	f.Say.Blur()
	f.OptionChoice.Blur()
	for i := range f.Options {
		f.Options[i].Blur()
	}
	f.Next.Blur()
}

func (f dialogueForm) IsDone() bool {
	return f.quit
}

func (f *dialogueForm) SetSize(w int, h int) {
	f.width = w
	f.height = h
}
func (f dialogueForm) Update(msg tea.Msg) (dialogueForm, tea.Cmd) {
	if !f.focus {
		return f, nil
	}

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	f.Name, cmd = f.Name.Update(msg)
	cmds = append(cmds, cmd)
	f.Say, cmd = f.Say.Update(msg)
	cmds = append(cmds, cmd)
	f.OptionChoice, cmd = f.OptionChoice.Update(msg)
	cmds = append(cmds, cmd)
	for i := range f.Options {
		f.Options[i], cmd = f.Options[i].Update(msg)
		cmds = append(cmds, cmd)
	}
	f.Next, cmd = f.Next.Update(msg)
	cmds = append(cmds, cmd)

	return f, tea.Batch(cmds...)
}

func (f dialogueForm) View() string {
	var availHeight = f.height

	v := f.Name.View() + "\n"
	v += f.Say.View() + "\n"
	v += f.OptionChoice.View() + "\n"

	title := f.titleView()
	desc := f.descView()

	availHeight -= lipgloss.Height(title) + lipgloss.Height(desc)

	content := lipgloss.NewStyle().Height(availHeight).Render(v)
	return lipgloss.JoinVertical(lipgloss.Left, title, desc, content)
}

func (f dialogueForm) titleView() string {
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

	return titleStyle.Width(f.width - gap).Render(f.title)
}

func (f dialogueForm) descView() string {
	descStyle := lipgloss.NewStyle().
		PaddingLeft(4).
		PaddingBottom(1).
		Foreground(lipgloss.Color("240"))

	gap := descStyle.GetHorizontalPadding()

	return descStyle.Width(f.width - gap).Render(f.desc)
}
