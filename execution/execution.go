package execution

import (
	"fmt"
	"strconv"

	"karakorum.in/dev/go-calculator/information"
	loggs "karakorum.in/dev/go-calculator/logger"
	model "karakorum.in/dev/go-calculator/models"
	"karakorum.in/dev/go-calculator/utils"
)

func GetOperation() string {
	var option string
	_, err := fmt.Scanln(&option)

	if err != nil {
		loggs.LogData(err.Error())
	}

	return option
}

func scan() string {
	var value string
	_, err := fmt.Scanln(&value)

	if err != nil {
		fmt.Println("Something went wrong while scanning.")
		loggs.LogData(err.Error())
	}

	return value
}

func scanConvert() int {

	number, err := strconv.Atoi(scan())

	if err != nil {
		fmt.Println("Inavlid value.")
		loggs.LogData(err.Error())
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
		result := nums.Num1 + nums.Num2
		information.Output("Addition", result)
		loggs.LogData("Addition result : " + strconv.Itoa(result))
		return true
	case "2":

		fmt.Println("You Choose Subtraction")
		nums := getOperands()
		utils.ClearScreen()
		result := nums.Num1 - nums.Num2
		information.Output("Subtraction", result)
		loggs.LogData("Subtraction result : " + strconv.Itoa(result))
		return true
	case "3":

		fmt.Println("You Choose Multiplication")
		nums := getOperands()
		utils.ClearScreen()
		result := nums.Num1 * nums.Num2
		information.Output("Multiplication", result)
		loggs.LogData("Multiplication result : " + strconv.Itoa(result))
		return true
	case "4":

		fmt.Println("You Choose Divide")
		nums := getOperands()
		utils.ClearScreen()
		result := nums.Num1 / nums.Num2
		information.Output("Division", result)
		loggs.LogData("Division result : " + strconv.Itoa(result))
		return true
	case "5":
		fmt.Println("You Choose to exit. Exiting ...")
		loggs.LogData("You Choose to exit. Exiting ...")
		return false
	default:
		fmt.Println("Invalid Option! Exiting Application!")
		loggs.LogData("Invalid Option! Exiting Application!")
		return false
	}
}
