package main

import (
	"calculator/stack"
	"fmt"
	"unicode"
)

func isOperator(b rune) bool {
	if b == '+' || b == '-' || b == '*' || b == '/' {
		return true
	} else {
		return false
	}
}

func convertStrToRpn(s string) string {

	// operators = [...]byte{'+', '-', '/', '*'}

	operatorPriorities := map[byte]byte{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 3,
	}
	operatorStack := stack.CreateNewByteStack()
	var output string

	for _, token := range s {
		if unicode.IsDigit(token) {
			output = string(append([]rune(output), token))
		}
		if isOperator(token) {
			if operatorStack.Empty() {
				operatorStack = operatorStack.Push(byte(token))
			} else {

				for operatorPriorities[operatorStack.Peek()] > operatorPriorities[byte(token)] {

					topStackOperator, _ := operatorStack.Pop()
					output = string(append([]rune(output), rune(topStackOperator)))
					fmt.Println("Pushed op to output")
				}
			}
		}
	}

	for !operatorStack.Empty() {

		topStackOperator, _ := operatorStack.Pop()

		output = string(append([]rune(output), rune(topStackOperator)))
	}

	return output
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(convertStrToRpn(input))
	// bs := stack.CreateNewByteStack()

	// bs = bs.Push(2)
	// bs = bs.Push(2)
	// bs = bs.Push(2)

	// b, _ := bs.Pop()
	// fmt.Println(b)
}
