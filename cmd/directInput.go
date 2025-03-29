package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type inputModel struct {
	textInput  textinput.Model
	err        error
	prompt     string
	takaAction func(string)
}

func (m inputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyEnter:
			m.takaAction(m.textInput.Value())
			return m, tea.Quit

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m inputModel) View() string {
	return fmt.Sprintf(m.prompt+"\n\n%s\n\n%s", m.textInput.View(), "(esc to quit)") + "\n"
}

func initialTextModel(prompt string, takeAction func(string)) inputModel {
	ti := textinput.New()
	ti.Placeholder = "Type here..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return inputModel{
		textInput:  ti,
		err:        nil,
		prompt:     prompt,
		takaAction: takeAction,
	}
}

func DirectInput(prompt string, takeAction func(string)) error {
	// p := tea.NewProgram(inputModel{prompt: prompt, takaAction: takeAction})

	p := tea.NewProgram(initialTextModel(prompt, takeAction))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	return nil
}
