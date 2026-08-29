package main

import(
    "fmt"
		"strings"
		"os"
		tea "charm.land/bubbletea/v2"
		"log"

	"charm.land/bubbles/v2/textinput"
	)

	var choices = []string{"Add", "Read"}
	type model struct {
		cursor int
		choice string
	}

	func (m model) Init() tea.Cmd{
		return nil
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
				return nil
		}
		
		}
		return m, nil
	}
func (m model) View() tea.View {
	s := strings.Builder{}
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
	s.WriteString("\n(press q to quit)\n")

	return tea.NewView(s.String())
}
	func main(){
p := tea.NewProgram(model{})

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




/*	mode := &serial.Mode{
		BaudRate : 9600,
	}

		port, err := serial.Open("/dev/ttyUSB1", mode)
	if err != nil {
		panic(err)
	}
	defer port.Close()
	
	time.Sleep(2 * time.Second)

	_, err = port.Write([]byte(`{"command","value"}`))
 if err != nil{
	 panic(err)
 }
 fmt.Println("done")*/

 
}
