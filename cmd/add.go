package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.bug.st/serial"
)

var addCmd = &cobra.Command{
	Use:  "add",
	Args: cobra.ExactArgs(3),

	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		login := args[1]
		password := args[2]

		// Open serial port
		mode := &serial.Mode{
			BaudRate: 9600,
		}

		port, err := serial.Open("/dev/ttyUSB0", mode)
		if err != nil {
			log.Fatal(err)
		}
		defer port.Close()

		// Data to send
		message := "add|" + name + "|" + login + "|" + password + "\n"

		_, err = port.Write([]byte(message))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Sent:", message)

		// -------------------------
		// Read Arduino response
		// -------------------------

		buffer := make([]byte, 1)
		var response []byte

		for {

			n, err := port.Read(buffer)
			if err != nil {
				log.Fatal(err)
			}

			if n > 0 {

				// If Arduino sends '1',
				// stop reading
				if buffer[0] == '1' {
					break
				}

				response = append(response, buffer[0])
			}
		}

		fmt.Println("Arduino:", string(response))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
