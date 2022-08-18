package calculator

import "math"

func Abs(value int) int {
	if value > 00 {
		return value
	}

	return -value
}

func IsSquareNumber(value int) bool {
	//sqrt := math.Sqrt(float64(value))
	if sqrt := math.Sqrt(float64(value)); sqrt != float64(int(sqrt)) {
		return false
	} else {
		return true
	}

}

func RunOperation(operation string, left, right int) int {

	var result int

	switch operation { // operation kann weggelassen werden
	//case operation == "add"
	case "add":
		return left + right
	case "substract":
		return left - right
	default:
		// Intensionally left blank
	}
	return result
}
