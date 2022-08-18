package calculator


// For Schleife ( Zählschleife )
func Sum(min, max int) int {
	sum := 0

	for i := min; i <= max; i++ {
		sum += i
	}

	return sum
}

//Klassische While Schleife
func SumUntil(max int) int {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	return sum
}

// unendliche Schleife
func SumInfinite () {
	for {
		// ....
	}
}