package main

import (
	"calculator/rpn"
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Printf("%v:", rpn.ConvertStrToRpn(input))
}
