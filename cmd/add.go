package cmd
import(
	"fmt"
	"github.com/spf13/cobra"
	"go.bug.st/serial"
)
var addCmd = &cobra.Command{
	Use : "add",

	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string){
		Name := args[0]
		Login := args[1]
		Password := args[2]
		fmt.Println(Name, Login, Password)

		mode := &serial.Mode{
		BaudRate: 9600,
	}

		port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		panic(err)
	}
	defer port.Close()

values := []string{Name, Login, Password}

data := []byte("add" + "|" + values[0] + "|" + values[1] + "|" + values[2] + "\n")

port.Write(data)

if err != nil {
		panic(err)
	}

	fmt.Println("Sent:", values)

	},
}
