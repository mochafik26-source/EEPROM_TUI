package main

import(
    "fmt"
		"strings"
		"os"
		tea "charm.land/bubbletea/v2"

	"charm.land/bubbles/v2/textinput"
	)

	var choices = []string{"Add", "Read"}
	type model struct {
		cursor int
		choice string
		situation string
		step int
		input textinput.Model
		showInput bool
	}

	func (m model) Init() tea.Cmd{
		return nil
	}

func initialModel() model {
	ti := textinput.New()
	
	return model{
		showInput: false,
		input: ti,
		cursor: 0,
		situation: "add",
	}
}
	func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String(){
			case "up":
				m.cursor++
				if m.cursor >= len(choices){
					m.cursor = 0
				}
			case "down":
				m.cursor--
				if m.cursor < 0{
					m.cursor = len(choices) - 1
				}
			case "q":
			return m, tea.Quit
			case "enter":
				if m.cursor == 0{
					m.showInput = true
					m.input.Focus()
					return m, textinput.Blink
				}
		}

		if m.showInput {
			var cmd tea.Cmd
			m.input, cmd = m.input.Update(msg)
			return m, cmd
		}
		
		}
		return m, nil
	}
func (m model) View() tea.View {
	var s strings.Builder

	s.WriteString("What kind of Bubble Tea would you like to order?\n\n")

	for i := range choices {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}

		s.WriteString(choices[i])
		s.WriteString("\n")
	}

	if m.situation == "read" {
		s.WriteString("\nYou are reading\n")
	}

	if m.situation == "add" {
		if !m.showInput {
			s.WriteString("\nPress SPACE to enter text\n")
		} else {
			s.WriteString("\n")
			s.WriteString(m.input.View())
			s.WriteString("\n")
		}
	}

	return tea.NewView(s.String())
}


func main(){
	p := tea.NewProgram(initialModel())
	// Run returns the model as a tea.Model.
	m, err := p.Run()

	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.choice)
	}






 
}
