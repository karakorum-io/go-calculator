package information

import "fmt"

func ChooseOperation() {
	fmt.Println("Choose operation! \n 1. Add \n 2. Subtract\n 3. Multiply\n 4. Divide\n 5. Exit")
}

func EnterFirst() {
	fmt.Println("Enter First Operand:")
}

func EnterSecond() {
	fmt.Println("Enter Second Operand:")
}

func Output(operation string, output int) {
	fmt.Println("Result for the Operation", operation, "is", output)
}
