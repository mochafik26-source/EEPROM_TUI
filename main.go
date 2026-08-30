

package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/textinput"
)

var choices = []string{
	"Add",
	"Read",
}

type model struct {
	cursor    int
	input     textinput.Model
	showInput bool
	inputVal  []string
	step      int
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter value..."

	return model{
		cursor:    0,
		input:     ti,
		showInput: false,
		inputVal:  []string{},
		step:      0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		// Quit
		case "q":
			return m, tea.Quit

		// Move up
		case "up":
			m.cursor--

			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}

		// Move down
		case "down":
			m.cursor++

			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		// Enter
		case "enter":

			// Select "Add"
			if !m.showInput && m.cursor == 0 {
				m.showInput = true
				m.input.Focus()

				return m, textinput.Blink
			}

			// Save current input
			if m.showInput {

				value := m.input.Value()

				// Add value to slice
				m.inputVal = append(m.inputVal, value)

				m.step++
				// Clear input box
				m.input.Reset()
				if m.step > 2{
					return m, tea.Quit
				}
				// Go to next input

				return m, nil
			}
		}
	}

	// If we are entering text,
	// give the message to the text input.
	if m.showInput {

		var cmd tea.Cmd

		m.input, cmd = m.input.Update(msg)

		return m, cmd
	}

	return m, nil
}

func (m model) View() tea.View {

	var s strings.Builder

	// -------------------------
	// INPUT SCREEN
	// -------------------------

	if m.showInput && m.step <= 2{

		s.WriteString("Add\n\n")

		fmt.Fprintf(
			&s,
			"Input %d:\n\n",
			m.step+1,
		)

		s.WriteString(m.input.View())

		s.WriteString("\n\n")
		s.WriteString("Press Enter to save\n")
		s.WriteString("Press q to quit\n")

		return tea.NewView(s.String())
	}

	// -------------------------
	// MENU SCREEN
	// -------------------------

	s.WriteString("Choose an option:\n\n")

	for i, choice := range choices {

		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}

		s.WriteString(choice)
		s.WriteString("\n")
	}

	s.WriteString("\n")
	s.WriteString("↑/↓ to move\n")
	s.WriteString("Enter to select\n")
	s.WriteString("q to quit\n")

	return tea.NewView(s.String())
}

func main() {

	p := tea.NewProgram(initialModel())

	m, err := p.Run()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Get final model
	finalModel, ok := m.(model)

	if !ok {
		return
	}

	// Print saved values
	fmt.Println("\nSaved values:")

	for i, value := range finalModel.inputVal {
		fmt.Printf("%d: %s\n", i+1, value)
	}
}


