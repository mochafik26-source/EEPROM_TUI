package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
	"go.bug.st/serial"
)

	var readCmd = &cobra.Command{
	Use: "Read",

	Args: cobra.ExactArgs(1),
	Run : func(cmd *cobra.Command, args []string){
		Name := args[0]
		fmt.Println(Name)
		
		mode := &serial.Mode{
			BaudRate: 9600,
		}

		port, err := serial.Open("/dev/ttyUSB0", mode)
		if err != nil {
			panic(err)
		}

		defer port.Close()
		data := []byte("read" + "|" + Name + "|")
		port.Write(data)

		fmt.Println(Name)
	},
}
