package cmd

import (
	"bufio"
	"fmt"
	"log"
	"strings"

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

		// Send data to Arduino
		message := "add|" + name + "|" + login + "|" + password + "|"

		_, err = port.Write([]byte(message))
		if err != nil {
			log.Fatal(err)
		}


		reader := bufio.NewReader(port)

		for {
			data, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			data = strings.TrimSpace(data)

			fmt.Println(data)

			if data == "done" {
				break
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
