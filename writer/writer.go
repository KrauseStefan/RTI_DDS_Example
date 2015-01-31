package main

import (
	"github.com/KrauseStefan/RTI_DDS_Example/lib_wrapper"
	"bufio"
	"fmt"
	"os"
	"strings"
	//"time"
)

func main() {
	participant := lib_wrapper.Create_participant()

	topic := participant.CreateTopic("Hello, World")

	dataWriter := participant.CreateStringDatawriter(topic)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please type a message> ")

		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		if message == "" {
			break
		}

		if !dataWriter.Write(message) {
			fmt.Println("Write failed")
		}
	}

}
