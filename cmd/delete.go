package cmd
import(

	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"go.bug.st/serial"
)


var deleteCmd = &cobra.Command{
	Use : "delete",
	Args: cobra.ExactArgs(1),

	Run : func(cmd *cobra.Command, args []string){
		Name := args[0]

		mode := &serial.Mode{
			BaudRate: 9600,
		}

		port, err := serial.Open("/dev/ttyUSB0", mode)
		if err != nil {
			panic(err)
		}
		
		defer port.Close()

		data := []byte("delete" + "|" + Name + "|")

		port.Write(data)
	
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
