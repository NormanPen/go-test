package main

import (
	"fmt"

	"github.com/normanpen/go-test/calculator"
	"github.com/normanpen/go-test/calculator/composite"
)

func main() {
	fmt.Println("23 + 42 = ", calculator.Add(23, 42))
	fmt.Println(calculator.Divide(17, 3))
	fmt.Println(calculator.Sum(1, 10))
	fmt.Println(calculator.SumUntil(1000))
	fmt.Println(calculator.IsSquareNumber(25))
	fmt.Println(calculator.MultiplyFromAToB(1, 10))
	fmt.Println(composite.Add(2, 6))

	point := composite.Point{
		X: 3,
		Y: 7,
	}
	fmt.Println(point.X, point.Y)
}
