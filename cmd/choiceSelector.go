package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices    []string
	cursor     int
	choice     string
	takaAction func(string)
}

func initialModel(choices []string, takeAction func(string)) model {
	return model{
		choices:    choices,
		choice:     "",
		takaAction: takeAction,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.choice != "" {
		return m, tea.Quit
	}

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.choice = m.choices[m.cursor]

			m.takaAction(m.choices[m.cursor])

			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Select a project\n\n"

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func ChoiceSelector(choices []string, takeAction func(string)) error {
	p := tea.NewProgram(initialModel(choices, takeAction))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	return nil
}
