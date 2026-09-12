package cmd

import(
	"fmt"
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
		fmt.Println("reset")
	},

}
