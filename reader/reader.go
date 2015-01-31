package main

import (
	"github.com/KrauseStefan/RTI_DDS_Example/lib_wrapper"
	"bufio"
	"fmt"
	"os"
	//"strings"
	//"time"
)

func main() {
	participant := lib_wrapper.Create_participant()

	topic := participant.CreateTopic("Hello, World")

	participant.CreateDatareader(topic)

	fmt.Print("Press enter to exit\n")
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')

}
