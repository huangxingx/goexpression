package main

import (
	"fmt"

	"github.com/huangxingx/goexpression"
)

func main() {
	expressStr := "1+2+3"
	exp := goexpression.NewExpress(expressStr)
	result := exp.Execute(nil)
	fmt.Printf("%s = %.2f", expressStr, result.(float64))
}
