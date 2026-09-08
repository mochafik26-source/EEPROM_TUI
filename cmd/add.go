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
		fmt.Println(Name)
		mode := &serial.Mode{
		BaudRate: 9600,
	}

		port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		panic(err)
	}
	defer port.Close()

	value := "Hello Arduino\n"

	_, err = port.Write([]byte(value))
	if err != nil {
		panic(err)
	}

	fmt.Println("Sent:", value)

	time.Sleep(time.Second)
	},
}
