package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type inputModel struct {
	input      string
	prompt     string
	takaAction func(string)
}

func (m inputModel) Init() tea.Cmd {
	return nil
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "enter":
			m.takaAction(m.input)
			return m, tea.Quit
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			m.input += msg.String()
			return m, nil
		}
	}
	return m, nil
}

func (m inputModel) View() string {
	return fmt.Sprintf("%s\n\n%s", m.prompt, m.input)
}

func DirectInput(prompt string, takeAction func(string)) error {
	p := tea.NewProgram(inputModel{prompt: prompt, takaAction: takeAction})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	return nil
}
