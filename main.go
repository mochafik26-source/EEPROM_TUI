package main

import(
    "fmt"
    "time"
    "go.bug.st/serial"
		tea "charm.land/bubbletea/v2"
	)

func main(){
	mode := &serial.Mode{
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
 fmt.Println("done")
}
