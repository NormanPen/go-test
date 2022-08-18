package calculator

var Pi float64 = 3.141592654 // Typ kann weggelsassen werden
// Pi := 3.141592654 nicht im package möglich

func Add(left, right int) int {
	// var sum int = left + right
	// var sum = left + right
	// sum := left + right  kürzere Zuweisung innerhalb von funktion
	var sum int
	sum = left + right

	return sum //left + right
}

func Divide(left, right int) (quotient int, remainder int) {
	quotient = left / right
	remainder = left % right
	return //left / right, left % right
}
