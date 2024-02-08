package loggs

import (
	"fmt"
	"os"
)

func LogData(data string) {
	file, err := os.OpenFile("karakorum.io.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("#- Log file not found or  could not be created")
		return
	}

	defer file.Close()

	_, err2 := file.WriteString(data + "\n")

	if err2 != nil {
		fmt.Println("#-  Unable to write logs in the log file")
	}
}
