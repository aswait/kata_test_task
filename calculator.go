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

type operand struct {
	value      int
	num_system string
}

func main() {
	expression, err := readExpression()
	if err != nil {
		panic(err)
	}
	f, err := operandToStruct(expression[0])
	fmt.Print(f)
	if err != nil {
		panic(err)
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
	myErr := errors.New("string is not a mathematical operation")
	arab, err := strconv.Atoi(operand)
	if err != nil || arab > 10 || arab <= 0 {
		return 0, myErr
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

	arab, err := operandIsArab(stringOperand)
	if err != nil {
		rome, err := operandIsRome(stringOperand)
		if err != nil {
			return operandStruct, err
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
