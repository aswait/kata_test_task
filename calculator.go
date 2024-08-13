package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	expression, err := readExpression()
	if err != nil {
		panic(err)
	}
	fmt.Print(expression)
}

func readExpression() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := strings.Fields(scanner.Text())
	if len(inputString) != 3 {
		err := errors.New("the format of a mathematical operation does not satisfy the condition - two operands and one operator (+, -, /, *).")
		return inputString, err
	}
	return inputString, nil
}
