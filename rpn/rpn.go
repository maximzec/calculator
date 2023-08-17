package rpn

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

func ConvertStrToRpn(s string) string {

	operatorPriorities := map[byte]byte{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	operatorStack := stack.CreateNewByteStack()
	var output string

	for _, token := range s {
		if unicode.IsDigit(token) {
			output = string(append([]rune(output), token))
		}
		if isOperator(token) {
			if operatorStack.Empty() {
				operatorStack.Push(byte(token))
			} else {
				for operatorPriorities[operatorStack.Peek()] > operatorPriorities[byte(token)] {
					topStackOperator := operatorStack.Pop()
					output = string(append([]rune(output), rune(topStackOperator)))
				}
				operatorStack.Push(byte(token))
			}
		}
	}

	for !operatorStack.Empty() {

		fmt.Println("Pushed op to output")
		topStackOperator := operatorStack.Pop()
		output = string(append([]rune(output), rune(topStackOperator)))
	}

	return output
}
