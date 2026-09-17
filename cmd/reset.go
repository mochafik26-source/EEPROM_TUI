package cmd

import(
	"fmt"
	"strings"
	"bufio"
	"log"
	"github.com/spf13/cobra"
	"go.bug.st/serial"
)

var resetCmd = &cobra.Command{
	Use : "reset",

	Run: func(cmd *cobra.Command, args []string){
		
		mode := &serial.Mode{
			BaudRate: 9600,
		}

		port, err := serial.Open("/dev/ttyUSB0", mode)

		if err != nil {
			panic(err)
		}
		defer port.Close()

		port.Write([]byte("reset|"))

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
