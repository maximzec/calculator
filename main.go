package main

import (
	"calculator/rpn"
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(rpn.ConvertStrToRpn(input))
}
