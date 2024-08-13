package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romeToArab = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var arabToRome = []struct {
	value int
	rome  string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

type operand struct {
	value      int
	num_system string
}

func main() {
	expression, err := readExpression()
	if err != nil {
		panic(err)
	}
	operands, err := checkNumSystems(expression[0], expression[2])
	if err != nil {
		panic(err)
	}
	result, err := mathOperation(operands, expression[1])
	if err != nil {
		panic(err)
	}
	if operands[0].num_system == "arab" {
		fmt.Println(result)
	} else {
		arabResult, err := intToRome(result)
		if err != nil {
			panic(err)
		}
		fmt.Println(arabResult)
	}
}

func readExpression() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := strings.Fields(scanner.Text())
	if len(inputString) > 3 {
		err := errors.New("the format of a mathematical operation does not satisfy the condition - two operands and one operator (+, -, /, *).")
		return inputString, err
	} else if len(inputString) < 3 {
		err := errors.New("string is not a mathematical operation")
		return inputString, err
	}
	return inputString, nil
}

func operandIsArab(operand string) (int, error) {
	arab, err := strconv.Atoi(operand)

	if err != nil {
		err := errors.New("string is not a mathematical operation")
		return 0, err
	} else if arab > 10 || arab <= 0 {
		err := errors.New("too large values that do not satisfy the condition")
		return 0, err
	}
	return arab, nil
}

func operandIsRome(operand string) (int, error) {
	myErr := errors.New("string is not a mathematical operation")
	if rome, ok := romeToArab[operand]; ok {
		return rome, nil
	}
	return 0, myErr
}

func operandToStruct(stringOperand string) (operand, error) {
	var operandStruct operand

	arab, arab_err := operandIsArab(stringOperand)
	if arab_err != nil {
		rome, err := operandIsRome(stringOperand)
		if err != nil {
			return operandStruct, arab_err
		} else {
			operandStruct.value = rome
			operandStruct.num_system = "rome"
		}

	} else {
		operandStruct.value = arab
		operandStruct.num_system = "arab"
	}
	return operandStruct, nil
}

func checkNumSystems(firstOperand string, secondOperand string) ([]operand, error) {
	operandsArray := make([]operand, 2)

	firstOperandStruct, err := operandToStruct(firstOperand)
	if err != nil {
		return operandsArray, err
	}

	secondOperandStruct, err := operandToStruct(secondOperand)
	if err != nil {
		return operandsArray, err
	}

	if firstOperandStruct.num_system != secondOperandStruct.num_system {
		err = errors.New("different number systems are used simultaneously.")
		return operandsArray, err
	}

	operandsArray[0], operandsArray[1] = firstOperandStruct, secondOperandStruct
	return operandsArray, err
}

func mathOperation(operands []operand, operator string) (int, error) {
	var result int

	switch operator {
	case "+":
		result = operands[0].value + operands[1].value
	case "-":
		result = operands[0].value - operands[1].value
	case "/":
		result = operands[0].value / operands[1].value
	case "*":
		result = operands[0].value * operands[1].value
	default:
		err := errors.New("mathematical operation does not satisfy the condition - only operators (+, -, /, *)")
		return 0, err
	}
	return result, nil
}

func intToRome(arab int) (string, error) {
	if arab <= 0 {
		err := errors.New("a mathematical operation on Roman numerals cannot be less than or equal to zero")
		return "", err
	}
	var romeString string
	for _, romeStruct := range arabToRome {
		for arab >= romeStruct.value {
			romeString += romeStruct.rome
			arab -= romeStruct.value
		}
	}
	return romeString, nil
}
