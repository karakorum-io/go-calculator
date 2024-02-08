package main

import (
	"fmt"

	"karakorum.in/dev/go-calculator/execution"
	"karakorum.in/dev/go-calculator/information"
	"karakorum.in/dev/go-calculator/utils"
)

var ApplicationStatus bool = true

func main() {
	utils.ClearScreen()

	fmt.Println("Karakorum Simple Calculator")

	for ApplicationStatus {
		information.ChooseOperation()
		option := execution.GetOperation()
		status := execution.LaunchExecution(option)

		if !status {
			ApplicationStatus = false
		}
	}
}
