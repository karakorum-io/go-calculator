package execution

import (
	"fmt"
	"log"
	"strconv"

	"karakorum.in/dev/go-calculator/information"
	model "karakorum.in/dev/go-calculator/models"
	"karakorum.in/dev/go-calculator/utils"
)

func GetOperation() string {
	var option string
	_, err := fmt.Scanln(&option)

	if err != nil {
		log.Fatal(err)
	}

	return option
}

func scan() string {
	var value string
	_, err := fmt.Scanln(&value)

	if err != nil {
		log.Fatal(err)
	}

	return value
}

func scanConvert() int {

	number, err := strconv.Atoi(scan())

	if err != nil {
		fmt.Println(err)
	}

	return number
}

func getOperands() model.Data {
	numbers := model.Data{}

	information.EnterFirst()
	numbers.Num1 = scanConvert()

	utils.ClearScreen()

	information.EnterSecond()
	numbers.Num2 = scanConvert()
	return numbers
}

func LaunchExecution(option string) bool {
	utils.ClearScreen()
	switch option {
	case "1":

		fmt.Println("You Choose Addition")
		nums := getOperands()
		utils.ClearScreen()
		information.Output("Addition", nums.Num1+nums.Num2)

		return true
	case "2":

		fmt.Println("You Choose Subtraction")
		nums := getOperands()
		utils.ClearScreen()
		information.Output("Subtraction", nums.Num1-nums.Num2)

		return true
	case "3":

		fmt.Println("You Choose Multiplication")
		nums := getOperands()
		utils.ClearScreen()
		information.Output("Multiplication", nums.Num1*nums.Num2)

		return true
	case "4":

		fmt.Println("You Choose Divide")
		nums := getOperands()
		utils.ClearScreen()
		information.Output("Division", nums.Num1/nums.Num2)

		return true
	case "5":

		fmt.Println("You Choose to exit. Exiting ...")
		return false
	default:
		fmt.Println("Invalid Option! Exiting Application!")
		return false
	}
}
