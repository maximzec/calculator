package rpn

import (
	"calculator/stack"
	"strconv"
	"unicode"
)

func isOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	} else {
		return false
	}
}

func operatorPriority(s string) byte {

	operatorPriorities := map[string]byte{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	return operatorPriorities[s]

}

// TODO: Этот метод мне кажется некрасивым
// Хотелось бы его отрефакторить потом
func splitString(s string) []string {
	var tmp string
	output := make([]string, 0)

	for _, symbol := range s {
		if unicode.IsDigit(symbol) {
			tmp = string(append([]rune(tmp), symbol))
		} else {
			output = append(output, tmp)
			output = append(output, string(symbol))
			tmp = ""
		}
	}

	output = append(output, tmp)
	return output
}

func ConvertStrToRpn(s string) []string {

	operatorStack := stack.CreateNewStringStack()
	var output []string

	for _, token := range splitString(s) {
		if _, err := strconv.Atoi(token); err == nil {
			output = append(output, token)
		} else {
			if isOperator(token) {
				if operatorStack.Empty() {
					operatorStack.Push(token)
				} else {
					for operatorPriority(operatorStack.Pop()) > operatorPriority(token) {
						topStackOperator := operatorStack.Pop()
						output = append(output, topStackOperator)
					}
					operatorStack.Push(token)
				}
			} else {
				// TODO: Эта строка сигнализирует о нарушении SPR - принципа.
				// Нужен отдельный метод, валидирующий строку на регламент
				panic("Input string contains unknown symbols")
			}
		}
	}

	for !operatorStack.Empty() {
		output = append(output, operatorStack.Pop())
	}

	return output
}
